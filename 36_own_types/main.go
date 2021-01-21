/* package main

import (
	"fmt"
	"strings"
)

// struct --> underlying type, employee ---> Defined Type, Named Type

type mile float64
type kilometer float64

//type mystring string

func main() {

	/* 	f1 := float64(4.4)

	fmt.Println()
	fmt.Printf("%T, %v", (m1 + mile(f1)), (m1 + mile(f1)))
	fmt.Println()
	fmt.Printf("%T, %v", (float64(m1) + f1), (float64(m1) + f1)) */

/* 	var m1 mile
   	m1 = 3.2
   	//fmt.Println(m1)
   	//fmt.Printf("%T, %v", m1, m1)

   	//fmt.Println()

   	m2 := mile(4.6) */

/* 	k1 := kilometer(7.8)
   	fmt.Printf("%T, %v", k1, 11)

	   fmt.Println(m1 + k1) */

/* 	fmt.Println(m1 + m2)
   	fmt.Printf("%.2f", (m1 * m2))
   	fmt.Println()
   	fmt.Println(m1 + m2 + 2.1)
   	fmt.Println()
   	//fmt.Println(m1 + m2 + "arin") */

/* 	s1 := "arin"

	//m1 := mile(4.6)

	fmt.Println(strings.ToUpper(s1))
	//fmt.Println(strings.ToUpper(m1))


} */

package main

import "fmt"

type mile float64
type kilometer float64

func main() {

	// m1 = 10 k1=?
	m1 := mile(10)

	k1 := toKilometer(m1)

	fmt.Println(k1)

	// k2 = 10 m2=?

	k2 := kilometer(10)

	m2 := toMile(k2)
	fmt.Println(m2)

}

func toKilometer(m mile) kilometer {
	return kilometer(m * 1.6)
}

func toMile(k kilometer) mile {
	return mile(k * 0.62)
}
