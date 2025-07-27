package main

type triangle struct {
	base float64
	height float64
}

func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}

type square struct {
	sideLength float64
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

type shape interface {
	getArea() float64;
}

func main()  {
	t := triangle{base: 5, height: 10}
	s := square{sideLength: 4}

	printArea(t)
	printArea(s)
}

func printArea(s shape) {
	println("Area:", s.getArea())
}