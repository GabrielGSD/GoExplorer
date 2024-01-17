package main

import (
	"fmt"

	"golang.org/x/tour/pic"
)

func Pic(dx, dy int) [][]uint8 {
	fmt.Println(dx, dy)
	pic := make([][]uint8, dy)

	for y := range pic {
		row := make([]uint8, dx)

		for x := range row {
			row[x] = uint8((x + y) / 2)
		}
		pic[y] = row
	}
	return pic
}

func main() {
	pic.Show(Pic)
}
