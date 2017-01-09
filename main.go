package main

import (
	"fmt"

	"github.com/hamdimuzakkiy/golang/src/regex"
)

func main() {
	s := ReverseString("test")
	fmt.Println(s)

	tigger := Animal{
		Name: "Tigger",
		Type: "Mamal",
	}

	person := Human{
		Name:  "hamdi",
		Phone: "0809894124",
		Hobby: []string{"swimming", "gym", "play"},
	}

	var livingThings []Mortal

	livingThings = append(livingThings, tigger)
	livingThings = append(livingThings, person)

	for _, living := range livingThings {
		fmt.Println(living.Talk())
		fmt.Println(living.Run())
	}

	regex.Plus(1, 2, 3, 4)
}
