package main

import (
	"fmt"
)

type Animal interface {
	GetWalkSpeed() int
	GetVoice() string
}

type Wolf struct {
	Weight int
}

type Snake struct {
	Poison bool
	Tall   int
}

func (w Wolf) GetWalkSpeed() int {
	return 10
}

func (s Snake) GetWalkSpeed() int {
	return 4
}

func (w Wolf) GetVoice() string {
	return "Wofff"
}

func (s Snake) GetVoice() string {
	return "Tsahh"
}

func (s Snake) GetPoisonLevel() int {
	return 8
}

func GetAnimal(option string) Animal {
	if option == "mammals" {
		return Wolf{}
	}
	return Snake{}
}

func AbstractExample() {
	w := GetAnimal("mammals")
	// w := GetAnimal("reptil")

	fmt.Println(w.GetWalkSpeed())
	fmt.Println(w.GetVoice())
}
