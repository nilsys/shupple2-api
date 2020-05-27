package main

import (
	"encoding/csv"
	"flag"
	"io"
	"log"
	"os"

	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model"

	"github.com/stayway-corp/stayway-media-api/pkg/adaptor/logger"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model/serror"
	"go.uber.org/zap"

	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
)

type Script struct {
	DB *gorm.DB
}

// local: docker-compose run --rm app go run ./scripts/import_interest -s ./scripts/import_interest/interest.csv
func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	script, err := InitializeScript("./config.yaml")
	if err != nil {
		return errors.Wrap(err, "faled to initialize script")
	}

	return script.Run()
}

func (s Script) Run() error {
	op := flag.String("s", "", "Resource path")
	flag.Parse()
	if *op == "" {
		// TODO: serrorではない
		return serror.New(nil, serror.CodeInvalidParam, "file not found")
	}

	file, err := os.Open(*op)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	reader := csv.NewReader(file)

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return errors.Wrap(err, "failed read csv")
		}

		interest := NewInterest(record[1], record[2])

		if err := s.DB.Save(interest).Error; err != nil {
			logger.Debug("failed save interest", zap.String("id=", record[0]))
			continue
		}
	}

	return nil
}

func NewInterest(name, group string) *entity.Interest {
	var interestGroup model.InterestGroup

	switch group {
	case "旅のスタイル":
		interestGroup = model.InterestGroupStyle
	case "シーン":
		interestGroup = model.InterestGroupScene
	case "グルメ":
		interestGroup = model.InterestGroupGourmet
	case "ライフスタイル":
		interestGroup = model.InterestGroupLifeStyle
	case "アクティビティ":
		interestGroup = model.InterestGroupActivity
	case "スポーツ":
		interestGroup = model.InterestGroupSport
	default:
		interestGroup = model.InterestGroupUndefined
	}
	return &entity.Interest{
		Name:          name,
		InterestGroup: interestGroup,
	}
}
