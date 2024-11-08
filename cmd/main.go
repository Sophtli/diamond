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

	d := diamond.New(iterations)

	fmt.Println(d.Build())
}
