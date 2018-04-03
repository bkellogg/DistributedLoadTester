package main

import (
	"fmt"

	"github.com/Pallinder/go-randomdata"
)

func main() {
	fmt.Println("Hello World!")
	fmt.Printf("My name is %s %s! I live in %s, %s. This is a random number: %d\n",
		randomdata.FirstName(randomdata.RandomGender),
		randomdata.LastName(),
		randomdata.City(),
		randomdata.State(randomdata.Large),
		randomdata.Number(0, 100))
}
