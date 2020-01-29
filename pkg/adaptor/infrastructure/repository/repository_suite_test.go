package repository

import (
	"testing"

	"github.com/stayway-corp/stayway-media-api/pkg/config"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jinzhu/gorm"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

const (
	testDB         = "stayway_test"
	migrationsDir  = "./../../../../migrations"
	configFilePath = "./../../../../config.test.yaml"
)

var db *gorm.DB

func TestRepository(t *testing.T) {
	RegisterFailHandler(Fail)

	RunSpecs(t, "Repository Suite")
}

var _ = BeforeSuite(func() {
	var err error
	db, err = prepareDB()
	Expect(err).To(Succeed())
})

func prepareDB() (*gorm.DB, error) {
	c, err := config.GetConfig(configFilePath)
	if err != nil {
		return nil, err
	}

	db, err := ProvideDB(c)
	if err != nil {
		return nil, err
	}

	if err := migrateUp(c.Database); err != nil {
		return nil, err
	}

	return db, nil
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

func migrateUp(database string) error {
	source := "file://" + migrationsDir
	db := "mysql://" + database
	m, err := migrate.New(source, db)
	if err != nil {
		return err
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return err
	}

	return nil
}
