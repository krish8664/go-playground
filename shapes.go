package main

import "math"

type Shape interface {
	area() float64
}

type rectangle struct {
	width, height float64
}

type circle struct {
	radius float64
}

//This can mutate the reciever struct as it be reference
func (r *rectangle) canMutate(w, h float64) {
	r.width = w
	r.height = h
}

//This can't mutate the reciever strct as it is by Value
func (r rectangle) cantMutate(w, h float64) {
	r.width = w
	r.height = h
}

func (r rectangle) area() float64 {
	return r.width * r.height
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
