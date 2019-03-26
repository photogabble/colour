package colour

import (
	"image/color"
	"math"
)

// Mixes the two provided colors together by calculating the average of each of the RGB
// components weighted to the first color by the provided weight.
//
// Ported from https://github.com/styled-components/polished/blob/master/src/color/mix.js
func Mix(weight float64, colourA, colourB color.RGBA) color.RGBA {
	if colourA.A == 0 && colourB.A == 0 {
		return color.RGBA{R: 0, G: 0, B: 0, A: 0}
	} else if colourA.A == 0 {
		return colourB
	} else if colourB.A == 0 {
		return colourA
	}

	var (
		alphaDelta, x, y, z, weightA, weightB float64
	)

	// The formula is copied from the original Sass implementation:
	// http://sass-lang.com/documentation/Sass/Script/Functions.html#mix-instance_method
	alphaDelta = float64(colourA.A - colourB.A)
	x = weight*2 - 1
	if int8(x*alphaDelta) == -1 {
		y = x
	} else {
		y = x + alphaDelta
	}
	z = 1 + x*alphaDelta
	weightA = (y/z + 1) / 2.0
	weightB = 1 - weightA

	return color.RGBA{
		R: uint8(math.Floor(float64(colourA.R)*weightA + float64(colourB.R)*weightB)),
		G: uint8(math.Floor(float64(colourA.G)*weightA + float64(colourB.G)*weightB)),
		B: uint8(math.Floor(float64(colourA.B)*weightA + float64(colourB.B)*weightB)),
		A: uint8(float64(colourA.A) + (float64(colourB.A)-float64(colourA.A))*(weight/1.0)),
	}
}
