package main

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

type DbInterface interface {
	DriverName() string
}

type DbInterfaceGet interface {
	Get(dest interface{}, query string, args ...interface{}) error
}

type Person struct {
	Name string `db:"name"`
	Age  int    `db:"age"`
	Sex  int    `db:"sex"`
}

func GetPerson(d DbInterfaceGet) Person {
	q := `	SELECT 
				name,
				age
			FROM
				PERSON
			LIMIT 1
	`
	var p Person
	d.Get(&p, q)
	p.Age *= 10
	return p
}

func GetDriverName(d DbInterface) string {
	return fmt.Sprintf("hi, driver name is %s", d.DriverName())
}

func DoSomethingWithDriverName() {
	s := sqlx.DB{}
	GetDriverName(&s)
}

func DoSomethingWithInsert() {
	s := sqlx.DB{}
	GetPerson(&s)
}
