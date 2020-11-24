// 1 -) 5 string elemandan oluşan bir array ve 5 int elemandan oluşan slice oluşturup index numaralarıyla birlikte gösterelim.

/* package main

import "fmt"

func main() {

	myArr := [5]string{"tahran", "belgrad", "roma", "tiflis", "moskova"}

	for index, value := range myArr {
		fmt.Println(index, value)
	}

	fmt.Println()

	mySlc := []int{1, 2, 3, 4, 5}

	for i, v := range mySlc {
		fmt.Println(i, v)
	}

} */

// 2 -) [0,1,2,3,4,5,6,7,8] slice dan [0,1,2,3], [4,5,6], [6,7,8], [2,3,6,7] slicelarını oluşturunuz.

/* package main

import "fmt"

func main() {

	mySlc := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}

	mySlc1 := mySlc[:4]
	fmt.Println(mySlc1)
	mySlc2 := mySlc[4:7]
	fmt.Println(mySlc2)
	mySlc3 := mySlc[6:]
	fmt.Println(mySlc3)

	mySlc4 := append(mySlc[2:4], mySlc[6:8]...)
	fmt.Println(mySlc4)

} */

// 3 -) slicelar için copy metodunu ve assign ( = ) ile farkını açıklayınız.
/*
package main

import "fmt"

func main() {

	 	mySlc := []int{1, 2, 3}
	   	mySlc2 := make([]int, 2)

	   	fmt.Println(mySlc)
	   	fmt.Println(mySlc2)

	   	copy(mySlc2, mySlc)
	   	fmt.Println(mySlc)
	   	fmt.Println(mySlc2)

	   	mySlc[0] = 100
	   	fmt.Println(mySlc)
	   	fmt.Println(mySlc2) */

/* 	mySlc := []int{1, 2, 3}
	mySlc2 := make([]int, 2)

	fmt.Println(mySlc)
	fmt.Println(mySlc2)

	mySlc2 = mySlc
	fmt.Println(mySlc)
	fmt.Println(mySlc2)

	mySlc[0] = 100
	fmt.Println(mySlc)
	fmt.Println(mySlc2)

}*/

// 4 -) map gösterimi ile yazar ve yazarlara ait kitapların isimlerini for range ile gösterin.

package main

import "fmt"

func main() {

	myMap := map[string][]string{
		"Yaşar Kemal":     []string{"İnce Memed", "Yer Demir Gök Bakır"},
		"Sabahattin Ali":  []string{"Kuyucaklı Yusuf", "Kürk Mantolu Madonna", "Değirmen"},
		"Haruki Murakami": []string{"1Q84", "Dans Dans Dans", "Kumandanı Öldürmek"},
	}

	/* 	fmt.Println(myMap)
	   	fmt.Println(myMap["Sabahattin Ali"])
	   	fmt.Println(myMap["Haruki Murakami"][0]) */

	for key, value := range myMap {
		fmt.Println("Yazarımız: ", key)
		fmt.Println("Bazı kitapları aşağıdadır:")
		for i, v := range value {
			fmt.Println("\t", i+1, v)
		}
	}

}
