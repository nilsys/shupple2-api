package main

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"path/filepath"
	"strconv"

	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
)

const (
	areaCSVName    = "area.csv"
	subAreaCSVName = "subarea.csv"
	innTypeCSVName = "inntype.csv"
	tagCSVName     = "tag.csv"
)

type Script struct {
	DB *gorm.DB
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	script, err := InitializeScript("./config.yaml")
	if err != nil {
		return errors.Wrap(err, "failed to initialize script")
	}

	return script.Run(os.Args)
}

func (s Script) Run(args []string) error {
	if len(args) != 2 {
		return errors.New("exactly one argument csv_directory required")
	}

	dir := args[1]

	if err := s.updateArea(dir); err != nil {
		return errors.Wrap(err, "failed to update areas")
	}

	if err := s.updateSubArea(dir); err != nil {
		return errors.Wrap(err, "failed to update sub areas")
	}

	if err := s.updateInnType(dir); err != nil {
		return errors.Wrap(err, "failed to update inn types")
	}

	if err := s.updateTag(dir); err != nil {
		return errors.Wrap(err, "failed to update tags")
	}

	return nil
}

func skipRecords(reader *csv.Reader, num int) error {
	for i := 0; i < num; i++ {
		if _, err := reader.Read(); err != nil {
			return errors.WithStack(err)
		}
	}
	return nil
}

func (s Script) updateArea(dir string) error {
	const skipRecordsNum = 4

	file, err := os.Open(filepath.Join(dir, areaCSVName))
	if err != nil {
		return errors.Wrap(err, "failed to open csv")
	}
	defer file.Close()

	reader := csv.NewReader(file)
	if err := skipRecords(reader, skipRecordsNum); err != nil {
		return errors.Wrap(err, "failed to skip csv record")
	}

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Println(errors.Wrap(err, "failed to decode csv row"))
			continue
		}

		areaID := mustToInt(record[0])
		categoryID := mustToInt(record[2])
		if !(areaID > 0 && categoryID > 0) {
			continue
		}

		res := s.DB.Exec("update area_category set metasearch_area_id = ? where id = ?", areaID, categoryID)
		if err := res.Error; err != nil {
			log.Println(errors.Wrapf(err, "failed to set area; id = %d", areaID))
		}
		if res.RowsAffected == 0 {
			log.Printf("area not updated; id = %d\n", areaID)
		}
	}
	return nil
}

func (s Script) updateSubArea(dir string) error {
	const skipRecordsNum = 4

	file, err := os.Open(filepath.Join(dir, subAreaCSVName))
	if err != nil {
		return errors.Wrap(err, "failed to open csv")
	}
	defer file.Close()

	reader := csv.NewReader(file)
	if err := skipRecords(reader, skipRecordsNum); err != nil {
		return errors.Wrap(err, "failed to skip csv record")
	}

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Println(errors.Wrap(err, "failed to decode csv row"))
			continue
		}

		subAreaID := mustToInt(record[1])
		categoryID := mustToInt(record[5])
		if !(subAreaID > 0 && categoryID > 0) {
			continue
		}

		res := s.DB.Exec("update area_category set metasearch_sub_area_id = ? where id = ?", subAreaID, categoryID)
		if err := res.Error; err != nil {
			log.Println(errors.Wrapf(err, "failed to set sub_area; id = %d", subAreaID))
		}
		if res.RowsAffected == 0 {
			log.Printf("sub_area not updated; id = %d\n", subAreaID)
		}
	}
	return nil
}

func (s Script) updateInnType(dir string) error {
	const skipRecordsNum = 2

	file, err := os.Open(filepath.Join(dir, innTypeCSVName))
	if err != nil {
		return errors.Wrap(err, "failed to open csv")
	}
	defer file.Close()

	reader := csv.NewReader(file)
	if err := skipRecords(reader, skipRecordsNum); err != nil {
		return errors.Wrap(err, "failed to skip csv record")
	}

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Println(errors.Wrap(err, "failed to decode csv row"))
			continue
		}

		innTypeID := mustToInt(record[1])
		categoryID := mustToInt(record[5])
		if !(innTypeID > 0 && categoryID > 0) {
			continue
		}

		res := s.DB.Exec("update theme_category set metasearch_inn_type_id = ? where id = ?", innTypeID, categoryID)
		if err := res.Error; err != nil {
			log.Println(errors.Wrapf(err, "failed to set inn_type; id = %d", innTypeID))
		}
		if res.RowsAffected == 0 {
			log.Printf("inn_type not updated; id = %d\n", innTypeID)
		}
	}
	return nil
}

func (s Script) updateTag(dir string) error {
	const skipRecordsNum = 2

	file, err := os.Open(filepath.Join(dir, tagCSVName))
	if err != nil {
		return errors.Wrap(err, "failed to open csv")
	}
	defer file.Close()

	reader := csv.NewReader(file)
	if err := skipRecords(reader, skipRecordsNum); err != nil {
		return errors.Wrap(err, "failed to skip csv record")
	}

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Println(errors.Wrap(err, "failed to decode csv row"))
			continue
		}

		tagID := mustToInt(record[1])
		categoryID := mustToInt(record[5])
		if !(tagID > 0 && categoryID > 0) {
			continue
		}

		res := s.DB.Exec("update theme_category set metasearch_tag_id = ? where id = ?", tagID, categoryID)
		if err := res.Error; err != nil {
			log.Println(errors.Wrapf(err, "failed to set tag; id = %d", tagID))
		}
		if res.RowsAffected == 0 {
			log.Printf("tag not updated; id = %d\n", tagID)
		}
	}
	return nil
}

func mustToInt(s string) int {
	if s == "#N/A" || s == "" {
		return 0
	}

	res, err := strconv.Atoi(s)
	if err != nil {
		log.Printf("%+v\n", errors.WithStack(err))
	}
	return res
}
