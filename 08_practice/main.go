// 1-) studentName --> John Doe, grade --> 77, isPassed --> true değişkenlerini
// 3 farklı yöntem ile oluşturup, çıktısını yazdırınız.

/* package main

import "fmt"

func main() { */

/* 	var studentName string = "John Doe"
   	var grade int = 77
   	var isPassed bool = true */

/* 	var studentName = "John Doe"
   	var grade = 77
	   var isPassed = true */

/* 	studentName := "John Doe"
	grade := 77
	isPassed := true

	fmt.Println(studentName)
	fmt.Println(grade)
	fmt.Println(isPassed)
}

*/

// 2-) yukarıda belirtilen değişkenleri tek satır içerisinde tanımlayınız.
/*
package main

import "fmt"

func main() {

	/* 	var studentName string = "John Doe"
	   	var grade int = 77
	   	var isPassed bool = true */

/* var studentName, grade, isPassed = "John Doe", 77, true */

/* 	studentName, grade, isPassed := "John Doe", 77, true

	fmt.Println(studentName)
	fmt.Println(grade)
	fmt.Println(isPassed)
}  */

// 3-) "Declaration", "Assign", "Initialization", "Initial Value" kavramlarını açıklayınız. (Terminoloji)

/* package main

import "fmt"

func main() {

	var studentName string = "John Doe"

	studentName = "Mahmut Erdem"

	fmt.Println(studentName)

} */

// 4-) "Statically Typed" vs "Dynamically Typed" ifadelerini GO ve Python üzerinden gösteriniz.

// 5-) ":=" vs "=" aradaki farkı gösteriniz, double declaration

package main

import "fmt"

func main() {

	/* 	var studentName string = "John Doe"
	   	studentName = "Mahmut Erdem" */

	studentName := "John Doe"
	studentName = "Mahmut Erdem"

	fmt.Println(studentName)

}
