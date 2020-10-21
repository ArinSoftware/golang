// 1 -) Iki rakam arasında toplama, çıkarma ve çarpma
// işleminin yapıldığı bir fonkiyon yazınız.

/* package main

import "fmt"

func main() {

	x, y := 10, 4
	sum1, dif1, prod1 := calculation(x, y)
	fmt.Println("Toplam: ", sum1)
	fmt.Println("Fark: ", dif1)
	fmt.Println("Çarpım: ", prod1)

}

func calculation(num1, num2 int) (int, int, int) {
	sum := num1 + num2
	dif := num1 - num2
	prod := num1 * num2

	return sum, dif, prod
} */

// 2 -) Kullanıcı tarafından girilen nota göre geçtiniz
// veya kaldınız geri dönüşünü yazdırınız.

/* package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	fmt.Print("Lütfen aldığınız notu giriniz: ")
	grade, _ := getGrade()

	var result string

	if grade >= 50 {
		result = "Geçtin"
	} else {
		result = "Kaldın"
	}

	fmt.Println(result)

}

func getGrade() (int, error) {
	reader := bufio.NewReader(os.Stdin)

	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}

	input = strings.TrimSpace(input)
	num, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println(err)
	}

	return num, nil
} */

// 3 -) 1 ile yüz arasındaki bir sayıyı tahmin etme uygulaması
// yazınız. Toplam tahmin hakkınız 10 olsun.

package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {

	target := numRand(1, 100)

	fmt.Println("1 ile 100 arasındaki sayıyı bulmaya çalışın")

	reader := bufio.NewReader(os.Stdin)

	for attempts := 0; attempts < 10; attempts++ {
		fmt.Println(10-attempts, "hakkınız kaldı")
		fmt.Print("Lütfen tahmininizi yapınız: ")

		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
		}

		input = strings.TrimSpace(input)
		num, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println(err)
		}

		if num > target {
			fmt.Println("Tahmininiz daha büyük, daha küçük bir sayı giriniz.")
		} else if num < target {
			fmt.Println("Tahmininiz daha küçük, daha büyük bir sayı giriniz.")
		} else {
			fmt.Println("Doğru Tahmin, hedf sayı ", target, " idi ", attempts, " seferde bulundunuz")
			break
		}

	}

}

func numRand(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max-min) + min
}

// -------------------------------------

/* package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	fmt.Print("Lütfen Sınav Sonucunuzu Giriniz: ")
	reader := bufio.NewReader(os.Stdin)
	value, _ := reader.ReadString('\n')
	fmt.Println(value)
} */

/* package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	fmt.Print("Lütfen Celcius derecesi giriniz: ")
	celcius, _ := getCelcius()

	fahrenheit := (celcius * 9 / 5) + 32

	fmt.Println(fahrenheit, "Fahrenheit derecedir.")

}

func getCelcius() (int, error) {
	reader := bufio.NewReader(os.Stdin)

	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}

	input = strings.TrimSpace(input)
	num, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println(err)
	}

	return num, nil
}
*/

/* package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	fmt.Print("Lütfen Celcius derecesi giriniz: ")
	celcius, _ := getCelcius()

	kelvin := celcius + 273

	fmt.Println(kelvin, "Kelvin derecedir.")

}

func getCelcius() (int, error) {
	reader := bufio.NewReader(os.Stdin)

	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}

	input = strings.TrimSpace(input)
	num, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println(err)
	}

	return num, nil
}
*/

/* package main

import (
	"greeting"
)

func main() {
	greeting.Hello()
	greeting.Hi()
	greeting.Allgreet()

}

*/

/* package main

import (
	"fmt"
	"getcelcius"
)

func main() {

	fmt.Print("Lütfen Celcius derecesi giriniz: ")
	celcius, _ := getcelcius.GetCelcius()

	kelvin := celcius + 273

	fmt.Println(kelvin, "Kelvin derecedir.")

} */
