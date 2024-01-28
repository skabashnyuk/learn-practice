package main

import (
	"fmt"
	"log"
)

type Square struct {
	X      int
	Y      int
	Length int
}

func NewSquare(x int, y int, length int) (*Square, error) {
	if length <= 0 {
		return nil, fmt.Errorf("length must be > 0 (was %d)", length)
	}
	return &Square{X: x, Y: y, Length: length}, nil
}
func (s *Square) Move(dx int, dy int) {
	s.X += dx
	s.Y += dy
}
func (s *Square) Area() int {
	return s.Length * s.Length
}

func main() {
	s, err := NewSquare(1, 1, 10)
	if err != nil {
		log.Fatalf("Error: %s\n", err)
	}
	s.Move(2, 3)
	fmt.Println("%+v\n", s)
	fmt.Println(s.Area())
}
