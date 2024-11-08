package main

import (
	"fmt"
	"os"
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

	i := 0
	up := true
	for i >= 0 {
		// add new line
		builder.WriteRune('\n')

		// add spaces
		add_spaces(&builder, iterations-i-1)

		// calculate the next letter
		letter := letter_start + rune((i % 26))
		// add letters and dashes
		add_letters(&builder, i, letter)

		if i >= iterations-1 {
			up = false
		}
		if up {
			i++
		} else {
			i--
		}
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
