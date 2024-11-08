package main

import (
	"diamond/diamond"
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Argument is required")
		return
	}
	iterations, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Argument must be an int")
		return
	}

	var d *diamond.Diamond

	if len(os.Args) >= 4 {
		letterStart := rune(os.Args[2][0])
		if err != nil {
			fmt.Println("Argument must be an char")
			return
		}
		letterEnd := rune(os.Args[3][0])
		if err != nil {
			fmt.Println("Argument must be an char")
			return
		}
		d = diamond.NewWithLetters(iterations, rune(letterStart), rune(letterEnd))
	} else {
		d = diamond.New(iterations)
	}

	fmt.Println(d.Build())
}
