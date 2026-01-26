package models

type Artists struct {
	Id          int      `json:"id"`
	Name        string   `json:"name"`
	Image       string   `json:"image"`
	CreatedDate int      `json:"creationDate"`
	Members     []string `json:"members"`
	FirstAlbum  string   `json:"firstAlbum"`

	Locations    string `json:"locations"`
	ConcertDates string `json:"concertDates"`

	Relations string `json:"relations"`

	RelationsData Relations `json:"-"`
}
