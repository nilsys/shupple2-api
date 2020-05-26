package main

import (
	"fmt"
	"log"
	"path"
	"strconv"

	"github.com/stayway-corp/stayway-media-api/pkg/adaptor/infrastructure/dto"
	"github.com/stayway-corp/stayway-media-api/pkg/adaptor/logger"

	"github.com/stayway-corp/stayway-media-api/pkg/config"

	"github.com/stayway-corp/stayway-media-api/pkg/adaptor/infrastructure/client"

	"github.com/pkg/errors"

	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"

	"github.com/jinzhu/gorm"
)

type Script struct {
	DB     *gorm.DB
	Client client.Client
	Config *config.Config
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

func (s *Script) Run() error {
	excludeIDs, err := s.FindMappingTouristSpotIDs()
	if err != nil {
		return errors.Wrap(err, "failed list mapping tourist_spot ids")
	}

	touristSpots, err := s.FindTouristSpotNotExcludeIDs(excludeIDs)
	if err != nil {
		return errors.Wrap(err, "failed list tourist_spot")
	}

	for _, touristSpot := range touristSpots {
		if !touristSpot.Lng.Valid || !touristSpot.Lat.Valid {
			logger.Debug(fmt.Sprintf("tourist_spot.id=%d.Lng or Lat is null", touristSpot.ID))
			continue
		}

		var res dto.Inns
		opts := &client.Option{
			QueryParams: map[string][]string{},
		}
		u := s.Config.Stayway.Metasearch.BaseURL
		u.Path = path.Join(u.Path, "/api/inns")
		opts.QueryParams.Add("geocode", innsGeoCodeQuery(touristSpot.Lat.Float64, touristSpot.Lng.Float64))
		opts.QueryParams.Add("per_page", "1")
		if err := s.Client.GetJSON(u.String(), opts, &res); err != nil {
			logger.Debug("failed metasearch inns api")
			continue
		}

		innIDs := innIDs(&res)
		if len(innIDs) == 0 {
			continue
		}

		// 最初のinnのareaを取得する
		var ares dto.InnArea
		aopts := &client.Option{
			QueryParams: map[string][]string{},
		}
		au := s.Config.Stayway.Metasearch.BaseURL
		au.Path = path.Join(au.Path, "/api", fmt.Sprintf("/%d/inn_area", innIDs[0]))
		if err := s.Client.GetJSON(au.String(), aopts, &ares); err != nil {
			logger.Debug("failed metasearch inns area api")
			continue
		}

		// 紐づけるべきarea_categoryのmetasearchIDを取得
		metasearchID := mappingMetasearchAreaID(&ares)
		if metasearchID == 0 {
			logger.Debug("missing area id")
			continue
		}

		var areaCategory entity.AreaCategory
		if err := s.DB.Where("metasearch_area_id = ?", metasearchID).Or("metasearch_sub_area_id = ?", metasearchID).Or("metasearch_sub_sub_area_id = ?", metasearchID).First(&areaCategory).Error; err != nil {
			logger.Debug(fmt.Sprintf("failed find area_category metasearch_id=%d", metasearchID))
			continue
		}

		if err := s.DB.Exec("INSERT INTO tourist_spot_area_category(tourist_spot_id,area_category_id) VALUES(?,?)", touristSpot.ID, areaCategory.ID).Error; err != nil {
			return errors.Wrap(err, "failed insert tourist_spot_area_category")
		}
	}

	return nil
}

// すでにarea_categoryとマッピング済tourist_spotのidを返す
func (s *Script) FindMappingTouristSpotIDs() ([]int, error) {
	var rows []*entity.TouristSpotTiny

	if err := s.DB.Where("id IN (SELECT tourist_spot_id FROM tourist_spot_area_category)").Find(&rows).Error; err != nil {
		return nil, errors.Wrap(err, "failed list mapping tourist_spot")
	}

	ids := make([]int, len(rows))
	for i, touristSpot := range rows {
		ids[i] = touristSpot.ID
	}

	return ids, nil
}

func (s *Script) FindTouristSpotNotExcludeIDs(excludeIDs []int) ([]*entity.TouristSpotTiny, error) {
	var rows []*entity.TouristSpotTiny

	if err := s.DB.Not("id", excludeIDs).Find(&rows).Error; err != nil {
		return nil, errors.Wrap(err, "failed find tourist_spot not in ids")
	}
	return rows, nil
}

func innsGeoCodeQuery(lat, lng float64) string {
	return strconv.FormatFloat(lat, 'f', -1, 64) + "," + strconv.FormatFloat(lng, 'f', -1, 64)
}

func innIDs(inns *dto.Inns) []int {
	innIDs := make([]int, len(inns.Inns))
	for i, inn := range inns.Inns {
		innIDs[i] = inn.ID
	}
	return innIDs
}

// 紐づいてるareaの最下層のIDを返す
func mappingMetasearchAreaID(ares *dto.InnArea) int {
	// areaと紐づいてない場合は下とも紐づいてないのでreturn
	// errorパターン
	if ares.Area.ID == 0 {
		return 0
	}

	if ares.SubSubArea.ID != 0 {
		return ares.SubSubArea.ID
	}

	if ares.SubArea.ID != 0 {
		return ares.SubArea.ID
	}

	return ares.Area.ID
}
