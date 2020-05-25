package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/stayway-corp/stayway-media-api/pkg/adaptor/logger"

	"github.com/pkg/errors"
	"googlemaps.github.io/maps"

	"github.com/stayway-corp/stayway-media-api/pkg/config"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"

	"github.com/jinzhu/gorm"
)

type (
	Script struct {
		DB     *gorm.DB
		Config *config.Config
	}

	Config struct {
		GoogleMap struct {
			APIKey string `yaml:"api_key" validate:"required"`
		} `yaml:"google_map" validate:"required"`
		MicrosoftTranslate struct {
			APIKey string `yaml:"api_key" validate:"required"`
		} `yaml:"microsoft_translate" validate:"required"`
	}

	TransScript struct {
		Text string `json:"text"`
	}

	Translation struct {
		Translations []TransScript `json:"translations"`
	}
)

const (
	limit   = 100
	uriBase = "https://api.cognitive.microsofttranslator.com"
	uriPath = "/translate?api-version=3.0"
)

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
	var config Config
	if err := s.Config.Scripts.Decode(&config); err != nil {
		return errors.Wrap(err, "failed to load script config")
	}

	c, err := maps.NewClient(maps.WithAPIKey(config.GoogleMap.APIKey))
	if err != nil {
		return errors.Wrap(err, "failed auth google_map api client")
	}

	lastID := 0
	for {
		touristSpots, err := s.FindNoGeoTouristSpot(lastID)
		if err != nil {
			return errors.Wrap(err, "failed find no geo tourist_spot")
		}
		if len(touristSpots) == 0 {
			break
		}
		for _, touristSpot := range touristSpots {
			geocodeResult, err := s.getGeocodeResult(c, touristSpot.Address, config.MicrosoftTranslate.APIKey)
			if err != nil {
				logger.Debug(fmt.Sprintf("not found geocoding id=%d", touristSpot.ID))
				continue
			}
			if len(geocodeResult) == 0 {
				logger.Debug(fmt.Sprintf("not found geocoding id=%d", touristSpot.ID))
				continue
			}
			err = s.UpdateTouristSpotGeo(touristSpot.ID, geocodeResult[0].Geometry.Location)
			if err != nil {
				return errors.Wrap(err, "failed update tourist_spot lat lng")
			}
		}
		lastID = touristSpots[len(touristSpots)-1].ID
	}
	return nil
}

func (s *Script) FindNoGeoTouristSpot(lastID int) ([]*entity.TouristSpotTiny, error) {
	var rows []*entity.TouristSpotTiny

	if err := s.DB.Where("id > ?", lastID).Where("lat is NULL AND lng is NULL").Order("id").Limit(limit).Find(&rows).Error; err != nil {
		return nil, errors.Wrap(err, "failed find no geo tourist_spot")
	}

	return rows, nil
}

func (s *Script) UpdateTouristSpotGeo(id int, latLng maps.LatLng) error {
	if err := s.DB.Exec("UPDATE tourist_spot SET lat = ?, lng = ? WHERE id = ?", latLng.Lat, latLng.Lng, id).Error; err != nil {
		return errors.Wrap(err, "failed update tourist_spot.lat & tourist_spot.lng")
	}
	return nil
}

func (s *Script) getGeocodeResult(c *maps.Client, name string, translateAPIKey string) ([]maps.GeocodingResult, error) {
	resp, err := getResponseFromGeocodeAPI(c, name, "ja")
	if err != nil {
		return nil, errors.Wrapf(err, "failed to getResponseFromGeocodeAPI")
	}
	if len(resp) > 0 {
		return resp, nil
	}

	// Geocode API から生データの地名で値が取得できなかった場合
	// ①Nameの中に括弧があったら()より前の文字列をを取り出す(ex. 嘉峪関（カヨクカン）→ 嘉峪関)
	excludeBracketsName := excludeBrackets(name)
	resp, err = getResponseFromGeocodeAPI(c, excludeBracketsName, "ja")
	if err != nil {
		return nil, errors.Wrapf(err, "failed to getResponseFromGeocodeAPI")
	}
	if len(resp) > 0 {
		return resp, nil
	}

	// ②APIによる英語翻訳
	addressToEn, err := s.translate(name, translateAPIKey)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to Translate")
	}
	resp, err = getResponseFromGeocodeAPI(c, addressToEn, "en")
	if err != nil {
		return nil, errors.Wrapf(err, "failed to getResponseFromGeocodeAPI")
	}
	if len(resp) > 0 {
		return resp, nil
	}

	// ③手動で緯度経度が取れるような地名を用意
	addressByManual, err := manualTranslate(name)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to manualTranslate")
	}
	resp, err = getResponseFromGeocodeAPI(c, addressByManual, "")
	if err != nil {
		return nil, errors.Wrapf(err, "failed to getResponseFromGeocodeAPI")
	}
	if len(resp) > 0 {
		return resp, nil
	}

	// ①括弧整形, ②英語翻訳, ③手動翻訳しても緯度経度データが取得できないときはerrにする
	return nil, errors.Wrapf(err, "not fount geocode result: %v", addressByManual)
}

