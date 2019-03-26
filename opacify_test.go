package colour_test

import (
	"colour"
	"colour/test"
	"image/color"
	"math"
	"testing"
)

func Test_Opacify_will_return_transparent_when_passed_transparent(t *testing.T) {
	helpers.AssertColourEquals(color.RGBA{}, colour.Opacify(1, color.RGBA{R: 255}), t)
}

func Test_Opacify_will_increase_the_opacity_of_fff_by_point_one_and_still_be_one(t *testing.T) {
	helpers.AssertColourEquals(color.RGBA{R: 255, G: 255, B: 255, A: 255}, colour.Opacify(1, color.RGBA{R: 255, G: 255, B: 255, A: 255}), t)
}

func Test_Opacify_will_increase_the_opacity_of_colours(t *testing.T) {
	helpers.AssertColourEquals(color.RGBA{R: 101, G: 100, B: 205, A: uint8(math.Floor(0.8 * 255))}, colour.Opacify(0.1, color.RGBA{R: 101, G: 100, B: 205, A: uint8(math.Floor(0.7 * 255))}), t)
	helpers.AssertColourEquals(color.RGBA{R: 255, G: 0, B: 0, A: 255}, colour.Opacify(0.5, color.RGBA{R: 255, G: 0, B: 0, A: uint8(math.Floor(0.5 * 255))}), t)
}
