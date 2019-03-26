package colour

import (
	"image/color"
	"math"
)

// Returns a number (float) representing the luminance of a color.
//
// Ported from https://github.com/styled-components/polished/blob/master/src/color/getLuminance.js
func Luminance(colour color.RGBA) float64 {
	if colour.A == 0 {
		return 0
	}

	rChannel := mutateChannel(float64(colour.R) / 255)
	gChannel := mutateChannel(float64(colour.G) / 255)
	bChannel := mutateChannel(float64(colour.B) / 255)

	return 0.2126*rChannel + 0.7152*gChannel + 0.0722*bChannel
}

func mutateChannel(channel float64) float64 {
	if channel < 0.03928 {
		return channel / 12.92
	}

	return math.Pow(((channel + 0.055) / 1.055), 2.4)
}
