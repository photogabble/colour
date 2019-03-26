package colour

import (
	"image/color"
	"math"
)

// Increases the opacity of a color. Its range for the amount is between 0 to 1.
//
// Ported from https://github.com/styled-components/polished/blob/master/src/color/opacify.js
func Opacify(y float64, x color.RGBA) color.RGBA {
	if x.A == 0 {
		return color.RGBA{}
	}

	c := Guard(0, 255, float64(x.A)/255)
	n := Guard(0, 255, (c+y)*255)
	return color.RGBA{R: x.R, G: x.G, B: x.B, A: uint8(math.Ceil(n))}
}
