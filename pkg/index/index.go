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
	standardRamp = common.NewColorRamp(
		common.ColorPoint{Factor: 0, Color: color.Black, Blend: common.BlendLinear},
		common.ColorPoint{Factor: 1, Color: color.White, Blend: common.BlendConstant},
	)
	stripeRamp = common.NewColorRamp(
		common.ColorPoint{Factor: 0, Color: color.RGBA{0xff, 0x00, 0x00, 0xff}, Blend: common.BlendConstant},
		common.ColorPoint{Factor: 0.1, Color: color.White, Blend: common.BlendConstant},
		common.ColorPoint{Factor: 0.2, Color: color.Black, Blend: common.BlendConstant},
		common.ColorPoint{Factor: 0.3, Color: color.White, Blend: common.BlendConstant},
		common.ColorPoint{Factor: 0.4, Color: color.Black, Blend: common.BlendConstant},
		common.ColorPoint{Factor: 0.5, Color: color.White, Blend: common.BlendConstant},
		common.ColorPoint{Factor: 0.6, Color: color.Black, Blend: common.BlendConstant},
		common.ColorPoint{Factor: 0.7, Color: color.White, Blend: common.BlendConstant},
		common.ColorPoint{Factor: 0.8, Color: color.Black, Blend: common.BlendConstant},
		common.ColorPoint{Factor: 0.9, Color: color.White, Blend: common.BlendConstant},
		common.ColorPoint{Factor: 1, Color: color.White, Blend: common.BlendConstant},
	)
	lowCatch = common.NewColorRamp(
		common.ColorPoint{Factor: 0, Color: color.RGBA{0xff, 0x00, 0x00, 0xff}, Blend: common.BlendConstant},
		common.ColorPoint{Factor: 0.05, Color: color.Black, Blend: common.BlendLinear},
		common.ColorPoint{Factor: 1, Color: color.White, Blend: common.BlendConstant},
	)

	generators map[string]generate.Generator = map[string]generate.Generator{
		"blank":    &simple.BlankGenerator{},
		"gradient": &simple.GradientGenerator{},
		"ramp":     simple.RampGenerator{},
		"perlin": simple.TextureGenerator{
			Ramp: standardRamp,
			Factory: texture.PerlinFactory{
				Interpolate: interpolate.SmoothStep,
				Scale:       50,
			},
		},
		"ramp-tex": simple.TextureGenerator{
			Ramp: standardRamp,
			Factory: texture.GradientFactory{
				Interpolate: interpolate.Linear,
			},
		},
		"worley": simple.TextureGenerator{
			Ramp: stripeRamp,
			Factory: texture.WorleyFactory{
				CellsX: 5, CellsY: 5, Points: 1,
			},
		},
		"voronoi": simple.TextureGenerator{
			Ramp: stripeRamp,
			Factory: texture.VoronoiFactory{
				CellsX: 5, CellsY: 5, Points: 1,
			},
		},
		"stack": simple.TextureGenerator{
			Ramp: stripeRamp,
			Factory: texture.NewStack(
				texture.TextureStackFactoryInfo{
					Factor: 3, Factory: texture.VoronoiFactory{
						CellsX: 5, CellsY: 5, Points: 1,
					},
				},
				texture.TextureStackFactoryInfo{
					Factor: 1, Factory: texture.VoronoiFactory{
						CellsX: 20, CellsY: 20, Points: 1,
					},
				},
			),
		},
		"voronoi-lines": simple.TextureGenerator{
			Ramp: lowCatch,
			Factory: texture.VoronoiLinesFactory{
				CellsX: 5, CellsY: 5, Points: 1,
			},
		},
		"tiles": scene.TileScene{},
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
