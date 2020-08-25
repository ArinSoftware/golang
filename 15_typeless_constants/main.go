package main

import "fmt"

func main() {

	/* 	const x = 5

	   	fmt.Printf("%T, %v\n", x, x) */

	/* 	const x = 3
	   	var y int16 = 12

	   	fmt.Printf("%T, %v\n", x, x)

	   	fmt.Printf("%T, %v\n", y, y)

	   	fmt.Printf("%T, %v\n", x+y, x+y) // int16(x) + y

		   fmt.Printf("%T, %v\n", x, x) */

	/* 	const x = int16(5.2 + 4.8)

	   	fmt.Printf("%T, %v\n", x, x) */

	/* 	const x = 3
	   	const y = 5.6

	   	fmt.Printf("%T, %v\n", x, x)

	   	fmt.Printf("%T, %v\n", y, y)

	   	fmt.Printf("%T, %v\n", x+y, x+y) */

	const x = 3
	const y = 5.6
	const z = true
	const t = "test"

	fmt.Printf("%T, %v\n", x, x)
	fmt.Printf("%T, %v\n", y, y)
	fmt.Printf("%T, %v\n", z, z)
	fmt.Printf("%T, %v\n", t, t)

}
