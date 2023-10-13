package calculator_test

import (
	"testing"

	"github.com/PacktPublishing/Test-Driven-Development-in-Go/chapter02/calculator"
)

func TestAdd(t *testing.T) {

	e := calculator.Engine{}
	x, y := 2.6, 3.4
	want := 6.0

	got := e.Add(x, y)

	if got != 6.0 {
		t.Errorf("Add (%.2f, %.2f) incorrect, got: %.2f, want: %.2f", x, y, got, want)
	}
}
