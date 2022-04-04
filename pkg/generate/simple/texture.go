package simple

import (
	"image"

	"github.com/TheWozard/GoGenerate/pkg/common"
	"github.com/TheWozard/GoGenerate/pkg/common/texture"
	"github.com/TheWozard/GoGenerate/pkg/generate/params"
)

type TextureGenerator struct {
	Ramp    common.ColorRamp
	Factory texture.TextureFactory
}

func (tg TextureGenerator) Gen(params *params.GenerationParams) (*image.RGBA, error) {
	img := params.Image()
	texture := tg.Factory.Texture(params)

	width, height := img.Bounds().Dx(), img.Bounds().Dy()
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			img.Set(x, y, tg.Ramp.ColorAt(texture.FactorAt(x, y)))
		}
	}

	return img, nil
}
