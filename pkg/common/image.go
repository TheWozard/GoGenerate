package common

import "image"

func BlendImage(a, b *image.RGBA, factor float64) {
	width, height := a.Bounds().Dx(), a.Bounds().Dy()
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			a.Set(x, y, BlendColor(a.At(x, y), b.At(x, y), factor))
		}
	}
}
