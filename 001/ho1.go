package main

import (
	"fmt"
	"math"
)

type square struct {
	side float64
	}
type circle struct {
	radius float64
	}

func (s square) area() float64 {
	return s.side * s.side
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

type shape interface {
	area() float64
}

func info (s shape) {
	fmt.Println(s.area())
}

func main() {
	sq := square{20}
	crcl := circle{20}
	info(sq)
	info(crcl)
}