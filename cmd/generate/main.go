package main

import (
	"flag"
	"fmt"
	"image/png"
	"os"
	"time"

	"github.com/TheWozard/GoGenerate/pkg/generate/params"
	"github.com/TheWozard/GoGenerate/pkg/index"
)

func main() {
	seed := flag.String("seed", fmt.Sprintf("%d", time.Now().UTC().UnixNano()), "seed to be used for generation")
	height := flag.Int("height", 200, "height of output image file")
	width := flag.Int("width", 200, "width of the output image file")
	output := flag.String("output", "out.png", "output location of the output image file")

	flag.Parse()

	if len(flag.Args()) != 1 {
		fmt.Printf("Unexpected number of arguments '%d', expected usage 'cmd [generator]'\n", len(flag.Args()))
		os.Exit(1)
	}

	name := flag.Arg(0)
	gen, err := index.GetGenerator(name)
	if err != nil {
		fmt.Printf("Failed to get generator %v\n", err)
		os.Exit(1)
	}

	start := time.Now()
	// Generate image
	img, err := gen.Gen(&params.GenerationParams{
		Seed:   *seed,
		Height: *height,
		Width:  *width,
	})
	if err != nil {
		fmt.Printf("Failed to generate output: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Image generation completed in %dms", time.Since(start).Milliseconds())

	// Output to file
	file, err := os.Create(*output)
	if err != nil {
		fmt.Printf("Failed to create output file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()
	png.Encode(file, img)
}
