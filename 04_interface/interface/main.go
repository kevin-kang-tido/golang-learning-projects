package main

import "fmt"

// Define an interface
type Shape interface {
    Area() float64   // method 
    Perimeter() float64
}

// Define a struct for Circle
type Circle struct {
    Radius float64
}

// Implement the Shape interface methods for Circle
func (c Circle) Area() float64 {
    return 3.14 * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
    return 2 * 3.14 * c.Radius
}

// Define a struct for Rectangle
type Rectangle struct {
    Width  float64
    Height float64
}

// Implement the Shape interface methods for Rectangle
func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
    return 2 * (r.Width + r.Height)
}

// A function that takes the Shape interface as a parameter
func PrintShapeDetails(s Shape) {
    fmt.Printf("Area: %f\n", s.Area())
    fmt.Printf("Perimeter: %f\n", s.Perimeter())
}

func main() {
    // Create instances of Circle and Rectangle
    circle := Circle{Radius: 5}
    rectangle := Rectangle{Width: 4, Height: 6}

    // Print details using the Shape interface
    fmt.Println("Circle Details:")
    PrintShapeDetails(circle)

    fmt.Println("\nRectangle Details:")
    PrintShapeDetails(rectangle)
}
