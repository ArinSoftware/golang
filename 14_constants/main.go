package main

import "fmt"

func main() {

	/* 5
	   3.14
	   "passed"
	   true */

	/* 	r := 3.0

	   	const pi float64 = 3.14

	   	areaOfCircle := pi * (math.Pow(r, 2))

	   	fmt.Println(areaOfCircle) */

	/* 	const x int = 2
	   	const y float32 = 3.4
	   	const z string = "test"
	   	const t bool = false

	   	fmt.Printf("%T, %v\n", x, x)
	   	fmt.Printf("%T, %v\n", y, y)
	   	fmt.Printf("%T, %v\n", z, z)
		   fmt.Printf("%T, %v\n", t, t) */

	/* 	const x = 2

	   	//x = 5
	   	// x++
	   	// x = x + 1

		   fmt.Printf("%T, %v\n", x, x) */

	// const ---> compile time
	// var ---> run time

	// var x = math.Pow(3, 4)
	// const x = math.Pow(3, 4)

	/* 	const x = 5

	   	fmt.Printf("%T, %v\n", x, x) */

	/* 	y := 3

	   	const x = y

	   	fmt.Printf("%T, %v\n", y, y)
		   fmt.Printf("%T, %v\n", x, x) */

	/* 	const x = 1
	   	var y = 3

	   	fmt.Printf("%T, %v\n", x, x)
	   	fmt.Printf("%T, %v\n", y, y)
	   	fmt.Printf("%T, %v\n", x+y, x+y)
	*/

	const x, y = 3, 5

	fmt.Printf("%T, %v\n", x, x)
	fmt.Printf("%T, %v\n", y, y)
}
