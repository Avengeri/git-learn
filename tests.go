package main

import "fmt"

type Square struct {
	Side int
}

func (s Square) Perimeter() {
	//fmt.Printf("%T,%#v \n", s, s)
	fmt.Printf("Периметр фигуры: %d сантиметров\n", s.Side*4)
}

func (s Square) Volume() {
	fmt.Printf("Объем фигуры: %d кубических сантиметров\n", s.Side*s.Side*s.Side)
}
func (s *Square) Scale(multiplier int) {
	s.Side *= multiplier
}
func definition() {
	s := Square{5}
	ps := &Square{5}
	s.Perimeter()
	s.Volume()
	ps.Perimeter()
	ps.Volume()
	ps.Scale(2)
	fmt.Printf("что это %d", ps.Scale)
}
func main() {
	definition()

}
