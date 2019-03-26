package colour_test

import (
	"colour"
	"colour/test"
	"image/color"
	"testing"
)

func Test_Shade_should_shade_the_provided_colour_with_white_by_the_given_percentage(t *testing.T) {
	helpers.AssertColourEquals(
		color.RGBA{R: 0, G: 0, B: 191, A: 255},
		colour.Shade(0.25, color.RGBA{R: 0, G: 0, B: 255, A: 255}),
		t)
}

func Test_Shade_should_return_transparent_when_passed_transparent(t *testing.T) {
	helpers.AssertColourEquals(
		color.RGBA{R: 0, G: 0, B: 0, A: 0},
		colour.Shade(0.25, color.RGBA{R: 0, G: 0, B: 0, A: 0}),
		t)
}
