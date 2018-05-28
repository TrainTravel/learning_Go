package main

import (
    "fmt"
    "math"
)

type Shape interface {
    // to implement a Shape interface, it should have an area method
    area() float64
}

type Square struct {
    side float64
}

type Circle struct {
    radius float64
}

func (s Square) area() float64{
    return s.side * s.side
}

func (c Circle) area() float64{
    return c.radius * c.radius * math.Pi
}

func info (s Shape) {
    fmt.Println(s)
    fmt.Println(s.area())
}

func main () {
    s := Square{2.5}
    c := Circle{2}
    info(s)
    info(c)

    shapes := []Shape{s, c}
    for key, s := range(shapes) {
        fmt.Println(key, "---", s.area())
    }
}
