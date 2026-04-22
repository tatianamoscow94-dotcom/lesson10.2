package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
}

type Circle struct{ Radius float64 }

func (c Circle) Area() float64 { return math.Pi * c.Radius * c.Radius }

type Rectangle struct{ Width, Height float64 }

func (r Rectangle) Area() float64 { return r.Width * r.Height }

func printAreas(shapes []Shape) {
	for _, s := range shapes {
		fmt.Printf("%.2f\n", s.Area())
	}
}
func main() {
	shapes := []Shape{Circle{7}, Rectangle{7, 9}}
	printAreas(shapes)
}
