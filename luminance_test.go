package colour_test

import (
	"colour"
	"colour/test"
	"image/color"
	"math"
	"testing"
)

func Test_Luminance_will_return_the_luminance_of_a_rgba_colour(t *testing.T) {
	helpers.AssertFloat64Equals(0.16288822420427432, colour.Luminance(color.RGBA{R: 101, G: 100, B: 205, A: uint8(math.Floor(255 * 0.7))}), t)
}

func Test_Luminance_will_return_the_luminance_of_a_rgb_colour(t *testing.T) {
	helpers.AssertFloat64Equals(0.5742011250098834, colour.Luminance(color.RGBA{R: 204, G: 205, B: 100, A: 255}), t)
	helpers.AssertFloat64Equals(0.05780543019106723, colour.Luminance(color.RGBA{R: 68, G: 68, B: 68, A: 255}), t)
}

func Test_Luminance_will_return_zero_when_passed_transparent(t *testing.T) {
	helpers.AssertFloat64Equals(0, colour.Luminance(color.RGBA{R: 255}), t)
}
