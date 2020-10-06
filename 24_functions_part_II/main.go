/* package main

import "fmt"

func main() {

	merhaba("Arin", 6) // Argument Function Run

}

func merhaba(name string, age int) { // Parametre Function Write
	fmt.Printf("Adım %s, yaşım %d", name, age)
} */

/* package main

import "fmt"

func main() {

	fmt.Println(result(40))

}

func result(grade int) string {
	if grade >= 50 {
		return "geçtiniz"

	}
	return "kaldınız"
	fmt.Println("FONKSIYON ICINDEYIM")

} */

/* package main

import "fmt"

func main() {

	merhaba("Arin", 6)

	name := "Elis"
	age := 4
	fmt.Printf("Adım %s, yaşım %d", name, age)
}

func merhaba(name string, age int) {

	fmt.Printf("Adım %s, yaşım %d\n", name, age)

} */

/* package main

import (
	"fmt"
	"time"
)

func main() {

	var x int = 10

	fmt.Println(x)

	var moment time.Time = time.Now() // Now ---> method

	fmt.Println(moment)
} */

/* package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	fmt.Print("Lütfen Sınav Sonucunuzu Giriniz: ")
	reader := bufio.NewReader(os.Stdin)
	value, _ := reader.ReadString('\n') // _ blank identifier
	fmt.Println(value)
} */

package main

import "fmt"

func main() {

	bölüm, kalan := bölme(104, 5)
	fmt.Println(bölüm, kalan)

}

// 104 / 5 =====> 20 - 4

func bölme(bölünen, bölen int) (bölüm, kalan int) {
	bölüm = bölünen / bölen
	kalan = bölünen % bölen

	return bölüm, kalan
}
