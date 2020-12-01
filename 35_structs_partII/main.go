package main

import "fmt"

type employee struct { // underlying type
	name      string
	age       int
	isMarried bool
}

type manager struct {
	employee
	hasDegree bool
}

// IS A Relation --> Klasik
// HAS A Relation

func main() {

	e1 := employee{
		name:      "Gurcan",
		age:       40,
		isMarried: true,
	}

	fmt.Println(e1)

	/* 	m1 := manager{
		employee: employee{
			name:      "Ayşe",
			age:       28,
			isMarried: false,
		},
		hasDegree: true,
	} */

	m1 := manager{}
	m1.name = "Ayşe"
	m1.age = 28
	m1.isMarried = false
	m1.hasDegree = true

	fmt.Println(m1)

	// Anonim Struct

	theBoss := struct {
		name  string
		money bool
	}{name: "THE BOSS", money: true}

	fmt.Println(theBoss)

}
