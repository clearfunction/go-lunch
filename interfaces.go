package main

import "fmt"

// START OMIT
type geometry interface {
	area() float64
}

type rect struct {
	width, height float64
}

func (r rect) area() float64 {
	return r.width * r.height
}

func announceArea(g geometry) { // HL
	fmt.Println("My area is: ", g.area())
}

func main() {
	aSquare := rect{width: 1, height: 1}
	announceArea(aSquare)
}

// END OMIT
