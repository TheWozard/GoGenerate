package simple

import (
	"image"
	"image/color"

	"github.com/TheWozard/GoGenerate/pkg/common"
	"github.com/TheWozard/GoGenerate/pkg/generate"
)

type BlankGenerator struct {
	Color color.Color
}

func (bg *BlankGenerator) Gen(params *generate.GenerationParams) (image.Image, error) {
	img := params.Image()
	background := bg.color(params)

	for x := 0; x < img.Bounds().Dx(); x++ {
		for y := 0; y < img.Bounds().Dy(); y++ {
			img.Set(x, y, background)
		}
	}

	return img, nil
}

func (bg *BlankGenerator) color(params *generate.GenerationParams) color.Color {
	if bg.Color == nil {
		bg.Color = common.RandomColor(params.Rand())
	}
	return bg.Color
}
