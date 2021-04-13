//created a new copy of struct with new title. In case of long struct it can be problematic
package main

import (
	"fmt"
	"github.com/basant93/goexamples/standard/lib/reflectauxiliary"
)

func main() {

	movie := reflectauxiliary.Movie{}
	movie.Name = "Five Feet Apart"
	movie.Rating = 8.0

	fmt.Println(movie.Name)
	fmt.Println(movie.Rating)
	movie.Name = "Schilder's List"
	movie.Rating = 9.9
	fmt.Println(movie.Name)
	fmt.Println(movie.Rating)

}
