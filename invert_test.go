package colour_test

import (
	"colour"
	"colour/test"
	"image/color"
	"math"
	"testing"
)

func Test_Invert_will_return_transparent_when_passed_transparent(t *testing.T) {
	helpers.AssertColourEquals(color.RGBA{}, colour.Invert(color.RGBA{R: 255}), t)
}

func Test_Invert_will_invert_a_colour_with_transparency(t *testing.T) {
	helpers.AssertColourEquals(color.RGBA{R: 154, G: 155, B: 50, A: uint8(math.Floor(0.7 * 255))}, colour.Invert(color.RGBA{R: 101, G: 100, B: 205, A: uint8(math.Floor(0.7 * 255))}), t)
}

func Test_Invert_will_invert_a_colour(t *testing.T) {
	helpers.AssertColourEquals(color.RGBA{R: 51, G: 50, B: 155, A: 255}, colour.Invert(color.RGBA{R: 204, G: 205, B: 100, A: 255}), t)
}
