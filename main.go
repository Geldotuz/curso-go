package main

import (
	"encoding/json"
	"fmt"
)

type Persona struct {
	Name string
	Age  int
	Eye  string
}

func main() {
	p1 := Persona{
		Name: "Angel",
		Age:  30,
		Eye:  "Brown",
	}
	p2 := Persona{
		Name: "Marin",
		Age:  24,
		Eye:  "Green",
	}

	Personas := []Persona{p1, p2}

	fmt.Println(Personas)

	bs, err := json.Marshal(Personas)
	if err != nil {
		panic(err.Error())
	}

	fmt.Println(string(bs))
}
