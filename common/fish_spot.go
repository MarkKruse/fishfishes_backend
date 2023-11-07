package common

type Fish_spots struct {
	Fish_spots []Fish_spot `json:"fish_spots"`
}

type Fish_spot struct {
	Id      string  `json:"id"`
	Marker  Marker  `json:"marker"`
	Catches []Catch `json:"catches"`
}

type Catch struct {
	Fish      string    `json:"id"`
	Number    int       `json:"number"`
	Size      float32   `json:"size"`
	Equipment Equipment `json:"equipment"`
	Deep      int       `json:"deep"`
	Time      string    `json:"time"` //Morning, Day, Afternoon, night
}

type Equipment struct {
	Name   string `json:"name"`
	Bait   string `json:"bait"`
	Leader string `json:"leader"`
}
