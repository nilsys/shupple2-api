package config

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/labstack/gommon/log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssm"

	"github.com/pkg/errors"

	"gopkg.in/go-playground/validator.v9"
	"gopkg.in/yaml.v3"
)

// build時にldflagで注入する
var Version = "unknown"

const (
	Region           = "ap-northeast-1"
	StgSsmKey        = "sw-stg-media-api-config"
	PrdSsmKey        = "sw-prd-media-api-config"
	StgContainerName = "sw-stg-media-api"
	PrdContainerName = "sw-prd-media-api"
)

func GetConfig(filename FilePath) (*Config, error) {
	// 環境を判断
	url := os.Getenv("ECS_CONTAINER_METADATA_URI")
	if utf8.RuneCountInString(url) > 0 {
		resp, err := http.Get(url)
		if err != nil || resp.StatusCode == http.StatusOK {
			var metaData MetaData
			bodyBytes, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				log.Debug("failed aws meta data read response body")
			}
			err = json.Unmarshal(bodyBytes, &metaData)
			if err != nil {
				log.Debug("failed aws meta data unmarshal")
			}
			log.Debugf("ECS CONTAINER METADATA: %s", string(bodyBytes))
			return getConfigFromSSM(metaData.GetSSMKEY())
		}
		defer func() {
			err := resp.Body.Close()
			if err != nil {
				log.Debug("failed to close response body")
			}
		}()
	}

	// ローカルの場合
	f, err := os.Open(string(filename))
	if err != nil {
		return nil, errors.Wrap(err, "failed to read config file")
	}

	var config Config
	if err = yaml.NewDecoder(f).Decode(&config); err != nil {
		return nil, errors.WithStack(err)
	}

	if err := validator.New().Struct(&config); err != nil {
		return nil, err
	}

	config.Version = Version

	return &config, nil
}

func getConfigFromSSM(ssmKey string) (*Config, error) {
	var config Config
	res, err := fetchParameterStore(ssmKey)
	if err != nil {
		return nil, err
	}
	r := strings.NewReader(res)
	if err := yaml.NewDecoder(r).Decode(&config); err != nil {
		return nil, errors.WithStack(err)
	}
	if err := validator.New().Struct(&config); err != nil {
		return nil, err
	}
	return &config, nil
}

// SSMから引数で受けたkeyのvalueを取得
func fetchParameterStore(param string) (string, error) {
	sess := session.Must(session.NewSession())
	svc := ssm.New(
		sess,
		aws.NewConfig().WithRegion(Region),
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

// SSMKEYを環境で切り替え
func (m *MetaData) GetSSMKEY() string {
	if m.Name == PrdContainerName {
		return PrdSsmKey
	}
	return StgSsmKey
}
