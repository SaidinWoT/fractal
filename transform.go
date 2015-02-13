package fractal

type Coef struct {
	P, A, B, C, D, E, F float64
}

type Transform func(Point) Point

type Variation struct {
	W float64
	F Transform
}

func f(c Coef, vs []Variation) Transform {
	return func(p Point) Point {
		out := Point{}
		vp := Point{
			c.A*p.X + c.B*p.Y + c.C,
			c.D*p.X + c.E*p.Y + c.F,
		}

		for _, v := range vs {
			r := v.F(vp)
			r = r.Scale(v.W)
			out = out.Add(r)
		}
		return out
	}
}

func RandomFunc(r float64, cs []Coef, vs []Variation) Transform {
	var j int
	for r < 1 {
		r += cs[j].P
		j++
	}
	return f(cs[j-1], vs)
}
