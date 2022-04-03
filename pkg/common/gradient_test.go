package common_test

import (
	"fmt"
	"testing"

	"github.com/TheWozard/GoGenerate/pkg/common"
	"github.com/stretchr/testify/require"
)

func TestLinearGradient(t *testing.T) {
	tests := []struct {
		name   string
		x      int
		y      int
		width  int
		height int
		output float64
	}{
		{
			name: "left side",
			x:    0, y: 0,
			height: 100, width: 100,
			output: 0,
		},
		{
			name: "right side",
			x:    100, y: 0,
			height: 100, width: 100,
			output: 1,
		},
		{
			name: "middle",
			x:    50, y: 0,
			height: 100, width: 100,
			output: 0.5,
		},
	}
	for i, test := range tests {
		t.Run(fmt.Sprintf("%d_%s", i, test.name), func(t *testing.T) {
			require.Equal(t, test.output, common.LinearGradient(test.width, test.height, test.x, test.y))
		})
	}
}
