/* package main

import "merhaba"

func main() {

	merhaba.Hola()
	merhaba.Hello()
	merhaba.Merhaba()

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

	fmt.Print("Lütfen Celcius dereceyi giriniz: ")
	celcius, err := getCelcius()
	if err != nil {
		fmt.Println(err)
	}

	fahrenheit := (celcius * 9 / 5) + 32

	fmt.Println(celcius, "Celcius derecesi ", fahrenheit, "Fahrenheit derecesine eşittir.")

}

func getCelcius() (float64, error) {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}

	input = strings.TrimSpace(input)
	num, err := strconv.ParseFloat(input, 64)
	if err != nil {
		fmt.Println(err)
	}

	return num, nil
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

	fmt.Print("Lütfen Celcius dereceyi giriniz: ")
	celcius, err := getCelcius()
	if err != nil {
		fmt.Println(err)
	}

	kelvin := celcius + 273

	fmt.Println(celcius, "Celcius derecesi ", kelvin, "Kelvin derecesine eşittir.")

}

func getCelcius() (float64, error) {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}

	input = strings.TrimSpace(input)
	num, err := strconv.ParseFloat(input, 64)
	if err != nil {
		fmt.Println(err)
	}

	return num, nil
} */

/* package main

import (
	"fmt"
	"getcelcius"
)

func main() {

	fmt.Print("Lütfen Celcius dereceyi giriniz: ")
	celcius, err := getcelcius.GetCelcius()
	if err != nil {
		fmt.Println(err)
	}

	kelvin := celcius + 273

	fmt.Println(celcius, "Celcius derecesi ", kelvin, "Kelvin derecesine eşittir.")

} */

package main

import (
	"fmt"
	"getcelcius"
)

func main() {

	fmt.Print("Lütfen Celcius dereceyi giriniz: ")
	celcius, err := getcelcius.GetCelcius()
	if err != nil {
		fmt.Println(err)
	}

	fahrenheit := (celcius * 9 / 5) + 32

	fmt.Println(celcius, "Celcius derecesi ", fahrenheit, "Fahrenheit derecesine eşittir.")

}
