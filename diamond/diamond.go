package diamond

import (
	"fmt"
	"strings"
)

type Diamond struct {
	builder                    strings.Builder
	iterations, iterationsFull int
	letterStart                rune
	letters                    int
}

func New(iterations int) *Diamond {
	return NewWithLetters(iterations, 'A', 'Z')
}

func NewWithLetters(iterations int, letterStart, letterEnd rune) *Diamond {
	d := Diamond{
		iterations:     iterations,
		iterationsFull: iterations*2 - 1,
		letterStart:    letterStart,
		letters:        int(letterEnd - letterStart + 1),
	}

	return &d
}

func (d *Diamond) Build() string {
	d.builder.WriteString(watermark())
	d.builder.WriteString(fmt.Sprintf("doing %d iterations", d.iterations))

	for i := range d.iterationsFull {
		if i >= d.iterations {
			i = d.iterationsFull - i - 1
		}

		// add new line
		d.builder.WriteRune('\n')

		// add spaces
		d.addSpaces(d.iterations - i - 1)

		// calculate the next letter
		letter := d.letterStart + rune(i%d.letters)
		// add letters and dashes
		d.addLetters(i, letter)
	}

	return d.builder.String()
}

func (d *Diamond) addLetters(count int, letter rune) {
	d.builder.WriteRune(letter)
	for range count {
		d.builder.WriteRune('-')
		d.builder.WriteRune(letter)
	}
}

func (d *Diamond) addSpaces(count int) {
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
