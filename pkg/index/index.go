package index

import (
	"fmt"
	"strings"

	"github.com/TheWozard/GoGenerate/pkg/generate"
	"github.com/TheWozard/GoGenerate/pkg/generate/simple"
)

var (
	generators map[string]generate.Generator = map[string]generate.Generator{
		"blank":    &simple.BlankGenerator{},
		"gradient": &simple.GradientGenerator{},
		"ramp":     simple.RampGenerator{},
	}
)

func GetGenerator(name string) (generate.Generator, error) {
	lower := strings.ToLower(name)
	gen, ok := generators[lower]
	if !ok {
		return nil, fmt.Errorf("unknown generator name '%s'", name)
	}

	return gen, nil
}
