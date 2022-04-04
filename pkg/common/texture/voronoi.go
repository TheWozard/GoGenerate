package texture

import (
	"math"
	"math/rand"

	"github.com/TheWozard/GoGenerate/pkg/common/texture/vector2"
	"github.com/TheWozard/GoGenerate/pkg/generate/params"
)

type VoronoiFactory struct {
	CellsX int
	CellsY int
	Points int
}

func (wf VoronoiFactory) Texture(param *params.GenerationParams) FactorTexture {
	return wf.VoronoiTexture(param)
}

func (wf VoronoiFactory) VoronoiTexture(param *params.GenerationParams) *VoronoiTexture {
	return &VoronoiTexture{
		factory: wf,
		rand:    param.Rand(),
		xScale:  float64(param.Width) / float64(wf.CellsX),
		yScale:  float64(param.Height) / float64(wf.CellsY),
		cells:   make(map[voronoiStoreKey][]vector2.Vector, (wf.CellsX+2)*(wf.CellsY+2)),
		points:  make(map[voronoiStoreKey][]vector2.Vector, wf.CellsX*wf.CellsY),
	}
}

type VoronoiTexture struct {
	factory VoronoiFactory
	rand    *rand.Rand
	xScale  float64
	yScale  float64
	cells   map[voronoiStoreKey][]vector2.Vector
	points  map[voronoiStoreKey][]vector2.Vector
}

type voronoiStoreKey struct {
	x int
	y int
}

func (wt *VoronoiTexture) FactorAt(x, y int) float64 {
	distance, _ := wt.pointInformation(float64(x)/wt.xScale, float64(y)/wt.yScale)
	return math.Min(distance, 1)
}

func (wt *VoronoiTexture) NearestPoint(x, y int) (int, int, float64) {
	distance, point := wt.pointInformation(float64(x)/wt.xScale, float64(y)/wt.yScale)
	return int(point.X * wt.xScale), int(point.Y * wt.yScale), math.Min(distance, 1)
}

func (wt *VoronoiTexture) TotalPoints() int {
	return (wt.factory.CellsX + 2) * (wt.factory.CellsY + 2) * wt.factory.Points
}

func (wt *VoronoiTexture) pointInformation(x, y float64) (float64, vector2.Vector) {
	x0, y0 := int(math.Floor(x)), int(math.Floor(y))
	key := voronoiStoreKey{x: x0, y: y0}
	points := wt.getPoints(key)

	f1distance, f1vector := 2.0, vector2.New(0, 0)
	f2distance := 2.0
	offset := vector2.New(float64(x0), float64(y0))
	point := vector2.New(x, y)
	for _, target := range points {
		truePoint := target.Add(offset)
		temp := point.Sub(truePoint).Length()
		if temp < f1distance {
			f2distance = f1distance
			f1distance = temp
			f1vector = truePoint
		} else if temp < f2distance {
			f2distance = temp
		}
	}
	return f2distance - f1distance, f1vector
}

func (wt *VoronoiTexture) getPoints(key voronoiStoreKey) []vector2.Vector {
	points, ok := wt.points[key]
	if !ok {
		points = []vector2.Vector{}
		for x := -1; x <= 1; x += 1 {
			for y := -1; y <= 1; y += 1 {
				offset := vector2.New(float64(x), float64(y))
				cell := wt.getCell(voronoiStoreKey{x: key.x + x, y: key.y + y})
				temp := make([]vector2.Vector, len(cell))
				for i, point := range cell {
					temp[i] = point.Add(offset)
				}
				points = append(points, temp...)
			}
		}
		wt.points[key] = points
	}
	return points
}

func (wt *VoronoiTexture) getCell(key voronoiStoreKey) []vector2.Vector {
	cell, ok := wt.cells[key]
	if !ok {
		cell = make([]vector2.Vector, wt.factory.Points)
		for i := range cell {
			cell[i] = vector2.New(wt.rand.Float64(), wt.rand.Float64())
		}
		wt.cells[key] = cell
	}
	return cell
}
