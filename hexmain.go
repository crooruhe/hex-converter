package main

import (
	// "fmt"
)

func main() {
	hexcodes := []string{"#FF0000", "#1133AA", "0xF8C22D", "0x332A8BBD1", "#3399CC"}
	//                    16711680,  1127338,    3381708   16302637
	for _, val := range hexcodes {
		Hexconvert(val)
		// fmt.Println(val)
	}
}
