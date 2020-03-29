package main

import (
	"flag"
	"log"
	"math"

	"github.com/jinzhu/gorm"

	"github.com/stayway-corp/stayway-media-api/pkg/domain/model/serror"

	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"

	"github.com/pkg/errors"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/repository"
)

type Script struct {
	TouristSpotQueryRepo repository.TouristSpotQueryRepository
	DB                   *gorm.DB
}

// local: docker-compose run --rm app go run ./scripts/import_tourist_spot_vendor_rate -s (resource_path)
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

	touristSpots, err := s.TouristSpotQueryRepo.FindAll()
	if err != nil {
		return errors.Wrap(err, "failed to find all tourist_spot")
	}

	touristSpotsFromCsv := s.touristSpotFromCsv(*op)

	s.importVendorRate(touristSpots, touristSpotsFromCsv)

	return nil
}

// TODO: stgでマッチするか確認
func (s Script) importVendorRate(touristSpots []*entity.TouristSpot, touristSpotsFromCSV []touristSpot) {
	for _, touristSpotFromCsv := range touristSpotsFromCSV {
		if touristSpotFromCsv.Lng == 0 {
			continue
		}
		for _, touristSpot := range touristSpots {
			if math.Round(touristSpot.Lng*10000)/10000 == touristSpotFromCsv.Lng && math.Round(touristSpot.Lat*10000)/10000 == touristSpotFromCsv.Lat {
				if err := s.DB.Exec("UPDATE tourist_spot SET vendor_rate = ? WHERE id = ?;", touristSpotFromCsv.Rate, touristSpot.ID).Error; err != nil {
					log.Fatal(err)
				}
			}
		}
	}
}
