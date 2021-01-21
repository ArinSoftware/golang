// 1 -) Underlying Type 'int' olacak şekilde kendi veri tipinizi oluşturunuz
// veri tipi ve değerini '10' yazdırınız.

/* package main

import "fmt"

type myType int

func main() {

	var x myType
	x = 10
	fmt.Printf("%T, %v", x, x)

} */

// 2 -) Başlangıç değeri 10 olan bir X değişkeni oluşturun sonrasında
// bulunduğu adres üzerinden y değişkenini tanımlayıp x değerini 20 yapınız.

/* package main

import "fmt"

func main() {

	x := 10
	fmt.Println(x)
	y := (&x)
	fmt.Println(y)
	*y = 20
	fmt.Println((*y))
	fmt.Println(x)
} */

// 3 -) Underlying Type struct olan Rectangle type oluşturunuz.
// display, area, circumference metodlarını yazınız.

/* package main

import "fmt"

type rectangle struct {
	a, b int
}

func (r rectangle) display() {
	fmt.Println("Kenar uzunlukları: ", r.a, " ve ", r.b, " olan dikdörtgen ")
}

func (r rectangle) area() int {
	return r.a * r.b
}

func (r rectangle) circumference() int {
	return 2 * (r.a + r.b)
}

func main() {

	r1 := rectangle{3, 8}
	r1.display()
	fmt.Println("Alanı: ", r1.area())
	fmt.Println("Çevresi: ", r1.circumference())

} */

// 4-) family aile bireyleri şeklinde veri tipi oluşturalım, underlying struct. Aile bireylerinin hepsini farklı
// şekilde tanımlayınız. Sonrasında for döngüsünde yazdırınız. field age, name, isMarried field.

package main

import "fmt"

type family struct {
	name      string
	age       int
	isMarried bool
}

func showFamily() []family {

	family1 := family{
		name:      "Arin",
		age:       5,
		isMarried: false,
	}

	family2 := family{
		name: "Elis",
		age:  3,
	}

	family3 := family{
		"Gurcan",
		40,
		true,
	}

	var family4 family
	family4.name = "Gamze"
	family4.age = 40
	family4.isMarried = true

	return []family{family1, family2, family3, family4}
}

func main() {
	families := showFamily()
	for i := 0; i < len(families); i++ {
		fmt.Printf("%v, %T\n", families[i], families[i])
	}
}
