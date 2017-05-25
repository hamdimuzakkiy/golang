package main

import (
	"fmt"

	"github.com/hamdimuzakkiy/mocktemplate/sqlxmock"
)

func sqlxlMockExample() {
	s := sqlxmock.DBTemplate{}
	fmt.Println(s)
}
