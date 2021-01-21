package main

import "fmt"

type Person struct {
	Surname string
}

func (person Person) getSurname() string {
	return person.Surname
}

type Kid struct {
	FavoriteColor string
	Person
}

func (kid Kid) getFavoriteColor() string {
	return kid.FavoriteColor
}

func main() {
	person1 := Person{
		"Sheikh",
	}
	kid1 := Kid{
		"Blue",
		person1,
	}
	fmt.Println("Kid's Surname: ", kid1.getSurname())
	fmt.Println("Kid's Favorite Color: ", kid1.getFavoriteColor())
}
