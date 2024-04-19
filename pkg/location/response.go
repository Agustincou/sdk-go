package location

type Country struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type City struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type GeoLocation struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type State struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type StateByZipCodeResponse struct {
	ZipCode string `json:"zip_code"`
	State   struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"state"`
	Country struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"country"`
}

type StateDetailsResponse struct {
	ID             string      `json:"id"`
	Name           string      `json:"name"`
	Country        Country     `json:"country"`
	GeoInformation GeoLocation `json:"geo_information"`
	TimeZone       string      `json:"time_zone"`
	TimeZoneName   string      `json:"time_zone_name"`
	Cities         []City      `json:"cities"`
}

type CityDetailsResponse struct {
	ID             string      `json:"id"`
	Name           string      `json:"name"`
	State          State       `json:"state"`
	Country        Country     `json:"country"`
	GeoInformation GeoLocation `json:"geo_information"`
}
