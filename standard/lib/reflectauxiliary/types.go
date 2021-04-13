package reflectauxiliary

import (
	"strings"
)

type Movie struct {
	Name   string
	Rating float32
}

func NewMovie(name string, rating float32) Movie {
	return Movie{
		Name:   name,
		Rating: rating,
	}
}

func (m *Movie) GetMovieName() string {
	return strings.ToUpper(m.Name)
}

func (m *Movie) GetMovieRating() float32 {
	return m.Rating
}

func (m *Movie) SetMovieName(name string) {
	m.Name = name
}

func (m *Movie) SetMovieRating(rating float32) {
	m.Rating = rating
}
