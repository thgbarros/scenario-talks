package main

import (
	"math"
)

type Vertex2 struct {
	X, Y float64
}

func (v Vertex2) Abs() float64 {
	v.X += 5
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vertex2) AbsPointer() float64 {
	v.X += 5
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

type MyFloat float64

func (f MyFloat) ToFloat64() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}
