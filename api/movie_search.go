package api

import "time"

type MovieFilter struct  {
	Title    string `json:"title"`
	Genre    string `json:"genre"`
	Director string `json:"director"`
}

type Movie struct {
	Title 		string 	  `json:"title"`
	Cast 		string 	  `json:"cast"`
	ReleaseDate time.Time `json:"release_date"`
	Genre 		string 	  `json:"genre"`
	Director 	string    `json:"director"`
}

type MovieSearch interface {
	Search(filter MovieFilter) ([]Movie, error)
}


type MovieService struct {

}

func (s *MovieService) Search(filter MovieFilter) ([]Movie, error){
	m1 := Movie{
		Title := "Ta cabron el himalaya",
		Cast := "Zaramay",
		ReleaseDate: time.Now(),
		Genre: "Cs Ficcion",
		Director: "",
	}
	m2 := Movie{
		Title := "Bizarap Music Session",
		Cast := "Zaramay",
		ReleaseDate: time.Now(),
		Genre: "Cs Ficcion",
		Director: "Bizarap",
	}

	var _movies []Movie

	_movies = append(_movies, m1)
	_movies = append(_movies, m2)

	return _movies, nil
}
