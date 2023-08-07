package main

import "fmt"

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func ScaleFunc(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	v.Scale(2)
	ScaleFunc(&v, 10)
	p := &Vertex{4, 3}

	// (&p).Scaleとして解釈する (利便性のため)
	p.Scale(3)
	ScaleFunc(p, 8)
	fmt.Println(v, p)
}
