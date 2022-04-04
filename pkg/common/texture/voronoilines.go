package texture

import (
	"math"
	"math/rand"

	"github.com/TheWozard/GoGenerate/pkg/common/texture/vector2"
	"github.com/TheWozard/GoGenerate/pkg/generate/params"
)

type VoronoiLinesFactory struct {
	CellsX int
	CellsY int
	Points int
}

func (vlf VoronoiLinesFactory) Texture(param *params.GenerationParams) FactorTexture {
	return vlf.VoronoiLinesTexture(param)
}

func (vlf VoronoiLinesFactory) VoronoiLinesTexture(param *params.GenerationParams) *VoronoiLinesTexture {
	return &VoronoiLinesTexture{
		factory: vlf,
		rand:    param.Rand(),
		xScale:  float64(param.Width) / float64(vlf.CellsX),
		yScale:  float64(param.Height) / float64(vlf.CellsY),
		cells:   make(map[voronoiLinesStoreKey][]vector2.Vector, (vlf.CellsX+2)*(vlf.CellsY+2)),
		points:  make(map[voronoiLinesStoreKey][]vector2.Vector, vlf.CellsX*vlf.CellsY),
	}
}

type VoronoiLinesTexture struct {
	factory VoronoiLinesFactory
	rand    *rand.Rand
	xScale  float64
	yScale  float64
	cells   map[voronoiLinesStoreKey][]vector2.Vector
	points  map[voronoiLinesStoreKey][]vector2.Vector
}

type voronoiLinesStoreKey struct {
	x int
	y int
}

func (wt *VoronoiLinesTexture) FactorAt(x, y int) float64 {
	distance, _ := wt.pointInformation(float64(x)/wt.xScale, float64(y)/wt.yScale)
	return math.Min(distance, 1)
}

func (wt *VoronoiLinesTexture) NearestPoint(x, y int) (int, int, float64) {
	distance, point := wt.pointInformation(float64(x)/wt.xScale, float64(y)/wt.yScale)
	return int(point.X * wt.xScale), int(point.Y * wt.yScale), math.Min(distance, 1)
}

func (wt *VoronoiLinesTexture) TotalPoints() int {
	return (wt.factory.CellsX + 2) * (wt.factory.CellsY + 2) * wt.factory.Points
}

func (wt *VoronoiLinesTexture) pointInformation(x, y float64) (float64, vector2.Vector) {
	x0, y0 := int(math.Floor(x)), int(math.Floor(y))
	key := voronoiLinesStoreKey{x: x0, y: y0}
	points := wt.getPoints(key)

	var focal vector2.Vector
	f1distance, f1vector := 8.0, vector2.New(0, 0)
	f2distance := 8.0
	offset := vector2.New(float64(x0), float64(y0))
	point := vector2.New(x, y)
	for _, target := range points {
		truePoint := target.Add(offset)
		r := truePoint.Sub(point)
		temp := r.Dot(r)
		if temp < f1distance {
			f1distance = temp
			f1vector = r
			focal = truePoint
		}
	}

	for _, target := range points {
		r := target.Add(offset).Sub(point)
		temp := f1vector.Add(r).DivF(2).Dot(r.Sub(f1vector).Norm())
		if temp < f2distance {
			f2distance = temp
		}
	}

	return f2distance, focal
}

func (wt *VoronoiLinesTexture) getPoints(key voronoiLinesStoreKey) []vector2.Vector {
	points, ok := wt.points[key]
	if !ok {
		points = []vector2.Vector{}
		for x := -2; x <= 2; x += 1 {
			for y := -2; y <= 2; y += 1 {
				offset := vector2.New(float64(x), float64(y))
				cell := wt.getCell(voronoiLinesStoreKey{x: key.x + x, y: key.y + y})
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

func (wt *VoronoiLinesTexture) getCell(key voronoiLinesStoreKey) []vector2.Vector {
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
