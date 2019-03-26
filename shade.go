package colour

import "image/color"

// Shades a color by mixing it with black. `shade` can produce
// Hue shifts, where as `darken` manipulates the luminance channel and therefore
// doesn't produce Hue shifts.
//
// Ported from https://github.com/styled-components/polished/blob/master/src/color/shade.js
func Shade(weight float64, colour color.RGBA) color.RGBA {
	if colour.A == 0 {
		return color.RGBA{R: 0, G: 0, B: 0, A: 0}
	}
	return Mix(weight, color.RGBA{R: 0, G: 0, B: 0, A: 255}, colour)
}
