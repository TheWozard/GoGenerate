package common

type GradientFunction func(height, width, x, y int) float64

// LinearGradient linear width wise gradient
func LinearGradient(width, height, x, y int) float64 {
	return float64(x) / float64(width)
}
