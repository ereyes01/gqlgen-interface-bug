package shapes

import "math"

type Shape interface {
	Area() float64
}

type Circle struct {
	Msg string
	Radius float64
}

func (c *Circle) Area() float64 {
	return c.Radius * math.Pi * math.Pi
}

func (c *Circle) GetMsg() string {
	return c.Msg
}

type Rectangle struct {
	Msg string
	Length, Width float64
}

func (r *Rectangle) Area() float64 {
	return r.Length * r.Width
}

func (r *Rectangle) GetMsg() string {
	return r.Msg
}

type Thing interface {
	GetMsg() string
}
