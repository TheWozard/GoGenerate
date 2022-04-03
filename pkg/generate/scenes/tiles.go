package scene

import (
	"image"
	"image/color"
	"math"

	"github.com/TheWozard/GoGenerate/pkg/common"
	"github.com/TheWozard/GoGenerate/pkg/common/texture"
	"github.com/TheWozard/GoGenerate/pkg/generate/params"
)

type TileScene struct {
}

type tileSceneKey struct {
	x int
	y int
}

func (ts TileScene) Gen(params *params.GenerationParams) (image.Image, error) {
	img := params.Image()
	cellSize := 50.0
	primary := texture.WorleyFactory{
		CellsX: int(math.Floor(float64(params.Width) / cellSize)),
		CellsY: int(math.Floor(float64(params.Height) / cellSize)),
		Points: 1,
	}.WorleyTexture(params)
	overlay := texture.WorleyFactory{
		CellsX: int(math.Floor(float64(params.Width) / (cellSize / 2))),
		CellsY: int(math.Floor(float64(params.Height) / (cellSize / 2))),
		Points: 1,
	}.WorleyTexture(params)
	index := common.NewColorIndex(primary.TotalPoints(), common.NewColorRamp(
		common.ColorPoint{Factor: 0, Color: color.RGBA{0x1f, 0x1f, 0x1f, 0xff}, Blend: common.BlendLinear},
		common.ColorPoint{Factor: 1, Color: color.RGBA{0xcf, 0xcf, 0xcf, 0xff}, Blend: common.BlendConstant},
	))

	width, height := img.Bounds().Dx(), img.Bounds().Dy()
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			cx, cy, _ := primary.NearestPoint(x, y)
			factor := overlay.FactorAt(x, y)
			if factor < 0.7 {
				img.Set(x, y, index.GetColor(tileSceneKey{cx, cy}))
			} else {
				img.Set(x, y, color.RGBA{0x42, 0x8a, 0x58, 0xff})
			}
		}
	}

	return img, nil
}
