package colour

import "image/color"

// Inverts the red, green and blue values of a color.
//
// Ported from https://github.com/styled-components/polished/blob/master/src/color/invert.js
func Invert(colour color.RGBA) color.RGBA {
	if colour.A == 0 {
		return color.RGBA{}
	}

	return color.RGBA{
		R: 255 - colour.R,
		G: 255 - colour.G,
		B: 255 - colour.B,
		A: colour.A,
	}
}
