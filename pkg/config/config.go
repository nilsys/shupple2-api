package config

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssm"

	"github.com/pkg/errors"

	"gopkg.in/go-playground/validator.v9"
	"gopkg.in/yaml.v3"
)

const (
	region                 = "ap-northeast-1"
	ssmKeyFormat           = "sw-%s-media-api-config"
	prdContainerNamePrefix = "sw-prd"

	DefaultConfigFilePath = "config.yaml"
)

func GetConfig(filename FilePath) (*Config, error) {
	// 環境を判断
	url := os.Getenv("ECS_CONTAINER_METADATA_URI")
	if url != "" {
		resp, err := http.Get(url)
		if err != nil {
			return nil, errors.Wrap(err, "failed to fetch ecs container metadata")
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			return nil, errors.Wrapf(err, "ecs container metadata api returns not OK status; %d", resp.StatusCode)
		}

		var metaData MetaData
		if err := json.NewDecoder(resp.Body).Decode(&metaData); err != nil {
			return nil, errors.Wrap(err, "failed aws meta data unmarshal")
		}

		return GetConfigFromSSM(getEnvFromEcsMetadata(&metaData))
	}

	// ローカルの場合
	f, err := os.Open(string(filename))
	if err != nil {
		return nil, errors.Wrap(err, "failed to read config file")
	}
	defer f.Close()

	return loadConfig(f, EnvDev)
}

func GetConfigFromSSM(env Env) (*Config, error) {
	ssmKey := fmt.Sprintf(ssmKeyFormat, env)
	res, err := fetchParameterStore(ssmKey)
	if err != nil {
		return nil, err
	}

	return loadConfig(strings.NewReader(res), env)
}

func loadConfig(reader io.Reader, env Env) (*Config, error) {
	var config Config
	if err := yaml.NewDecoder(reader).Decode(&config); err != nil {
		return nil, errors.WithStack(err)
	}

	v := validator.New()
	v.RegisterStructValidation(URLRequiredValidation, URL{}) // URLが必ずrequiredになってしまうので微妙
	if err := v.Struct(&config); err != nil {
		return nil, errors.WithStack(err)
	}

	return &config, nil
}

// SSMから引数で受けたkeyのvalueを取得
func fetchParameterStore(param string) (string, error) {
	sess := session.Must(session.NewSession())
	svc := ssm.New(
		sess,
		aws.NewConfig().WithRegion(region),
	)
	res, err := svc.GetParameter(&ssm.GetParameterInput{
		Name:           aws.String(param),
		WithDecryption: aws.Bool(true),
	})
	if err != nil {
		return "fetch Error", err
	}

	value := *res.Parameter.Value
	return value, nil
}

type MetaData struct {
	DockerID   string `json:"DockerId"`
	Name       string `json:"Name"`
	DockerName string `json:"DockerName"`
	Image      string `json:"Image"`
	ImageID    string `json:"ImageID"`
	Labels     struct {
		ComAmazonawsEcsCluster               string `json:"com.amazonaws.ecs.cluster"`
		ComAmazonawsEcsContainerName         string `json:"com.amazonaws.ecs.container-name"`
		ComAmazonawsEcsTaskArn               string `json:"com.amazonaws.ecs.task-arn"`
		ComAmazonawsEcsTaskDefinitionFamily  string `json:"com.amazonaws.ecs.task-definition-family"`
		ComAmazonawsEcsTaskDefinitionVersion string `json:"com.amazonaws.ecs.task-definition-version"`
	} `json:"Labels"`
	DesiredStatus string `json:"DesiredStatus"`
	KnownStatus   string `json:"KnownStatus"`
	Limits        struct {
		CPU    int `json:"CPU"`
		Memory int `json:"Memory"`
	} `json:"Limits"`
	CreatedAt time.Time `json:"CreatedAt"`
	Type      string    `json:"Type"`
	Networks  []struct {
		NetworkMode   string   `json:"NetworkMode"`
		IPv4Addresses []string `json:"IPv4Addresses"`
	} `json:"Networks"`
}

// TODO: 環境変数から取ったほうが良さそう
func getEnvFromEcsMetadata(m *MetaData) Env {
	if strings.HasPrefix(m.Name, prdContainerNamePrefix) {
		return EnvPrd
	}
	return EnvStg
}
