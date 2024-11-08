package diamond

import (
	"fmt"
	"strings"
)

type diamond struct {
	builder        strings.Builder
	iterations     int
	iterationsFull int
}

func New(iterations int) *diamond {
	d := diamond{
		iterations:     iterations,
		iterationsFull: iterations*2 - 1,
	}

	return &d
}

func (d *diamond) Build() string {
	d.builder.WriteString(watermark())
	d.builder.WriteString(fmt.Sprintf("doing %d iterations", d.iterations))

	letter_start := 'A'

	for i := range d.iterationsFull {
		if i >= d.iterations {
			i = d.iterationsFull - i - 1
		}

		// add new line
		d.builder.WriteRune('\n')

		// add spaces
		d.add_spaces(d.iterations - i - 1)

		// calculate the next letter
		letter := letter_start + rune(i)%26
		// add letters and dashes
		d.add_letters(i, letter)
	}

	return d.builder.String()
}

func (d *diamond) add_letters(count int, letter rune) {
	d.builder.WriteRune(letter)
	for range count {
		d.builder.WriteRune('-')
		d.builder.WriteRune(letter)
	}
}

func (d *diamond) add_spaces(count int) {
	for range count {
		d.builder.WriteRune(' ')
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
