package colour_test

import (
	"colour"
	"fmt"
	"image/color"
)

func Example_Mix() {
	// Take two RGBA values and return both mixed by 25%
	fmt.Printf("%+v\n", colour.Mix(0.25, color.RGBA{R: 255, G: 0, B: 0, A: 255}, color.RGBA{R: 0, G: 0, B: 255, A: 255}))

	// Output:
	// {R:63 G:0 B:191 A:255}
}
