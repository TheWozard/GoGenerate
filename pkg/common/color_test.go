package common_test

import (
	"fmt"
	"image/color"
	"math/rand"
	"testing"

	"github.com/TheWozard/GoGenerate/pkg/common"
	"github.com/stretchr/testify/require"
)

func TestRandomColor(t *testing.T) {
	attempts := 20
	random := rand.New(rand.NewSource(0))
	colors := make([]color.Color, attempts)
	for i := range colors {
		colors[i] = common.RandomColor(random)
	}
	for i := range colors {
		if i > 0 {
			require.NotEqual(t, colors[i], colors[i-1])
		}
		if i < len(colors)-1 {
			require.NotEqual(t, colors[i], colors[i+1])
		}
	}
}

func TestBlendColor(t *testing.T) {
	tests := []struct {
		name    string
		from    color.Color
		to      color.Color
		percent float64
		output  color.Color
	}{
		{
			name:    "none",
			from:    color.White,
			to:      color.Black,
			percent: 0,
			output:  color.RGBA{R: 0xff, G: 0xff, B: 0xff, A: 0xff}, // The blended output is always RGBA
		},
		{
			name:    "full",
			from:    color.White,
			to:      color.Black,
			percent: 1,
			output:  color.RGBA{R: 0x00, G: 0x00, B: 0x00, A: 0xff},
		},
		{
			name:    "half",
			from:    color.White,
			to:      color.Black,
			percent: 0.5,
			output:  color.RGBA{R: 0x7f, G: 0x7f, B: 0x7f, A: 0xff},
		},
		{
			name:    "mix",
			from:    color.RGBA{R: 0xff, G: 0x00, B: 0x00, A: 0x00},
			to:      color.RGBA{R: 0xff, G: 0x7f, B: 0x00, A: 0xff},
			percent: 0.10,
			output:  color.RGBA{R: 0xff, G: 0x0c, B: 0x00, A: 0x19},
		},
	}
	for i, test := range tests {
		t.Run(fmt.Sprintf("%d_%s", i, test.name), func(t *testing.T) {
			require.Equal(t, test.output, common.BlendColor(test.from, test.to, test.percent))
		})
	}
}
