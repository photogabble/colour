package colour_test

import (
	"colour"
	"colour/test"
	"image/color"
	"math"
	"testing"
)

func Test_Mix_should_mix_two_colours_by_a_weight_of_25_percent(t *testing.T) {
	// It should mix two colors with by a weight of 25%
	helpers.AssertColourEquals(
		color.RGBA{R: 63, G: 0, B: 191, A: 255},
		colour.Mix(0.25, color.RGBA{R: 255, G: 0, B: 0, A: 255}, color.RGBA{R: 0, G: 0, B: 255, A: 255}),
		t)
}

func Test_Mix_should_mix_two_colours_with_opacity_lower_than_one(t *testing.T) {
	// It should mix two colors with opacity lower than 1
	helpers.AssertColourEquals(
		color.RGBA{R: 63, G: 0, B: 191, A: uint8(math.Floor(float64(255 * 0.75)))},
		colour.Mix(0.5, color.RGBA{R: 255, G: 0, B: 0, A: uint8(255 / 2)}, color.RGBA{R: 0, G: 0, B: 255, A: 255}),
		t)
}

func Test_Mix_should_return_colourB_when_passed_transparent_colourA(t *testing.T) {
	// It should return colourB when passed transparent colourA
	helpers.AssertColourEquals(
		color.RGBA{R: 0, G: 255, B: 0, A: 255},
		colour.Mix(0.5, color.RGBA{R: 255, G: 0, B: 0, A: 0}, color.RGBA{R: 0, G: 255, B: 0, A: 255}),
		t)
}

func Test_Mix_should_return_colourA_when_passed_transparent_colourB(t *testing.T) {
	// It should return colourA when passed transparent colourB
	helpers.AssertColourEquals(
		color.RGBA{R: 50, G: 50, B: 50, A: 255},
		colour.Mix(0.5, color.RGBA{R: 50, G: 50, B: 50, A: 255}, color.RGBA{R: 0, G: 255, B: 0, A: 0}),
		t)
}

func Test_Mix_should_return_transparent_when_passed_transparent_for_both_colours(t *testing.T) {
	// It should return transparent when passed transparent for both colours
	helpers.AssertColourEquals(
		color.RGBA{R: 0, G: 0, B: 0, A: 0},
		colour.Mix(0.5, color.RGBA{R: 50, G: 50, B: 50, A: 0}, color.RGBA{R: 0, G: 255, B: 0, A: 0}),
		t)
}
