package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	s := make([][]uint8, dy)
	for y := range s{
		s[y] = make([]uint8, dx)
		for x := range s[y]{
			s[y][x] = uint8((x+1)^(y+3))	
		}
	}
		
	return s
}

func main() {
	pic.Show(Pic)
}
