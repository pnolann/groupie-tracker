package models

type Artists struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Image string `json:"image"`
	CreatedDate int `json:"createdDate"`
	Members []string `json:"members"`
	FirstAlbum int `json:"firstAlbum"`

	Location string `json:"location"`
	ConcertDate string `json:"concertDate"`
	relation string `json:"relation"`

}