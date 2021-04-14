package main

import (
	"fmt"
	"github.com/basant93/goexamples/standard/lib/gotypes/type/types"
)

func main() {
	per := NewPerson("konark", 25)
	fmt.Printf("person firstname: %s \n", per.FirstName)
	fmt.Println(per.details())

	home := types.Home{}
	home.SetHomeDetails(100, "India Pincode - 411000")
	fmt.Println(home.GetHomeDetails())

}

type person struct {
	FirstName string
	Age       int
}

func NewPerson(firstname string, age int) *person {
	return &person{FirstName: firstname, Age: age}

}

func (p *person) details() string {
	return fmt.Sprintf("%s has age %d", p.FirstName, p.Age)
}

//keep all methods as pointer based receiver, so that when we implement interface it would br consistent and struct would implement the method
//struct is not copied here and it allows for update of struct
func (p *person) setFirstName(name string) {
	p.FirstName = name
}
