package fractal

import (
	"math"
)

func Scale(factor float64) Transform {
	return func(p *Point) *Point {
		return p.Scale(factor)
	}
}

func Sinusoidal(p *Point) *Point {
	return &Point{
		math.Sin(p.X),
		math.Sin(p.Y),
	}
}

func Spherical(p *Point) *Point {
	r := p.R2()
	if r == 0 {
		return p
	}
	return p.Scale(1 / r)
}

func Swirl(p *Point) *Point {
	r := p.R2()
	return &Point{
		p.X*math.Sin(r) - p.Y*math.Cos(r),
		p.X*math.Cos(r) + p.Y*math.Sin(r),
	}
}

func Horseshoe(p *Point) *Point {
	r := p.R()
	if r != 0 {
		r = 1 / r
	}
	return &Point{
		r * (p.X - p.Y) * (p.X + p.Y),
		r * 2 * p.X * p.Y,
	}
}

func Bent(p *Point) *Point {
	if p.X >= 0 {
		p.X *= 2
	}
	if p.Y < 0 {
		p.Y /= -2
	}
	return p
}

func Fisheye(p *Point) *Point {
	return Eyefish(p)
}

func Exponential(p *Point) *Point {
	exp := math.Exp(p.X - 1)
	ypi := math.Pi * p.Y
	p.X = math.Cos(ypi)
	p.Y = math.Sin(ypi)
	return p.Scale(exp)
}

func Cosine(p *Point) *Point {
	return &Point{
		math.Cos(math.Pi*p.X) * math.Cosh(p.Y),
		-math.Sin(math.Pi*p.X) * math.Sinh(p.Y),
	}
}

func Eyefish(p *Point) *Point {
	return p.Scale(2 / (p.R() + 1))
}

func Bubble(p *Point) *Point {
	return p.Scale(4 / (p.R2() + 4))
}

func Cylinder(p *Point) *Point {
	return &Point{
		math.Sin(p.X),
		p.Y,
	}
}

func Tangent(p *Point) *Point {
	return &Point{
		math.Sin(p.X) / math.Cos(p.Y),
		math.Tan(p.Y),
	}
}

func Cross(p *Point) *Point {
	return p.Scale(1 / math.Abs(p.X*p.X-p.Y*p.Y))
}
