package main

import (
	"encoding/csv"
	"flag"
	"io"
	"log"
	"os"
	"strconv"

	"go.uber.org/zap"

	"github.com/stayway-corp/stayway-media-api/pkg/adaptor/logger"

	"github.com/pkg/errors"

	"github.com/jinzhu/gorm"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model/serror"
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

		id, err := strconv.Atoi(record[1])
		if err != nil {
			logger.Debug("invalid csv format column id", zap.String("id", record[1]))
			continue
		}
		sortOrder, err := strconv.Atoi(record[4])
		if record[4] == "" {
			continue
		}
		if err != nil {
			logger.Debug("invalid csv format column sort_order", zap.String("sort_order", record[4]))
			continue
		}

		if err := s.DB.Exec("UPDATE area_category SET sort_order = ? WHERE id = ?", sortOrder, id).Error; err != nil {
			logger.Debug("failed update area_category.sort_order", zap.String("id=", record[4]))
			continue
		}
	}

	return nil
}
