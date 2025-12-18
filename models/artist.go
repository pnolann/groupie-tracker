package models

type Artists struct {
	Id          int      `json:"id"`
	Name        string   `json:"name"`
	Image       string   `json:"image"`
	Members     []string `json:"members"`
	FirstAlbum  string   `json:"firstAlbum"`
	CreatedDate int      `json:"creationDate"`

	Locations    string `json:"locations"`
	ConcertDates string `json:"concertDates"`
	Relations    string `json:"relations"`
}
