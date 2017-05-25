package main

import (
	"fmt"
)

type CarInterface interface {
	GetMaxSpeed() int
	GetCarSound() string
	GetWheel() int
}

type Car struct {
	Price int64
}

type SportCar struct {
	Car
	Brand string
}

type Truck struct {
	Car
	types string
}

func (c Car) GetMaxSpeed() int {
	return 120
}

func (c Car) GetCarSound() string {
	return "Ngeeeng"
}

func (c Car) GetWheel() int {
	return 4
}

func (s SportCar) GetMaxSpeed() int {
	return s.Car.GetMaxSpeed() * 2
}

func (s SportCar) GetCarSound() string {
	return "Wuuuuzhh"
}

func (t Truck) GetWheel() int {
	return 6
}

func GetCar(option string) CarInterface {
	if option == "sport" {
		return SportCar{Car: Car{Price: 1000000000}, Brand: "xyz"}
	}
	return Truck{Car: Car{Price: 200000000}, types: "fruit"}
}

func OopExample() {
	car := GetCar("sport")
	// w := GetAnimal("reptil")

	fmt.Println(car.GetWheel())
	fmt.Println(car.GetCarSound())
	fmt.Println(car.GetMaxSpeed())
}
