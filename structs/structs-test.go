package main

import (
	"fmt"
	"math"
)

type Circle struct {
	x float64
	y float64
	r float64
}

type Rectangle struct {
	w float64
	h float64
}

type Traingle struct {
	h float64
	b float64
}

// Multi Shape
type MultiShape struct {
	shapes []Shape
}

// Defining an interface
type Shape interface {
	area() float64
}

func main() {

	fmt.Println("hello")

	// Defining a struct

	// c2 is a *Circle
	c2 := new(Circle)

	c1 := Circle{12, 10, 7}
	fmt.Println("Circle 1", c1.x, c1.y, c1.r)
	fmt.Println("Circle 2", c2.x, c2.y, c2.r)
	fmt.Println("Circle 1 Area:", CalcCircleArea(c1))
	fmt.Println("Circle 1 Area with method ", c1.area())

	rectangle1 := Rectangle{4.1, 2.3}
	trngl := Traingle{4, 6}

	fmt.Println("Rectangle", rectangle1.w, rectangle1.h, rectangle1.area())
	fmt.Println("Total area : ", totalArea(&c1, &rectangle1))

	fmt.Println("Traingle", trngl.h, trngl.b, trngl.area())

	multiShape1 := MultiShape{
		shapes: []Shape{
			&c1,
			&rectangle1,
			&trngl,
		},
	}

	fmt.Println("Multi-Shape :", multiShape1.area())

}

// Generic function that calculates the area of Circle
func CalcCircleArea(c Circle) float64 {
	return math.Pi * c.r * c.r
}

// Method of Circle for calculate area
func (c *Circle) area() float64 {
	return math.Pi * c.r * c.r
}

// Method of Rectangle for calculate area
func (r *Rectangle) area() float64 {
	return r.w * r.h
}

// Method of Traingle
func (t *Traingle) area() float64 {
	return t.b * t.h / 2
}

// Method of MultiShape for calculating area
func (multiS *MultiShape) area() float64 {
	var area float64
	for _, s := range multiS.shapes {
		area += s.area()
	}
	return area
}

// Function that calculates the total area
func totalArea(shapes ...Shape) float64 {
	var area float64
	for _, s := range shapes {
		area += s.area()
	}
	return area
}
