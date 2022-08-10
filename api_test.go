package skip32_test

import (
	"testing"

	"github.com/jmhobbs/skip32"
)

func TestKeyFromSlice(t *testing.T) {
	t.Run("does not pad correctly sized slices", func(t *testing.T) {
		input := []byte{0x11, 0x22, 0x33, 0x44, 0x55, 0x66, 0x77, 0x88, 0x99, 0x11}
		expected := [10]byte{0x11, 0x22, 0x33, 0x44, 0x55, 0x66, 0x77, 0x88, 0x99, 0x11}
		actual := skip32.KeyFromSlice(input)
		if actual != expected {
			t.Errorf("key did not match\nexpected: %v\n  actual: %v", expected, actual)
		}
	})

	t.Run("takes the first 10 bytes of an oversized slice", func(t *testing.T) {
		input := []byte{0x11, 0x22, 0x33, 0x44, 0x55, 0x66, 0x77, 0x88, 0x99, 0x11, 0x22, 0x33}
		expected := [10]byte{0x11, 0x22, 0x33, 0x44, 0x55, 0x66, 0x77, 0x88, 0x99, 0x11}
		actual := skip32.KeyFromSlice(input)
		if actual != expected {
			t.Errorf("key did not match\nexpected: %v\n  actual: %v", expected, actual)
		}
	})

	t.Run("correctly left pads undersized slices", func(t *testing.T) {
		input := []byte{0x11, 0x22, 0x33, 0x44, 0x55, 0x66}
		expected := [10]byte{0x00, 0x00, 0x00, 0x00, 0x11, 0x22, 0x33, 0x44, 0x55, 0x66}
		actual := skip32.KeyFromSlice(input)
		if actual != expected {
			t.Errorf("key did not match\nexpected: %v\n  actual: %v", expected, actual)
		}
	})
}
