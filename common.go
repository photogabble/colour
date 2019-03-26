package colour

import (
	"encoding/hex"
	"image/color"
	"math"
	"strings"
)

type HslColour struct {
	Hue, Saturation, Lightness float64
}

type HslaColour struct {
	*HslColour
	alpha float64
}

type convert func(R, G, B uint8) color.RGBA

func NumberToHexString(colour uint8) string {
	var input = []byte{colour}
	return hex.EncodeToString(input)
}

// Reduces hex values if possible e.g. ff8866 to f86
func ReduceHexStringValue(h string) string {
	a := strings.Split(h, "")
	if len(a) == 6 && a[0] == a[1] && a[2] == a[3] && a[4] == a[5] {
		return a[0] + a[2] + a[4]
	}
	return h
}

func ConvertRGBToHex(colour color.RGBA) string {
	return ConvertToHex(colour.R, colour.G, colour.B)
}

func ConvertToHex(R, G, B uint8) string {
	return ReduceHexStringValue(NumberToHexString(R) + NumberToHexString(G) + NumberToHexString(B))
}

func HSLToRGB(hsl HslColour, fn convert) color.RGBA {
	if hsl.Saturation == 0 {
		// achromatic
		return fn(uint8(math.Floor(hsl.Lightness)), uint8(math.Floor(hsl.Lightness)), uint8(math.Floor(hsl.Lightness)))
	}

	var (
		huePrime, chroma, secondComponent, red, green, blue, lightnessModification float64
	)

	// formula from https://en.wikipedia.org/wiki/HSL_and_HSV
	huePrime = math.Mod(hsl.Hue, 360) / 60
	chroma = (1 - math.Abs(2*hsl.Lightness-1)) * hsl.Saturation
	secondComponent = chroma * (1 - math.Abs(math.Mod(huePrime, 2)-1))
	lightnessModification = hsl.Lightness - chroma/2

	red = 0
	green = 0
	blue = 0

	if huePrime >= 0 && huePrime < 1 {
		red = chroma
		green = secondComponent
	} else if huePrime >= 1 && huePrime < 2 {
		red = secondComponent
		green = chroma
	} else if huePrime >= 2 && huePrime < 3 {
		green = chroma
		blue = secondComponent
	} else if huePrime >= 3 && huePrime < 4 {
		green = secondComponent
		blue = chroma
	} else if huePrime >= 4 && huePrime < 5 {
		red = secondComponent
		blue = chroma
	} else if huePrime >= 5 && huePrime < 6 {
		red = chroma
		blue = secondComponent
	}

	return fn(uint8(math.Floor((red+lightnessModification)*255)), uint8(math.Floor((green+lightnessModification))*255), uint8(math.Floor((blue+lightnessModification)*255)))
}

//func HSLToHex(Hue, Saturation, Lightness float64) string {
//
//}
