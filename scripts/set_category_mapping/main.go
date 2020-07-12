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
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model"
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

		msArea := &entity.MetasearchArea{
			MetasearchAreaID:   areaID,
			MetasearchAreaType: model.AreaCategoryTypeArea,
			AreaCategoryID:     categoryID,
		}
		res := s.DB.Create(msArea)
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

		msArea := &entity.MetasearchArea{
			MetasearchAreaID:   subAreaID,
			MetasearchAreaType: model.AreaCategoryTypeSubArea,
			AreaCategoryID:     categoryID,
		}
		res := s.DB.Create(msArea)
		if err := res.Error; err != nil {
			log.Println(errors.Wrapf(err, "failed to set sub_area; id = %d", subAreaID))
		}
		if res.RowsAffected == 0 {
			log.Printf("sub_area not updated; id = %d\n", subAreaID)
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
