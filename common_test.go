package colour_test

import (
	"colour"
	"colour/test"
	"image/color"
	"testing"
)

func Test_NumberToHex_should_return_uint_as_string(t *testing.T) {
	helpers.AssertStringEquals("00", colour.NumberToHexString(0), t)
	helpers.AssertStringEquals("ff", colour.NumberToHexString(255), t)
}

func Test_ReduceHexStringValue_should_return_shortened_hex_string(t *testing.T) {
	helpers.AssertStringEquals("f86", colour.ReduceHexStringValue("ff8866"), t)
	helpers.AssertStringEquals("fff", colour.ReduceHexStringValue("fff"), t)
}

func Test_ConvertToHex_should_return_shortened_hex_string(t *testing.T) {
	helpers.AssertStringEquals("f00", colour.ConvertToHex(255, 0, 0), t)
	helpers.AssertStringEquals("0f0", colour.ConvertToHex(0, 255, 0), t)
	helpers.AssertStringEquals("00f", colour.ConvertToHex(0, 0, 255), t)
	helpers.AssertStringEquals("fff", colour.ConvertToHex(255, 255, 255), t)
	helpers.AssertStringEquals("000", colour.ConvertToHex(0, 0, 0), t)
}

func Test_ConvertRGBToHex_should_return_shortened_hex_string(t *testing.T) {
	helpers.AssertStringEquals("f00", colour.ConvertRGBToHex(color.RGBA{R: 255}), t)
	helpers.AssertStringEquals("0f0", colour.ConvertRGBToHex(color.RGBA{G: 255}), t)
	helpers.AssertStringEquals("00f", colour.ConvertRGBToHex(color.RGBA{B: 255}), t)
}

func Test_HSLToRGB_should_return_colour_type_from_hsl_input(t *testing.T) {
	helpers.AssertColourEquals(color.RGBA{R: 255, G: 255, B: 255, A: 255}, colour.HSLToRGB(colour.HslColour{Lightness: 1.0, Hue: 0.0, Saturation: 1.0}, func(R, G, B uint8) color.RGBA {
		return color.RGBA{R: R, G: G, B: B, A: 255}
	}), t)

	helpers.AssertColourEquals(color.RGBA{R: 78, G: 78, B: 78, A: 255}, colour.HSLToRGB(colour.HslColour{Lightness: 0.3, Hue: 0, Saturation: 1.0}, func(R, G, B uint8) color.RGBA {
		return color.RGBA{R: R, G: G, B: B, A: 255}
	}), t)
}
