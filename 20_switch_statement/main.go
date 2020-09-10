/* 	if grade == 5 {
   		fmt.Println("Pekiyi")
   	} else if grade == 4 {
   		fmt.Println("İyi")
   	} else if grade == 3 {
   		fmt.Println("Orta")
   	} else if grade == 2 {
   		fmt.Println("Geçer")
   	} else if grade == 1 {
   		fmt.Println("Başarısız")
   	} else {
   		fmt.Println("Geçersiz Not")
   	} */

package main

import "fmt"

func main() {

	// Switch
	/*
		grade := -7

		switch grade {
		/* 	default:
		fmt.Println("Geçersiz Not") */

	/*	case 5: // if grade == 5 { fmt.Println("Pekiyi")  }
			fmt.Println("Pekiyi")

		case 4:
			fmt.Println("İyi")

		case 3:
			fmt.Println("Orta")

		case 2:
			fmt.Println("Geçer")
			/* 		y := 100
			   		fmt.Println(y) */
	/*	case 1:
		fmt.Println("Başarısız")

	} */

	switch {
	case false:
		fmt.Println("Bu yazdığımız konsolda görünmez.")

	case true:
		fmt.Println("Bu yazdığımız konsolda görünecek.")

	}

}
