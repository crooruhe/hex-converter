// crooruhe
package main

import (
	"fmt"
	"strings"
	"unicode"
	"math/big"
)

func Hexconvert(args string) {
	// this func is same as -> .SetString(hexNumber, 16)
	fmt.Println()
	fmt.Println(args)
	var result []int
	sum := new(big.Int)

	if strings.HasPrefix(args, "#") {
		args = args[1:]

		length := len(args)
		if length % 2 != 0 {
			fmt.Println("Not a hex number")
			return
		}
		for i := 0; i < length / 2; i++ {
			idx := i * 2
			val := 0

			if unicode.IsLetter(rune(args[idx])) {
				if args[idx] - 55 > 15{
					fmt.Println("Not a valid hexidecimal value")
					return
				}
				val += (int(byte(args[idx])) - 55) * 16
			} else {
				val += (int(args[idx]) - 48) * 16
			}

			if unicode.IsLetter(rune(args[idx + 1])) {
				val += int(byte(args[idx + 1])) - 55
			} else {
				val += int(args[idx + 1]) - 48
			}

			result = append(result, val)
		}
		fmt.Println(result)
		return
	} else if strings.HasPrefix(args, "0x") {
		args = args[2:]
		length := len(args)
		counter := length - 1
		reversecounter := 0
		tempsum := new(big.Int)
		for counter > -1 {
			tempsum.SetInt64(0)
			if unicode.IsLetter(rune(args[counter])) {
				if args[counter] - 55 > 15{
					fmt.Println("Not a valid hexidecimal value")
					return
				}
				tempsum.Add(tempsum, new(big.Int).Mul(new(big.Int).SetInt64(int64(args[counter]-'0'-7)), new(big.Int).Exp(big.NewInt(16), big.NewInt(int64(reversecounter)), nil)))
			} else {
				tempsum.Add(tempsum, new(big.Int).Exp(big.NewInt(16), big.NewInt(int64(reversecounter)), nil))
				tempsum.Mul(tempsum, new(big.Int).SetInt64(int64(args[counter]) - 48))
			}
			sum.Add(sum, tempsum)
			counter--
			reversecounter++
		}
		fmt.Println(sum)
	}
}
