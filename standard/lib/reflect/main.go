//created a new copy of struct with new title. In case of long struct it can be problematic
package main

import (
	"fmt"
	"github.com/basant93/goexamples/standard/lib/reflectauxiliary"
)

func main() {

	movie := reflectauxiliary.NewMovie("Five Feet Apart", 8.0)
	fmt.Println(movie.GetMovieName())
	fmt.Println(movie.GetMovieRating())

	movie.SetMovieName("Schinder's list")
	movie.SetMovieRating(9)

	fmt.Println(movie.GetMovieName())
	fmt.Println(movie.GetMovieRating())
}
