package texture

import (
	"math"
	"math/rand"

	"github.com/TheWozard/GoGenerate/pkg/common/texture/interpolate"
	"github.com/TheWozard/GoGenerate/pkg/common/texture/vector2"
	"github.com/TheWozard/GoGenerate/pkg/generate/params"
)

type PerlinFactory struct {
	Interpolate interpolate.Method
	Scale       int
}

func (pf PerlinFactory) Texture(param *params.GenerationParams) FactorTexture {
	return pf.PerlinTexture(param)
}

func (pf PerlinFactory) PerlinTexture(param *params.GenerationParams) *PerlinTexture {
	return &PerlinTexture{
		rand:        param.Rand(),
		interpolate: pf.Interpolate,
		scale:       pf.Scale,
		store:       map[storeKey]vector2.Vector{},
	}
}

type PerlinTexture struct {
	rand        *rand.Rand
	interpolate interpolate.Method
	scale       int
	store       map[storeKey]vector2.Vector
}

type storeKey struct {
	x int
	y int
}

func (pt *PerlinTexture) FactorAt(x, y int) float64 {
	return pt.perlin(float64(x)/float64(pt.scale), float64(y)/float64(pt.scale))
}

func (pt *PerlinTexture) perlin(x, y float64) float64 {
	// Grid points
	x0, y0 := int(math.Floor(x)), int(math.Floor(y))
	x1, y1 := x0+1, y0+1
	sx, sy := x-float64(x0), y-float64(y0)

	n0 := pt.dotGridGradient(x0, y0, x, y)
	n1 := pt.dotGridGradient(x1, y0, x, y)
	ix0 := pt.interpolate(n0, n1, sx)

	n0 = pt.dotGridGradient(x0, y1, x, y)
	n1 = pt.dotGridGradient(x1, y1, x, y)
	ix1 := pt.interpolate(n0, n1, sx)

	return (pt.interpolate(ix0, ix1, sy) + 1.0) / 2.0
	// return n0
}

func (pt *PerlinTexture) dotGridGradient(ix, iy int, x, y float64) float64 {
	gradient := pt.lookupGridPoint(ix, iy)
	return gradient.Dot(vector2.New(x-float64(ix), y-float64(iy)))
}

// lookupGridPoint returns the vector at a point on the grid
func (pt *PerlinTexture) lookupGridPoint(x, y int) vector2.Vector {
	key := storeKey{x: x, y: y}
	vec, ok := pt.store[key]
	if !ok {
		random := pt.rand.Float64() * (2 * math.Pi)
		vec := vector2.New(math.Cos(random), math.Sin(random))
		pt.store[key] = vec
	}
	return vec
}
