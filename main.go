package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
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

	var builder strings.Builder

	builder.WriteString(watermark())
	builder.WriteString(fmt.Sprintf("doing %d iterations", iterations))

	letter_start := 'A'

	a := make([]int, iterations)
	for i := range iterations {
		a[i] = i
	}

	b := make([]int, iterations-1)
	copy(b, a)
	slices.Reverse(b)

	c := append(a, b...)

	for _, i := range c {
		// add new line
		builder.WriteRune('\n')

		// add spaces
		add_spaces(&builder, iterations-i-1)

		// calculate the next letter
		letter := letter_start + rune((i % 26))
		// add letters and dashes
		add_letters(&builder, i, letter)
	}

	fmt.Println(builder.String())
}

func add_letters(builder *strings.Builder, count int, letter rune) {
	builder.WriteRune(letter)
	for range count {
		builder.WriteRune('-')
		builder.WriteRune(letter)
	}
}

func add_spaces(builder *strings.Builder, count int) {
	for range count {
		builder.WriteRune(' ')
	}
}

func watermark() string {
	return `   ^
  / \
 / X \
DIAMOND
 \ X /
  \ /
   v
`
}
