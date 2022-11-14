package main

import (
	"flag"
	"homework/internal/matrix"
)

var (
	matM    uint
	matN    uint
	colorsN uint
)

func main() {
	flag.UintVar(&matM, "m", 0, "colls number")
	flag.UintVar(&matN, "n", 0, "rows number")
	flag.UintVar(&colorsN, "c", 0, "colors number")
	flag.Parse()
	matrix.SetMaxColors(colorsN)
	mat := matrix.NewMatrix()
	mat.Fill(matM, matN)
	mat.PrintMaxEqualsInRow()
}
