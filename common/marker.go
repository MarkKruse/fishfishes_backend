package common

type Marker struct {
	Id          string      `json:"markertId"`
	Title       string      `json:"title"`
	Coordinates Coordinates `json:"coordinates"`
}

type Coordinates struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}
