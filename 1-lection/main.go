package main

import (
	"encoding/json"
	"fmt"
	"github.com/shopspring/decimal"
	"log"

	"github.com/stud-it-department/lections/1-lection/structures"
)

func main() {
	// стандатные типы данных
	var varInteger int
	var varString string
	var varFloat64 float64

	fmt.Printf("int: <%d>, string: <%s>, float: <%v>\n", varInteger, varString, varFloat64)

	integer := 2
	str := "common-string"
	fl := 2.2

	fmt.Printf("int: <%d>, string: <%s>, float: <%v>\n", integer, str, fl)

	// ...
	structure := structures.BaseStruct{}

	structure.SetString("github")
	fmt.Printf("структура с методом от значения: %+v\n", structure)

	structure.SetInteger(1)
	fmt.Printf("структура с методом от объекта: %+v\n", structure)

	// make(type, len, cap)
	slice := make([]int, 0, 0)
	fmt.Printf("slice init: %+v\n", slice)
	slice = append(slice, 1)
	fmt.Printf("slice after append: %+v\n", slice)

	var anotherSlice []int
	fmt.Printf("slice init: %+v\n", anotherSlice)
	anotherSlice = append(anotherSlice, slice...)
	fmt.Printf("slice after append: %+v\n", anotherSlice)

	d := decimal.NewFromFloat(125.1)
	f := 125.1
	jsonStruct := &structures.JSONStruct{
		Integer:          integer,
		DifficultInteger: integer,
		D:                &d,
		F:                &f,
	}
	bytes, err := json.Marshal(jsonStruct)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("структура после маршалинга: %s\n", bytes)
}
