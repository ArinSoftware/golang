/*

--- BASIC TYPES (BUILT-IN TYPES)
NUMERIC TYPES
uint8       the set of all unsigned  8-bit integers (0 to 255) = 256
uint16      the set of all unsigned 16-bit integers (0 to 65535)
uint32      the set of all unsigned 32-bit integers (0 to 4294967295)
uint64      the set of all unsigned 64-bit integers (0 to 18446744073709551615)

int8        the set of all signed  8-bit integers (-128 to 127) = 256
int16       the set of all signed 16-bit integers (-32768 to 32767)
int32       the set of all signed 32-bit integers (-2147483648 to 2147483647)
int64       the set of all signed 64-bit integers (-9223372036854775808 to 9223372036854775807)

float32     the set of all IEEE-754 32-bit floating-point numbers
float64     the set of all IEEE-754 64-bit floating-point numbers

complex64   the set of all complex numbers with float32 real and imaginary parts
complex128  the set of all complex numbers with float64 real and imaginary parts

byte        alias for uint8
rune        alias for int32


STRING TYPES
A string type represents the set of string values. A string value is a (possibly empty) sequence of bytes.
The number of bytes is called the length of the string and is never negative.

BOOLEAN TYPES
A boolean type represents the set of Boolean truth values denoted by the predeclared constants true and false.


--- COMPOSITE TYPES
Slice - Struct - Pointer - Function - Interface - Map - Channel - Array
*/

package main

import "fmt"

func main() {

	/* 	x := 10
	   	y := 10.0

	   	fmt.Printf("%v %T\n", x, x)
	   	fmt.Printf("%v %T\n", y, y)

	   	// Type Conversion type(value) => int(y) = 10.0 => 10

	   	fmt.Println(x + int(y))

	   	fmt.Printf("%v %T\n", y, y) */

	/* 	var x int8 = 10
	   	var y int16 = 10

	   	fmt.Println(x + y) */

	/* 	x := 10
	   	y := 10.0

	   	fmt.Printf("%v %T\n", x, x)
	   	fmt.Printf("%v %T\n", y, y)

	   	fmt.Println(float64(x) + y) */

	/* 	var x int16 = 128
	   	var y int8

	   	y = int8(x) // type(value)

	   	fmt.Println(y) */

	/* 	x := 10
	   	y := "10"

	   	fmt.Printf("%v %T\n", x, x)
	   	fmt.Printf("%v %T\n", y, y)

		   fmt.Println(x + int(y)) */

	num1 := 106
	str1 := string(num1)

	fmt.Printf("%v %T\n", num1, num1)
	fmt.Println()
	fmt.Printf("%v %T\n", str1, str1)

}
