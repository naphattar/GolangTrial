package models

type Camp struct {
	CampID   string `json:"campid"`
	Name     string `json:"name"`
	Time     string `json:"time"`
	Location string `json:"location"`
	Director string `json:"director"`
}
