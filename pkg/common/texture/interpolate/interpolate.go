package interpolate

import "math"

// Method function for interpolating between points
type Method = func(a0, a1, w float64) float64

func Linear(a0, a1, w float64) float64 {
	if w < 0.0 {
		return a0
	}
	if w > 1.0 {
		return a1
	}
	return (a1-a0)*w + a0
}

func SmoothStep(a0, a1, w float64) float64 {
	if w < 0.0 {
		return a0
	}
	if w > 1.0 {
		return a1
	}
	return (a1-a0)*(3.0-w*2.0)*w*w + a0
}

func SmootherStep(a0, a1, w float64) float64 {
	if w < 0.0 {
		return a0
	}
	if w > 1.0 {
		return a1
	}
	return (a1-a0)*((w*(w*6.0-15.0)+10.0)*w*w*w) + a0
}

func Exponential(a0, a1, w float64) float64 {
	if w < 0.0 {
		return a0
	}
	if w > 1.0 {
		return a1
	}
	return math.Pow(w, math.E) + a0
}
