package main

import "fmt"

// Crie uma estrutura (struct) chamada Address com os seguintes campos:
// street (string)
// city (string)
// state (string)
// zipCode (string)

// Crie uma segunda estrutura (struct) chamada Person com os seguintes campos:

// name (string)
// age (int)
// address (Address)

// Em seguida, crie uma variável do tipo Person e a inicialize com alguns valores de exemplo, incluindo um valor do tipo Address. Por fim, imprima os valores dos campos usando o pacote fmt.

type Address struct {
	street  string
	city    string
	state   string
	zipCode string
}

type Person struct {
	name    string
	age     int
	address Address
}

func main() {

	p := Person{
		name: "Kleber",
		age:  42,
		address: Address{
			street:  "Rua dos bobos",
			city:    "Caracas",
			state:   "MG",
			zipCode: "37090-000",
		},
	}
	fmt.Println(p)
}
