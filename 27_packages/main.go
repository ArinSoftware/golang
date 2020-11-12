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

/* package main

import "fmt"

func main() {

	fmt.Println("Arin")
} */
/*
package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("seafood", "foo"))

	fmt.Println(strings.Count("animatrix", "a"))
}
*/

package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.ToUpper("Gopher"))
}
