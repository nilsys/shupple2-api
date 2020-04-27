package main

import (
	"encoding/csv"
	"io"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type (
	touristSpot struct {
		Name string
		Rate float64
		Lng  float64
		Lat  float64
	}
)

func (s *Script) touristSpotFromCsv(filepath string) []touristSpot {
	touristSpots := make([]touristSpot, 0)

	// TODO: どっかのドメインに置いて良いかも
	file, err := os.Open(filepath)
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
			log.Fatal(err)
		}
		rate, _ := strconv.ParseFloat(record[1], 64)
		lng := geoCodeToLng(record[2])
		lat := geoCodeToLat(record[2])
		touristSpots = append(touristSpots, touristSpot{
			Name: record[0],
			Rate: rate,
			Lng:  lng,
			Lat:  lat,
		})
	}

	return touristSpots
}

func geoCodeToLng(str string) float64 {
	items := strings.Split(str, ",")
	if len(items) < 2 {
		return 0.0
	}
	lng, _ := strconv.ParseFloat(items[1], 64)
	return math.Round(lng*10000) / 10000
}

func geoCodeToLat(str string) float64 {
	items := strings.Split(str, ",")
	if len(items) < 2 {
		return 0.0
	}
	lat, _ := strconv.ParseFloat(items[0], 64)
	return math.Round(lat*10000) / 10000
}
