/* package main

import "fmt"

func main() {

	var employee struct { // structure
		name      string
		age       int
		isMarried bool
	}

	//fmt.Println(employee)
	//fmt.Printf("%#v\n", employee)
	//fmt.Println(employee.age)

	employee.name = "Gurcan"
	employee.age = 40
	employee.isMarried = true

	fmt.Printf("%#v\n", employee)
	fmt.Println(employee.name)
	fmt.Println(employee.age)
	fmt.Println(employee.isMarried)

	var employee2 struct { // structure
		name      string
		age       int
		isMarried bool
	}

	employee2.name = "Arin"
	employee2.age = 5
	employee2.isMarried = false

	fmt.Printf("%#v\n", employee2)
	fmt.Println(employee2.name)
	fmt.Println(employee2.age)
	fmt.Println(employee2.isMarried)

} */

/* package main

import "fmt"

type employee struct { // underlying type
	name      string
	age       int
	isMarried bool
}

func main() {

	var e1 employee
	e1.name = "Gurcan"
	e1.age = 40
	e1.isMarried = true

	var e2 employee
	e2.name = "Arin"
	e2.age = 5
	e2.isMarried = false

	e3 := employee{
		name:      "Elis",
		age:       3,
		isMarried: false,
	}

	fmt.Printf("%#v\n", e1)
	fmt.Printf("%#v\n", e2)
	fmt.Printf("%#v\n", e3)

} */

package main

import "fmt"

type employee struct { // underlying type
	name      string
	age       int
	isMarried bool
	kids      []string
}

func main() {

	e1 := employee{
		name:      "Gurcan",
		age:       40,
		isMarried: true,
		kids: []string{
			"Arin",
			"Elis",
		},
	}

	//fmt.Printf("%#v\n", e1)
	fmt.Println(e1)
	fmt.Println(e1.kids)
	fmt.Println(e1.kids[0])

}
