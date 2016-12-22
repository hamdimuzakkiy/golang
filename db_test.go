package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockSqlx struct {
	mock.Mock
}

func (m MockSqlx) DriverName() string {
	args := m.Called()
	return args.String(0)
}

func (m MockSqlx) Get(p interface{}, query string, argsInput ...interface{}) error {
	args := m.Called(p, query)
	p.(*Person).Name = "hamdi ahmadi muzakkiy"
	p.(*Person).Age = 22

	return args.Error(0)
}

func TestGetDriverName(t *testing.T) {
	m := new(MockSqlx)
	m.On("DriverName").Return("something")

	assert.Equal(t, "hi, driver name is something", GetDriverName(m), "TestGetDriverName")
}

func TestGetPerson(t *testing.T) {
	m := new(MockSqlx)
	var p Person
	q := `	SELECT 
				name,
				age
			FROM
				PERSON
			LIMIT 1
	`
	m.On("Get", &p, q).Return(nil)

	assert.Equal(t, Person{
		Name: "hamdi ahmadi muzakkiy",
		Age:  220,
	}, GetPerson(m), "TestGetPerson")
}
