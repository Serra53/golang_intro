package main

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	perim() float64
}

type rect struct {
	title         string
	width, height float64
}

type circle struct {
	title  string
	radius float64
}

func main() {
	newCircle := circle{title: "Golden Circle", radius: 4.0}
	newRect := rect{title: "Backyard", width: 10.0, height: 5.0}
	printFeatures(newCircle.title, newCircle)
	printFeatures(newRect.title, newRect)

}

func (r rect) area() float64 {
	return r.width * r.height
}

func (r rect) perim() float64 {
	return 2*r.height + 2*r.width
}

func (c circle) area() float64 {
	return math.Pow(math.Pi*c.radius, 2)
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func printFeatures(title string, g geometry) {
	fmt.Printf("The %v has a perimeter of %v and a an area of %v\n",
		title, g.perim(), g.area())
}
