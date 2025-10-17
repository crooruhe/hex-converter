package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	const (
		Reset  = "\033[0m"
		Red    = "\033[31m"
		Green  = "\033[32m"
		Yellow = "\033[33m"
		// Blue   = "\033[34m"
		// Purple = "\033[35m"
		// Cyan = "\033[36m"
		// White  = "\033[37m"

		// // Bold
		// BoldRed   = "\033[1;31m"
		// BoldGreen = "\033[1;32m"

		// // Background colors
		// BgRed   = "\033[41m"
		// BgGreen = "\033[42m"
	)
	// fmt.Println("Number of arguments:", len(os.Args)-1)
	hexcodes := os.Args[1:]
	// fmt.Println("# of args: ", len(hexcodes))
	for _, val := range hexcodes {
		convHex := fmt.Sprint(Hexconvert(val))
		if convHex == "<nil>" || convHex == "" {
			continue
		}
		hlen := len(convHex)
		if strings.HasPrefix(convHex, "[") {
			hlen = hlen - 2
		}
		fmt.Print(Red + val + Reset + Yellow)
		time.Sleep(80 * time.Millisecond)
		fmt.Print(" =")
		time.Sleep(80 * time.Millisecond)
		for i := 1; i < hlen; i++ {
			fmt.Print("=")
			time.Sleep(80 * time.Millisecond)
		}
		fmt.Print("> ")
		fmt.Print(Reset + Green + convHex + Reset)
		fmt.Println()
	}
}
