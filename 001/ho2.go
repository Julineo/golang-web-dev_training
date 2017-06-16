package main

import (
	"fmt"
)

type person struct {
	fName string
	lName string
}

type secretAgent struct {
	person
	blackSuburban bool
}

func (p person) pSpeak() {
	fmt.Println(p.fName, p.lName, "says, Beautiful morning, Tod.")
}

func (sa secretAgent) saSpeak() {
	fmt.Println(sa.fName, sa.lName, "says, Shaken, not stirred.")
}

func main() {
	vp := person {
		"John",
		"Chong",
	}

	vsa := secretAgent {
		person {
			"Stan",
			"Smith",
		},
		true,
	}

	fmt.Println(vp.fName)
	vp.pSpeak()
	fmt.Println(vsa.fName)
	vsa.saSpeak()
	vsa.person.pSpeak()
}
