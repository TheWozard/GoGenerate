package index

import (
	"fmt"
	"image/color"
	"strings"

	"github.com/TheWozard/GoGenerate/pkg/common"
	"github.com/TheWozard/GoGenerate/pkg/common/texture"
	"github.com/TheWozard/GoGenerate/pkg/common/texture/interpolate"
	"github.com/TheWozard/GoGenerate/pkg/generate"
	scene "github.com/TheWozard/GoGenerate/pkg/generate/scenes"
	"github.com/TheWozard/GoGenerate/pkg/generate/simple"
)

var (
	generators map[string]generate.Generator = map[string]generate.Generator{
		"blank":    &simple.BlankGenerator{},
		"gradient": &simple.GradientGenerator{},
		"ramp":     simple.RampGenerator{},
		"perlin": simple.TextureGenerator{
			Ramp: common.NewColorRamp(
				common.ColorPoint{Factor: 0, Color: color.Black, Blend: common.BlendLinear},
				common.ColorPoint{Factor: 1, Color: color.White, Blend: common.BlendConstant},
			),
			Factory: texture.PerlinFactory{
				Interpolate: interpolate.SmootherStep,
				Scale:       50,
			},
		},
		"ramp-tex": simple.TextureGenerator{
			Ramp: common.NewColorRamp(
				common.ColorPoint{Factor: 0, Color: color.Black, Blend: common.BlendLinear},
				common.ColorPoint{Factor: 1, Color: color.White, Blend: common.BlendConstant},
			),
			Factory: texture.GradientFactory{
				Interpolate: interpolate.Linear,
			},
		},
		"worley": simple.TextureGenerator{
			Ramp: common.NewColorRamp(
				common.ColorPoint{Factor: 0, Color: color.Black, Blend: common.BlendLinear},
				common.ColorPoint{Factor: 1, Color: color.White, Blend: common.BlendConstant},
			),
			Factory: texture.WorleyFactory{
				CellsX: 5, CellsY: 5, Points: 1,
			},
		},
		"tiles": scene.TileScene{},
		"stack-worley": simple.TextureGenerator{
			Ramp: common.NewColorRamp(
				common.ColorPoint{Factor: 0, Color: color.Black, Blend: common.BlendLinear},
				common.ColorPoint{Factor: 1, Color: color.White, Blend: common.BlendConstant},
			),
			Factory: texture.NewStack(
				texture.TextureStackFactoryInfo{Factor: -1, Factory: texture.WorleyFactory{
					CellsX: 5, CellsY: 5, Points: 1,
				}},
				texture.TextureStackFactoryInfo{Factor: 2, Factory: texture.WorleyFactory{
					CellsX: 20, CellsY: 20, Points: 3,
				}},
				texture.TextureStackFactoryInfo{Factor: 1, Factory: texture.WorleyFactory{
					CellsX: 10, CellsY: 10, Points: 1,
				}},
			),
		},
	}
)

func GetGenerator(name string) (generate.Generator, error) {
	lower := strings.ToLower(name)
	gen, ok := generators[lower]
	if !ok {
		return nil, fmt.Errorf("unknown generator name '%s'", name)
	}

	return gen, nil
}
