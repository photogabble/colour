package helpers

import (
	"image/color"
	"testing"
)

func AssertFloat64Equals(a, b float64, t *testing.T) {
	if a != b {
		t.Errorf("Expecting %+v, got '%+v'\n", a, b)
	}
}

func AssertColourEquals(a, b color.RGBA, t *testing.T) {
	if a.R != b.R || a.G != b.G || a.B != b.B || a.A != b.A {
		t.Errorf("Expecting %+v, got '%+v'\n", a, b)
	}
}

func AssertStringEquals(a, b string, t *testing.T) {
	if a != b {
		t.Errorf("Expecting %+v, got '%+v'\n", a, b)
	}
}
