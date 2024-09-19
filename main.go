package main

import (
	"fmt"
)

func main() {
	var tingkat int

	fmt.Print("Input tingkatan (>=5): ")
	fmt.Scan(&tingkat)

	cetakSegitiga(tingkat)
}

func cetakSegitiga(tingkat int) {
	if tingkat >= 5 {
		for i := 1; i <= tingkat; i++ {
			for j := 0; j < tingkat-i; j++ {
				fmt.Print(" ")
			}
			for k := 0; k < (i*2 - 1); k++ {
				fmt.Print("1")
			}
			fmt.Println()
		}
	} else {
		fmt.Println("Minimal tingkatan adalah 5")
	}
}
