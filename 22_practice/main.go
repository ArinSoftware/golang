// 1-) 1 ile 10 arasındaki sayıları if yapısıyla tek - çift olarak yazdırınız.

/* package main

import "fmt"

func main() {

	for i := 1; i <= 10; i++ {

		if i%2 == 0 {
			fmt.Println(i, "çittir")
		} else {
			fmt.Println(i, "tektir")
		}
	}

} */

// 2-) for yapısını kullanarak Go'da olmayan while döngüsüne örnek veriniz.

/* package main

import "fmt"

func main() {

	x := 0

	for x < 10 {
		fmt.Println(x)
		x++
	}

} */

// 3-) Switch fallthrough ifadesini açıklayınız.

/* package main

import "fmt"

func main() {

	switch x := 75; {
	case x < 20:
		fmt.Printf("%d küçüktür 20\n", x)
		fallthrough

	case x < 50:
		fmt.Printf("%d küçüktür 50\n", x)
		fallthrough

	case x < 100:
		fmt.Printf("%d küçüktür 100\n", x)
		fallthrough

	case x < 200:
		fmt.Printf("%d küçüktür 200\n", x)
	}

} */

// 4-) Aşağıdaki if döngüsünü daha idiomatic hale getiriniz.

/* package main

import (
	"fmt"
)

func main() {
	if x := 20; x%2 == 0 {
		fmt.Println(x, "çifttir")
	} else {
		fmt.Println(x, "tektir")
	}
} */

/* package main

import "fmt"

func main() {

	x := 20

	if x%2 == 0 {
		fmt.Println(x, "çifttir")
		return
	}

	fmt.Println(x, "tektir")
} */

// 5-) 1 ile 50 arasındaki asal sayıları gösteren bir program yazınız.

package main

import "fmt"

func main() {

	var x, y int

	for x = 2; x < 50; x++ {
		for y = 2; y < (x / y); y++ {
			if x%y == 0 {
				break
			}
		}

		if y > (x / y) {
			fmt.Printf("%d bir asal sayıdır\n", x)
		}
	}
}
