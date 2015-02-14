package fractal

import "math"

type Point struct {
	X, Y float64
}

func (p Point) R2() float64 {
	return p.X*p.X + p.Y*p.Y
}

func (p Point) R() float64 {
	return math.Sqrt(p.R2())
}

func (p *Point) Add(q Point) *Point {
	p.X += q.X
	p.Y += q.Y
	return p
}

func (p *Point) Scale(factor float64) *Point {
	p.X *= factor
	p.Y *= factor
	return p
}

func (p *Point) Pixel() (int, int) {
	return 0, 0
}
