package hello

import "testing"

func TestMultiply(t *testing.T) {
	if multiply(2, 2) != 4 {
		t.Fatal("Multiply failed")
	}
}
