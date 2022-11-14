package matrix

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix())
}

var maxColors uint = 2
var generatedColors []*Color = make([]*Color, 0)

func SetMaxColors(n uint) {
	maxColors = n
}

type Color [3]byte

func (c *Color) IsEqual(to *Color) bool {
	for idx, val := range c {
		if val != to[idx] {
			return false
		}
	}
	return true
}

func NewRandomColor() *Color {
	if len(generatedColors) == int(maxColors) {
		return generatedColors[rand.Intn(len(generatedColors))]
	}
	r, g, b := rand.Intn(255), rand.Intn(255), rand.Intn(255)
	color := Color{byte(r), byte(g), byte(b)}
	generatedColors = append(generatedColors, &color)
	return &color
}
