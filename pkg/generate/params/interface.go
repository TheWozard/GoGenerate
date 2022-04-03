package params

import (
	"image"
	"math/rand"
)

// GenerationParams common params between all generators
type GenerationParams struct {
	Seed   string
	Height int
	Width  int

	generator *rand.Rand
	image     *image.RGBA
}

func (gp *GenerationParams) Rand() *rand.Rand {
	if gp.generator == nil {
		seed := int64(0)
		for _, b := range []byte(gp.Seed) {
			seed = seed + int64(b)
		}
		gp.generator = rand.New(rand.NewSource(seed))
	}
	return gp.generator
}

func (gp *GenerationParams) Image() *image.RGBA {
	if gp.image == nil {
		gp.image = image.NewRGBA(image.Rectangle{image.Point{0, 0}, image.Point{gp.Width, gp.Height}})
	}
	return gp.image
}
