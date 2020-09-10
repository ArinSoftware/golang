package main

import "fmt"

func main() {

	/* 	x := 10

	   	if x := -5; x < 0 {
	   		fmt.Println(x, "negatif sayıdır")
	   	} else if x%2 == 0 {
	   		fmt.Println(x, "çift sayıdır")
	   	} else {
	   		fmt.Println(x, "tek sayıdır")
	   	}

		   fmt.Println(x) */

	if x := -25; x < 0 {
		fmt.Println(x, "negatif sayıdır")
		fmt.Println("Benim Adım Arin")
	} else {
		if x%2 == 0 {
			fmt.Println(x, "çift sayıdır")
		} else {
			fmt.Println(x, "tek sayıdır")
		}
	}

}
