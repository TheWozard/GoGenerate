package common

import (
	"image/color"
	"sort"
)

// BlendType enum for the possible blend types used in a Color Point
type BlendType int64

const (
	BlendLinear BlendType = iota
	BlendConstant
)

var (
	// When a color ramp does not contain a color this will be the default color returned
	defaultColorRampColor color.Color = color.RGBA{R: 0x00, G: 0x00, B: 0x00, A: 0xff}
)

// ColorPoint is a location on a color ramp
type ColorPoint struct {
	Factor float64
	Color  color.Color
	Blend  BlendType
}

// NewColorRamp creates a new color ramp based on the passed points
func NewColorRamp(points ...ColorPoint) ColorRamp {
	sort.Slice(points, func(i, j int) bool {
		return points[i].Factor < points[j].Factor
	})
	return ColorRamp(points)
}

// ColorRamp converts a passed factor [0,1] and converts it to a color based on pre defined rules
type ColorRamp []ColorPoint

// ColorAt get teh color on a ColorRamp at the passed factor
func (ramp ColorRamp) ColorAt(factor float64) color.Color {
	srcColor, destinationColor, blendMode := defaultColorRampColor, defaultColorRampColor, BlendConstant
	srcFactor, destinationFactor := 0.0, 1.0

	// Update the colors and blend method based on the ramp
	for i, point := range ramp {
		// The ramp is guaranteed to be in order
		if point.Factor <= factor && (i == len(ramp)-1 || ramp[i+1].Factor > factor) {
			srcColor = point.Color
			blendMode = point.Blend
			srcFactor = point.Factor
			if i < len(ramp)-1 {
				destination := ramp[i+1]
				destinationColor = destination.Color
				destinationFactor = destination.Factor
			}
			break
		}
	}

	// Calculate the output color
	switch blendMode {
	case BlendLinear:
		level := (factor - srcFactor) / (destinationFactor - srcFactor)
		return BlendColor(srcColor, destinationColor, level)
	case BlendConstant:
		return srcColor
	default:
		return defaultColorRampColor
	}
}
