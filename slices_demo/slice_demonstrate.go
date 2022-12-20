package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	out := make([][]uint8, dy)

	for i := 0; i < len(out); i++ {
		out[i] = make([]uint8, dx)
	}

	for y, _ := range out {
		for x, _ := range out[y] {
			out[y][x] = uint8(x ^ y)
		}
	}

	return out
}

func main() {
	pic.Show(Pic)
}
