package simple

import (
	"image"
	"image/color"

	"github.com/TheWozard/GoGenerate/pkg/common/texture"
	"github.com/TheWozard/GoGenerate/pkg/common/texture/vector2"
	"github.com/TheWozard/GoGenerate/pkg/generate/params"
)

type ShadowGenerator struct {
	Direction vector2.Vector
	Height    float64
	Texture   texture.FactorTexture
}

func (sg ShadowGenerator) Gen(params *params.GenerationParams) (*image.RGBA, error) {
	img := params.Image()

	width, height := img.Bounds().Dx(), img.Bounds().Dy()
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			amount := sg.Texture.FactorAt(x, y)
			location := vector2.New(float64(x), float64(y))
			for amount > 0 {
				location = location.Add(sg.Direction)
				cx, cy := location.Round()
				if cx >= 0 && cx < width && cy >= 0 && cy < height {
					amount = amount - sg.Height - sg.Texture.FactorAt(cx, cy)
					if amount > 0 {
						img.Set(cx, cy, color.Black)
					}
				} else {
					break
				}
			}
		}
	}

	return img, nil
}
