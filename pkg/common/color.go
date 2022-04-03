package common

import (
	"image/color"
	"math/rand"
)

// RandomColor creates a new random color.Color with the passed *rand.Rand
func RandomColor(r *rand.Rand) color.Color {
	return color.RGBA{
		R: uint8(r.Intn(256)),
		G: uint8(r.Intn(256)),
		B: uint8(r.Intn(256)),
		A: uint8(255),
	}
}

// BlendColor blends linearly between the passed from, to color.Colors.
// Amount is the percent of the final color to use with 0 returning from and 1 returning to.
// Amounts outside of [0, 1] are reduced
func BlendColor(from, to color.Color, amount float64) color.Color {
	// clean input
	if amount > 1 {
		amount = 1
	}
	if amount < 0 {
		amount = 0
	}

	fr, fg, fb, fa := from.RGBA()
	tr, tg, tb, ta := to.RGBA()
	inverse := 1 - amount
	return color.RGBA{
		R: uint8((float64(fr) * inverse / 256) + (float64(tr) * amount / 256)),
		G: uint8((float64(fg) * inverse / 256) + (float64(tg) * amount / 256)),
		B: uint8((float64(fb) * inverse / 256) + (float64(tb) * amount / 256)),
		A: uint8((float64(fa) * inverse / 256) + (float64(ta) * amount / 256)),
	}
}
