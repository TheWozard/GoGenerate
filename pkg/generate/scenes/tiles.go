package scene

import (
	"image"
	"image/color"
	"math"

	"github.com/TheWozard/GoGenerate/pkg/common"
	"github.com/TheWozard/GoGenerate/pkg/common/texture"
	"github.com/TheWozard/GoGenerate/pkg/common/texture/interpolate"
	"github.com/TheWozard/GoGenerate/pkg/common/texture/vector2"
	"github.com/TheWozard/GoGenerate/pkg/generate/params"
	"github.com/TheWozard/GoGenerate/pkg/generate/simple"
	"github.com/anthonynsimon/bild/blend"
)

type TileScene struct {
}

type tileSceneKey struct {
	x int
	y int
}

func (ts TileScene) Gen(param *params.GenerationParams) (*image.RGBA, error) {
	img := param.Image()
	cellSize := 200.0
	tileTexture := texture.VoronoiLinesFactory{
		CellsX: int(math.Floor(float64(param.Width) / cellSize)),
		CellsY: int(math.Floor(float64(param.Height) / cellSize)),
		Points: 1,
	}.VoronoiLinesTexture(param)
	tileColors := common.NewColorIndex(20, common.NewColorRamp(
		common.ColorPoint{Factor: 0, Color: color.RGBA{0x5f, 0x6f, 0x5f, 0xff}, Blend: common.BlendLinear},
		common.ColorPoint{Factor: 1, Color: color.RGBA{0x9f, 0x9f, 0x9f, 0xff}, Blend: common.BlendConstant},
	), *param)

	grassExtra := texture.NewStack(
		texture.TextureStackFactoryInfo{
			Factor: 1, Factory: texture.PerlinFactory{
				Interpolate: interpolate.SmoothStep,
				Scale:       int(cellSize) / 4,
			},
		},
		texture.TextureStackFactoryInfo{
			Factor: 1, Factory: texture.PerlinFactory{
				Interpolate: interpolate.SmoothStep,
				Scale:       int(cellSize) / 10,
			},
		},
	).Texture(param)
	grassColor := common.NewColorRamp(
		common.ColorPoint{Factor: 0, Color: color.RGBA{0x27, 0x4d, 0x32, 0xff}, Blend: common.BlendLinear},
		common.ColorPoint{Factor: 1, Color: color.RGBA{0x42, 0x8a, 0x58, 0xff}, Blend: common.BlendConstant},
	)

	shadow, _ := simple.ShadowGenerator{
		Direction: vector2.New(1, 1).Norm(),
		Height:    0.01,
		Texture: texture.ModifyTexture{
			F: func(f float64) float64 {
				return f
			},
			Wrapped: tileTexture,
		},
	}.Gen(&params.GenerationParams{
		Seed:   param.Seed,
		Height: param.Height,
		Width:  param.Width,
	})

	width, height := img.Bounds().Dx(), img.Bounds().Dy()
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			cx, cy, factor := tileTexture.NearestPoint(x, y)
			extraGrassFactor := grassExtra.FactorAt(x, y)
			factor = factor - interpolate.Exponential(0, 0.1, extraGrassFactor)
			tileColor := tileColors.GetColor(tileSceneKey{cx, cy})
			if factor > 0.02 {
				img.Set(x, y, tileColor)
			} else {
				img.Set(x, y, grassColor.ColorAt(1))
			}
		}
	}

	return blend.Darken(img, shadow), nil
}
