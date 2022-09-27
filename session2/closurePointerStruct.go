package session2

import "fmt"

type Person struct {
	name string
}

func main3() {

	// names := []*Person{
	// 	{"Giva"},
	// 	{"Iqbal"},
	// 	{"Ucup"},
	// 	{"Nugraha"},
	// 	{"Bayu"},
	// 	{"Putu"},
	// 	{"Satrio"},
	// 	{"Agus"},
	// 	{"Barru"},
	// 	{"Hasan"},
	// }

	// printPerson := func(persons []*Person) {
	// 	for _, person := range persons {
	// 		fmt.Println(*&person.name)
	// 	}
	// }

	names := []*Person{
		{name: "Giva"},
		{name: "Iqbal"},
		{name: "Ucup"},
		{name: "Nugraha"},
		{name: "Bayu"},
		{name: "Putu"},
		{name: "Satrio"},
		{name: "Agus"},
		{name: "Barru"},
		{name: "Hasan"},
	}

	printPerson := func(persons []*Person) {
		for _, person := range persons {
			fmt.Println(person.name)
		}
	}

	printPerson(names)

}
