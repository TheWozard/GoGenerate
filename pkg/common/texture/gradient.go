package texture

import (
	"fmt"

	"github.com/TheWozard/GoGenerate/pkg/common/texture/interpolate"
	"github.com/TheWozard/GoGenerate/pkg/common/texture/vector2"
	"github.com/TheWozard/GoGenerate/pkg/generate/params"
)

type GradientFactory struct {
	Interpolate interpolate.Method
}

func (gf GradientFactory) Texture(param *params.GenerationParams) FactorTexture {
	gradient := vector2.New(float64(param.Width), float64(param.Height))
	fmt.Println(gradient, gradient.Norm())
	return gradientTexture{
		interpolate: gf.Interpolate,
		length:      gradient.Length(),
		norm:        gradient.Norm(),
	}
}

type gradientTexture struct {
	interpolate interpolate.Method
	length      float64
	norm        vector2.Vector
}

func (gt gradientTexture) FactorAt(x, y int) float64 {
	return gt.interpolate(0, 1, gt.norm.Dot(vector2.New(float64(x), float64(y)).DivF(gt.length)))
}
