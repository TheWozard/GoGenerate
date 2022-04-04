package generate

import (
	"image"

	"github.com/TheWozard/GoGenerate/pkg/generate/params"
)

// Generator interface for a standard generator
type Generator interface {
	Gen(params *params.GenerationParams) (*image.RGBA, error)
}
