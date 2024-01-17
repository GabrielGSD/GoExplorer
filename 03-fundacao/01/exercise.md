# Slices Exercise: Pic

## Overview

This exercise involves implementing the `Pic` function, which should return a slice of length `dy`. Each element of this slice is itself a slice of `dx` 8-bit unsigned integers. When the program is executed, it will display the generated picture, interpreting the integers as grayscale (or bluescale) values.

The choice of image is up to you. Interesting functions include `(x+y)/2`, `x*y`, and `x^y`.

## Implementation Details

(You need to use a loop to allocate each `[]uint8` inside the `[][]uint8`.)

(Use `uint8(intValue)` to convert between types.)

## Base code

```go
package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
}

func main() {
	pic.Show(Pic)
}
```

## Reference

This exercise is part of the [Go Tour - Slices](https://go.dev/tour/moretypes/18).
