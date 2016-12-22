package main

import (
	"fmt"
	"strings"
)

type Human struct {
	Name  string
	Phone string
	Hobby []string
}

type Animal struct {
	Name string
	Type string
}

type Mortal interface {
	Talk() string
	Run() string
	LearningRate() float64
}

func (a Animal) Talk() string {
	return fmt.Sprintf("hey my name is %s and my type is %s", a.Name, a.Type)
}

func (a Animal) Run() string {
	return "i run very fast"
}

func (a Animal) LearningRate() float64 {
	return 0.01
}

func (h Human) Talk() string {
	return fmt.Sprintf("hey my name is %s and my phone number is %s and my hobbys are %s", h.Name, h.Phone, strings.Join(h.Hobby, ","))
}

func (h Human) Run() string {
	return "i run very slow"
}

func (h Human) LearningRate() float64 {
	return 99.9
}
