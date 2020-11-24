package main

import "fmt"

func main() {

	/* 	myArr := [3]int{1, 2, 3}
	   	myArr2 := [...]int{1, 2, 3, 4}
	   	fmt.Println(myArr)
	   	fmt.Println(myArr2)
	   	fmt.Println(len(myArr))
	   	fmt.Println(len(myArr2))

	   	var myArr3 [5]int
	   	fmt.Println(myArr3) */

	/* 	mySlc := []int{1, 2, 3} // LITERAL
	   	fmt.Println(mySlc)
	   	fmt.Println(len(mySlc)) */

	/* 	var myArr [4]int
	   	fmt.Println(myArr)
	   	myArr[0] = 5
	   	fmt.Println(myArr) */

	/* 	var mySlc []int
	   	mySlc = make([]int, 4) // MAKE METHOD
	   	fmt.Println(mySlc)
	   	mySlc[0] = 10
	   	fmt.Println(mySlc) */

	/* 	myArr := [3]int{1, 2, 3}
	   	fmt.Println(myArr)
	   	myArr2 := myArr
	   	fmt.Println(myArr2)

	   	myArr2[0] = 100
	   	fmt.Println(myArr2)
		   fmt.Println(myArr)   Pass By Value    */

	mySlc := []int{1, 2, 3}
	fmt.Println(mySlc)
	mySlc2 := mySlc
	fmt.Println(mySlc2)
	mySlc2[0] = 33
	fmt.Println(mySlc2)
	fmt.Println(mySlc) // Pass By Reference

}
