package simple

import (
	"image"

	"github.com/TheWozard/GoGenerate/pkg/common"
	"github.com/TheWozard/GoGenerate/pkg/generate/params"
)

type RampGenerator struct {
	Ramp common.ColorRamp
	Func common.GradientFunction
}

func (rg RampGenerator) Gen(params *params.GenerationParams) (image.Image, error) {
	img := params.Image()
	f := rg.function(params)
	ramp := rg.ramp(params)

	width, height := img.Bounds().Dx(), img.Bounds().Dy()
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			img.Set(x, y, ramp.ColorAt(f(width, height, x, y)))
		}
	}

	return img, nil
}

func (rg RampGenerator) ramp(params *params.GenerationParams) common.ColorRamp {
	if rg.Ramp == nil {
		rand := params.Rand()
		colors := make([]common.ColorPoint, rand.Intn(5)+5)
		for i := range colors {
			colors[i] = common.ColorPoint{
				Factor: (1.0 / float64(len(colors)-1)) * float64(i),
				Color:  common.RandomColor(rand),
				Blend:  common.BlendType(rand.Int63n(2)),
			}
		}
		return common.NewColorRamp(colors...)
	}
	return rg.Ramp
}

func (rg RampGenerator) function(params *params.GenerationParams) common.GradientFunction {
	if rg.Func == nil {
		return common.LinearGradient
	}
	return rg.Func
}
