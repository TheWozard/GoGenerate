package texture

import "github.com/TheWozard/GoGenerate/pkg/generate/params"

// FactorTexture returns a factor for the given location on an image
type FactorTexture interface {
	FactorAt(x, y int) float64
}

type TextureFactory interface {
	Texture(params *params.GenerationParams) FactorTexture
}

type ModifyTexture struct {
	F       func(float64) float64
	Wrapped FactorTexture
}

func (mt ModifyTexture) FactorAt(x, y int) float64 {
	return mt.F(mt.Wrapped.FactorAt(x, y))
}
