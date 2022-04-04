package common

import (
	"image/color"
	"math/rand"

	"github.com/TheWozard/GoGenerate/pkg/generate/params"
)

func NewColorIndex(items int, ramp ColorRamp, param params.GenerationParams) *ColorIndex {
	return &ColorIndex{
		count: items,
		rand:  param.Rand(),
		step:  1.0 / float64(items),
		ramp:  ramp,
		cache: map[interface{}]color.Color{},
	}
}

type ColorIndex struct {
	count int
	rand  *rand.Rand
	step  float64
	ramp  ColorRamp
	cache map[interface{}]color.Color
}

func (ci *ColorIndex) GetColor(index interface{}) color.Color {
	color, ok := ci.cache[index]
	if !ok {
		color = ci.ramp.ColorAt(float64(rand.Intn(ci.count)) * ci.step)
		ci.cache[index] = color
	}
	return color
}