func excludeBrackets(name string) string {
	if strings.Contains(name, "(") {
		return strings.Split(name, "(")[0]
	}
	if strings.Contains(name, "（") {
		return strings.Split(name, "（")[0]
	}
	return name
}

func getResponseFromGeocodeAPI(c *maps.Client, address, lang string) ([]maps.GeocodingResult, error) {
	req := &maps.GeocodingRequest{
		Address:  address,
		Language: lang,
	}
	return c.Geocode(context.Background(), req)
}

// see: https://github.com/MicrosoftTranslator/Text-Translation-API-V3-Go/blob/master/Translate.go
// MEMO: 6/30まで有効
func (s *Script) translate(originText string, translateAPIKey string) (string, error) {
	// requestを作成
	uri := uriBase + uriPath + "&to=en"
	r := strings.NewReader("[{\"Text\" : \"" + originText + "\"}]")
	req, err := http.NewRequest("POST", uri, r)
	if err != nil {
		return "", err
	}
	setHeader(req, translateAPIKey)

	client := &http.Client{Timeout: time.Second * 2}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var result []Translation
	if err := json.Unmarshal(body, &result); err != nil {
		return "", err
	}
	// 万が一Translate APIから翻訳結果が取得できない場合はerr
	if len(result) == 0 || len(result[0].Translations) == 0 {
		return "", fmt.Errorf("len(result) == 0 at translate")
	}
	return result[0].Translations[0].Text, nil
}

func setHeader(req *http.Request, apiKey string) {
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Content-Length", strconv.FormatInt(req.ContentLength, 10))
	req.Header.Add("Ocp-Apim-Subscription-Key", apiKey)
}

// APIによる英語翻訳によって取得した地名を用いても経度緯度データが取得できない地名リスト
func manualTranslate(strOrigin string) (string, error) {
	// key: 変換前の地名, value: 手動変換後の地名
	transMap := map[string]string{
		"カパシア":       "Kapasia",
		"トゥルミ":       "Turmi",
		"エイールススタージル": "Egilsstadir",
		"ポサリカ":       "Pozarica",
		"ロフトフース":     "Lofthus",
		"ナティティング":    "Natitingou",
		"ワンデュポダン":    "Wangdue Phodrang",
		"コヴァラム":      "Kovalam",
		"ハヴォルスヴォルール": "Hvolsvöllur",
		"ブロンドュオス":    "Blönduós",
		"カトビッツェ":     "Katowice",
		"ユッカスヤルビ":    "Jukkasjärvi",
		"カランドゥーラ":    "Kalandula",
		"カマグウェイ":     "Camagüey",
		"ユッラス":       "Ylläs",
		"ペリン":        "Pelin",
		"ネシャヴェトリル":   "Nesjavellir",
		"ハラホリン":      "Karakorum",
		"タメルザ":       "Tamerza",
		"ポブジカ":       "Phobjika",
		"バートライヘンハル":  "Bad Reichenhall",
		"ヘフナフィヨルド":   "Hornafjordur",
		"ムスティーク":     "Mustique",
		"ランコー":       "Lang Co",
		"瀬戸大橋":       "北備讃瀬戸大橋",
		"三段峡":        "SandankyoGorge",
		"澤当(ツェタン)":   "Tsetang",
	}

	if strByManual, ok := transMap[strOrigin]; ok {
		return strByManual, nil
	} else {
		return "", fmt.Errorf("not found %v in transMap", strOrigin)
	}
}
