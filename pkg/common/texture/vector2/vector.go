package vector2

import "math"

func New(x, y float64) Vector {
	return Vector{X: x, Y: y}
}

type Vector struct {
	X float64
	Y float64
}

func (v Vector) Dot(dv Vector) float64 {
	return (v.X * dv.X) + (v.Y * dv.Y)
}

func (v Vector) Norm() Vector {
	return v.DivF(v.Length())
}

func (v Vector) Length() float64 {
	return math.Sqrt(math.Pow(v.X, 2) + math.Pow(v.Y, 2))
}

func (v Vector) Div(dv Vector) Vector {
	return New(v.X/dv.X, v.Y/dv.Y)
}

func (v Vector) Sub(dv Vector) Vector {
	return New(v.X-dv.X, v.Y-dv.Y)
}

func (v Vector) Add(dv Vector) Vector {
	return New(v.X+dv.X, v.Y+dv.Y)
}

func (v Vector) DivF(factor float64) Vector {
	return New(v.X/factor, v.Y/factor)
}

func (v Vector) Round() (int, int) {
	return int(math.Floor(v.X)), int(math.Floor(v.Y))
}

func (v Vector) NextInt(dv Vector) Vector {
	xm := 1 / dv.X
	ym := 1 / dv.Y

	m := math.Min(xm, ym)

	return New(math.Floor(v.X+(dv.X*m)), math.Floor(v.Y+(dv.Y*m)))
}
