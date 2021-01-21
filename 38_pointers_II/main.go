/* package main

import "fmt"

func main() {       // GO pass by value
	x := 5
	fmt.Println(x)
	double(x)
	fmt.Println(x)
}

func double(num int) {
	num *= 2
	fmt.Println(num)
} */

/* package main

import "fmt"

func main() {

	mySlc := []int{1, 10, 100}

	fmt.Println(mySlc)

	double(mySlc)

	fmt.Println(mySlc)
}

func double(slc []int) {
	for i := 0; i < len(slc); i++ {
		slc[i] *= 2
	}

	fmt.Println(slc)
} */

/* package main

import "fmt"

func main() {

	myArr := [3]int{1, 10, 100}

	fmt.Println(myArr)

	double(myArr)

	fmt.Println(myArr)
}

func double(arr [3]int) {
	for i := 0; i < len(arr); i++ {
		arr[i] *= 2
	}

	fmt.Println(arr)
} */

package main

import "fmt"

func main() { // GO pass by value
	x := 5
	fmt.Println(x)
	double(&x)
	fmt.Println(x)

}
func double(num *int) { // double --> pointer method
	fmt.Println(num)
	*num *= 2
	fmt.Println(*num)
}
