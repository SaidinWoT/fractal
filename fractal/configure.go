package main

import . "github.com/SaidinWoT/fractal"

func configure(p float64) ([]Coef, []Variation) {
	var cs []Coef
	var vs []Variation

	cs = append(cs, Coef{.5, .5, 0, 0, 0, .5, 0})
	cs = append(cs, Coef{.25, .5, 0, 1, 0, .5, 0})
	cs = append(cs, Coef{.25, .5, 0, 0, 0, .5, 1})

	vs = append(vs, Variation{1, Scale(p * 2)})

	return cs, vs
}
