package main

import "fmt"

func main() {

	/* 	var x, y, sum int

	   	x = 5
	   	y = 10
	   	sum = x + y
	   	fmt.Printf("%d ve %d toplamı %d\n", x, y, sum)

	   	x = 7
	   	y = 11
	   	sum = x + y
	   	fmt.Printf("%d ve %d toplamı %d\n", x, y, sum) */

	// Moduler Programming
	// Readable code
	// From Complex To Simple

	/* 	fmt.Println(sum(5, 10)) */

	// func funcName(params) return type { code  }

	/* 	merhaba()

	   	fmt.Println(sum(5, 10))
	   	fmt.Println(sum(3, 5))
	   	fmt.Println(sum(2, 7))

	   	merhaba()
		   merhaba() */

	// fmt.Println(x)

	// return vs print

	/* 	z := sum(5, 10)
	   	fmt.Println(z)

		   sum2(6, 11) */

	merhaba2("Elis", 4)

	// Fonksiyon İsimlendirme
	// ilk karekter harf
	// camel Case -- mySum, myBestFunction
	// paket dışı --> ilk harf büyük

}

func sum(x, y int) int {
	return x + y
}

func sum2(x, y int) {
	fmt.Println(x + y)
}

func merhaba() {
	fmt.Println("Benim Adım Arin")
}

func merhaba2(name string, age int) {
	fmt.Printf("Adım %s, yaşım %d", name, age)
}
