package main

import "fmt"

// Implement types Rectangle, Circle and Shape
type Rectangle struct {
	//two attributes
	Width  int32
	Length int32
}

type Circle struct {
	//one attribute
	Radius int32
}

//now we create a string to "see" properties of Rectangle and Circle
func (r Rectangle) String() string {
	return fmt.Sprint("Square of width ", r.Width, " and length ", r.Length)
}

func (c Circle) String() string {
	return fmt.Sprint("Circle of radius ", c.Radius)
}

type Shape interface { // we create an interface that could be used in Rectangle and Circle i.e.
	fmt.Stringer
}

func main() {
	r := &Rectangle{2, 3}
	c := &Circle{5}

	shapes := []Shape{r, c}

	for _, s := range shapes {
		fmt.Println(s)
		// Expected output:
		// Square of width 2 and length 3
		// Circle of radius 5
	}
}
