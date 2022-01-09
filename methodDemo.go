package main

import "fmt"

type Circle struct {
	radius float64
}

func main() {
	var c1 Circle
	fmt.Println(c1.radius)

	c1.radius = 10.00
	fmt.Println("圆的面积 = ", c1.getArea)

	c1.changeRadius(20)
	fmt.Println(c1.radius)

	change(&c1, 30)
	fmt.Println(c1.radius)
}

func (c Circle) getArea() float64 {
	return 3.14 * c.radius * c.radius
}

func (c *Circle) changeRadius(radius float64) {
	c.radius = radius
}

func change(c *Circle, radius float64) {
	c.radius = radius
}
