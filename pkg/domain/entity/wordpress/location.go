package wordpress

type Location struct {
	ID          int                `json:"id"`
	Date        Time               `json:"date"`
	DateGmt     Time               `json:"date_gmt"`
	GUID        Text               `json:"guid"`
	Modified    Time               `json:"modified"`
	ModifiedGmt Time               `json:"modified_gmt"`
	Slug        string             `json:"slug"`
	Status      Status             `json:"status"`
	Title       Text               `json:"title"`
	Categories  []int              `json:"categories"`
	LocationCat []int              `json:"location__cat"`
	Attributes  LocationAttributes `json:"acf"`
}

type LocationAttributes struct {
	OfficialURL  string      `json:"officialurl"`
	City         string      `json:"spotplace2"`
	Address      string      `json:"address"`
	Map          LocationMap `json:"spotmapaddress"`
	AccessTrain  string      `json:"access"`
	AccessBus    string      `json:"bus_access"`
	AccessCar    string      `json:"car_access"`
	OpeningHours string      `json:"time"`
	TEL          string      `json:"spottell"`
	Price        string      `json:"spotprice"`
	Instagram    string      `json:"instagram"`
	Inn          string      `json:"inn"`
}

type LocationMap struct {
	Lat string `json:"lat"`
	Lng string `json:"lng"`
}
