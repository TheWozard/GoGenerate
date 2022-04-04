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
	for x := -1; x < width+1; x++ {
		for y := -1; y < height+1; y++ {
			amount := sg.Texture.FactorAt(x, y)
			location := vector2.New(float64(x), float64(y))
			for amount > 0 {
				location = location.NextInt(sg.Direction)
				cx, cy := location.Round()
				if cx >= 0 && cx < width && cy >= 0 && cy < height {
					amount = amount - sg.Height
					local := sg.Texture.FactorAt(cx, cy)
					if amount > local {
						img.Set(cx, cy, color.RGBA{0x00, 0x00, 0x00, 0x33})
					} else {
						break
					}
				} else {
					break
				}
			}
		}
	}

	return img, nil
}
