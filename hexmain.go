package main

import (
	"fmt"
)

func main() {
	hexcodes := []string{"#FF0000", "#1133AA", "0xF8C22D", "0x332A8BBD1", "#3399CC"}
	for _, val := range hexcodes {
		// x := Hexconvert(val)
		fmt.Println(Hexconvert(val))
	}
}
