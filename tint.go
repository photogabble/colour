package colour

import "image/color"

// Tints a color by mixing it with white. `tint` can produce
// Hue shifts, where as `lighten` manipulates the luminance channel and therefore
// doesn't produce Hue shifts.
//
// Ported from https://github.com/styled-components/polished/blob/master/src/color/tint.js
func Tint(weight float64, colour color.RGBA) color.RGBA {
	if colour.A == 0 {
		return color.RGBA{R: 0, G: 0, B: 0, A: 0}
	}

	return Mix(weight, color.RGBA{R: 255, G: 255, B: 255, A: 255}, colour)
}
