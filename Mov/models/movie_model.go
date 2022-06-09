package models

type Movie struct {
	ID          int    `gorm:"primary key;autoIncrement" json:"id"`
	Title       string `json:"title"`
	URL         string `json:"url"`
	ReleaseDate string `json:"releaseDate"`
}
