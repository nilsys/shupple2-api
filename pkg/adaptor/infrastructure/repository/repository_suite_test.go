package repository

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jinzhu/gorm"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/pkg/errors"
)

// TODO: 2重管理になっている点、fileの相対位置を指定しないといけない点がイケてない
const (
	testDB         = "stayway_test"
	migrationsDir  = "./../../../../migrations"
	configFilePath = "./../../../../config.test.yaml"
)

var (
	db    *gorm.DB
	tests *Test
)

func TestRepository(t *testing.T) {
	RegisterFailHandler(Fail)

	RunSpecs(t, "Repository Suite")
}

var _ = BeforeSuite(func() {
	Expect(beforeSuite()).To(Succeed())
})

func beforeSuite() error {
	var err error
	tests, err = InitializeTest(configFilePath)
	if err != nil {
		return errors.Wrap(err, "failed to initialize test")
	}
	db = tests.DB

	if err := MigrateUp(tests.Config.Database, migrationsDir); err != nil {
		return errors.Wrap(err, "failed to migrate up")
	}

	return nil
}

func truncate(db *gorm.DB) {
	type Result struct {
		Name string
	}

	Expect(db.Exec("SET FOREIGN_KEY_CHECKS=0").Error).To(Succeed())

	q := db.
		Table("information_schema.tables").
		Select("table_name AS name").
		Where("table_schema = ?", testDB).
		Where("table_rows > 0 AND table_name != 'schema_migrations'")

	var res []*Result
	Expect(q.Scan(&res).Error).To(Succeed())

	for _, row := range res {
		Expect(q.Exec("truncate " + row.Name).Error).To(Succeed())
	}

	Expect(db.Exec("SET FOREIGN_KEY_CHECKS=1").Error).To(Succeed())
}

func prepareBucket(sess *session.Session, bucket string) error {
	s3c := s3.New(sess)

	_, err := s3c.CreateBucket(&s3.CreateBucketInput{
		Bucket: &bucket,
	})

	if err == nil {
		return nil
	}

	awsErr, ok := err.(awserr.Error)
	if !(ok && awsErr.Code() == s3.ErrCodeBucketAlreadyExists) {
		return err
	}

	// Bucketが既に存在している場合

	var errDelete error
	listInput := &s3.ListObjectsV2Input{Bucket: &bucket}
	err = s3c.ListObjectsV2Pages(listInput, func(page *s3.ListObjectsV2Output, lastPage bool) bool {
		for _, obj := range page.Contents {
			_, err := s3c.DeleteObject(&s3.DeleteObjectInput{
				Bucket: &bucket,
				Key:    obj.Key,
			})
			if err != nil {
				errDelete = err
				return false
			}
		}
		return true
	})

	if errDelete != nil {
		return errDelete
	}

	return err
}

func responseTestdata(name string) (*http.Response, error) {
	body, err := ioutil.ReadFile("testdata/" + name)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to read testdata/%s", name)
	}

	recoder := httptest.NewRecorder()
	recoder.WriteHeader(http.StatusOK)
	if _, err := recoder.Write(body); err != nil {
		return nil, errors.Wrap(err, "failed to write body")
	}

	return recoder.Result(), nil
}
