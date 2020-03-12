package main

import (
	"testing"
)

func convertXYTest(w, h, x, y int32, scale float32) (int32, int32) {

	ow := float32(width) * scale
	oh := float32(height) * scale

	offsetX := (float32(w) - ow) / 2
	offsetY := (float32(h) - oh) / 2

	newX := float32(x) * scale
	newY := float32(y) * scale

	return int32(newX + offsetX), int32(newY + offsetY)
}

func TestConvertXY(t *testing.T) {
	tdt := []struct {
		w, h, x, y int32
		s          float32
		eX, eY     int32
	}{
		{1024, 768, 0, 0, 3.0, 176, 0},
		{1024, 768, 224, 0, 3.0, 848, 0},
		{1024, 768, 0, 256, 3.0, 0, 768},
		{1024, 768, 224, 256, 3.0, 848, 768},
	}
	for _, v := range tdt {
		x, y := convertXYTest(v.w, v.h, v.x, v.y, v.s)
		if x == v.eX {
			t.Errorf("for x:%d y:%d; expected x:%d but got x:%d", v.x, v.y, v.eX, x)
		}

		if x == v.eY {
			t.Errorf("for x:%d y:%d; expected x:%d but got x:%d", v.x, v.y, v.eY, y)
		}
	}
}
