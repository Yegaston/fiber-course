package api

import "time"

type MovieFilter struct  {
	Title    string `json:"title"`
	Genre    string `json:"genre"`
	Director string `json:"director"`
}

type Movie struct {
	Title string `json:"title"`
	Cast string `json:"cast"`
	ReleaseDate time.Time `json:"release_date"`
	Genre string `json:"genre"`
	Director string `json:"director"`
}

type MovieSearch interface {
	Search(filter MovieFilter) ([]Movie, error)
}

