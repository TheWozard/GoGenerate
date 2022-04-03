package common

import "image/color"

func NewColorIndex(items int, ramp ColorRamp) *ColorIndex {
	return &ColorIndex{
		factor: 0,
		step:   1.0 / float64(items),
		ramp:   ramp,
		cache:  map[interface{}]color.Color{},
	}
}

type ColorIndex struct {
	factor float64
	step   float64
	ramp   ColorRamp
	cache  map[interface{}]color.Color
}

func (ci *ColorIndex) GetColor(index interface{}) color.Color {
	color, ok := ci.cache[index]
	if !ok {
		color = ci.ramp.ColorAt(ci.factor)
		ci.cache[index] = color
		ci.factor = (ci.factor + ci.step)
	}
	return color
}
