package common_test

import (
	"fmt"
	"image/color"
	"testing"

	"github.com/TheWozard/GoGenerate/pkg/common"
	"github.com/stretchr/testify/require"
)

func TestColorRamp(t *testing.T) {
	empty := common.NewColorRamp()
	basic := common.NewColorRamp(
		common.ColorPoint{Factor: 0.5, Color: color.Black, Blend: common.BlendLinear},
		common.ColorPoint{Factor: 0.75, Color: color.White, Blend: common.BlendConstant},
	)

	tests := []struct {
		name   string
		ramp   common.ColorRamp
		factor float64
		output color.Color
	}{
		{
			name: "empty start",
			ramp: empty, factor: 0,
			output: color.RGBA{R: 0x00, G: 0x00, B: 0x00, A: 0xff},
		},
		{
			name: "empty mid",
			ramp: empty, factor: 0.5,
			output: color.RGBA{R: 0x00, G: 0x00, B: 0x00, A: 0xff},
		},
		{
			name: "empty end",
			ramp: empty, factor: 1,
			output: color.RGBA{R: 0x00, G: 0x00, B: 0x00, A: 0xff},
		},

		{
			name: "basic start",
			ramp: basic, factor: 0,
			output: color.RGBA{R: 0x00, G: 0x00, B: 0x00, A: 0xff},
		},
		{
			name: "basic mid",
			ramp: basic, factor: 0.5,
			output: color.RGBA{R: 0x00, G: 0x00, B: 0x00, A: 0xff},
		},
		{
			name: "basic blend",
			ramp: basic, factor: 0.6,
			output: color.RGBA{R: 0x66, G: 0x66, B: 0x66, A: 0xff},
		},
		{
			name: "basic tail",
			ramp: basic, factor: 0.9,
			output: color.RGBA{R: 0xff, G: 0xff, B: 0xff, A: 0xff},
		},
		{
			name: "basic end",
			ramp: basic, factor: 1,
			output: color.RGBA{R: 0xff, G: 0xff, B: 0xff, A: 0xff},
		},
	}
	for i, test := range tests {
		t.Run(fmt.Sprintf("%d_%s", i, test.name), func(t *testing.T) {
			require.Equal(t, test.output, test.ramp.ColorAt(test.factor))
		})
	}
}
