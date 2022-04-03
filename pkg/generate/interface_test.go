package generate_test

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/TheWozard/GoGenerate/pkg/generate"
	"github.com/stretchr/testify/require"
)

func TestRandomColor(t *testing.T) {
	attempts := 20
	random := rand.New(rand.NewSource(0))
	numbers := make([]int, attempts)
	for i := range numbers {
		numbers[i] = (&generate.GenerationParams{
			Seed: fmt.Sprintf("%d", random.Int()),
		}).Rand().Int()
	}
	for i := range numbers {
		if i > 0 {
			require.NotEqual(t, numbers[i], numbers[i-1])
		}
		if i < len(numbers)-1 {
			require.NotEqual(t, numbers[i], numbers[i+1])
		}
	}
}
