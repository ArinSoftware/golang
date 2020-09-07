package main

import "fmt"

func main() {

	// if <boolean expression> { code }  x%2 == 0 (false)

	/* 	x := 27

	   	if false {
	   		fmt.Println(x, "tek sayıdır")
	   	} */

	/* 	if !true {
		fmt.Println("Mesaj Gosterilecek")
	} */

	/* 	if 5 > 3 {
		fmt.Println("Mesaj Gosterilecek mi?")
	} */

	/* 	x := 4

	   	if x%2 == 0 {
	   		fmt.Println(x, "çift sayıdır")
	   	} else {
	   		fmt.Println(x, "tek sayıdır")
		   } */

	// if <boolean expression> { code } else { code }

	/* 	if false {
		fmt.Println("Mesaj Gosterilecek")
	} */

	x := -5

	if x < 0 {
		fmt.Println(x, "negatif sayıdır")
	} else if x%2 == 0 {
		fmt.Println(x, "çift sayıdır")
	} else {
		fmt.Println(x, "tek sayıdır")
	}

	// if <boolean expression> { code } else if <boolean expression> else { code }
}
