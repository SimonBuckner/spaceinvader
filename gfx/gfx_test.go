package gfx

import "testing"

func TestHexColorToRGBA(t *testing.T) {
	c := 0x11223344
	rgba := HexColorToRGBA(c)

	if rgba.R != 0x11 {
		t.Errorf("unexpected value for R; expecting 0x11, got %#x", rgba.R)
	}
	if rgba.G != 0x22 {
		t.Errorf("unexpected value for R; expecting 0x11, got %#x", rgba.G)
	}
	if rgba.B != 0x33 {
		t.Errorf("unexpected value for R; expecting 0x11, got %#x", rgba.B)
	}
	if rgba.A != 0x44 {
		t.Errorf("unexpected value for R; expecting 0x11, got %#x", rgba.A)
	}
}
