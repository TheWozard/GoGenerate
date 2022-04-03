package simple

import (
	"image"
	"image/color"

	"github.com/TheWozard/GoGenerate/pkg/common"
	"github.com/TheWozard/GoGenerate/pkg/generate"
)

type GradientGenerator struct {
	From color.Color
	To   color.Color

	Func common.GradientFunction
}

func (gg *GradientGenerator) Gen(params *generate.GenerationParams) (image.Image, error) {
	img := params.Image()
	from := gg.from(params)
	to := gg.to(params)
	f := gg.function(params)

	width, height := img.Bounds().Dx(), img.Bounds().Dy()
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			img.Set(x, y, common.BlendColor(from, to, f(width, height, x, y)))
		}
	}

	return img, nil
}

func (gg *GradientGenerator) from(params *generate.GenerationParams) color.Color {
	if gg.From == nil {
		gg.From = common.RandomColor(params.Rand())
	}
	return gg.From
}

func (gg *GradientGenerator) to(params *generate.GenerationParams) color.Color {
	if gg.To == nil {
		gg.To = common.RandomColor(params.Rand())
	}
	return gg.To
}

func (gg *GradientGenerator) function(params *generate.GenerationParams) common.GradientFunction {
	if gg.Func == nil {
		gg.Func = common.LinearGradient
	}
	return gg.Func
}
