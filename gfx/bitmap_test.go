package gfx

import (
	"testing"

	"github.com/veandco/go-sdl2/sdl"
)

var goodAtlas = BitmapAtlas{
	Keys: map[AtlasKey]AtlasCoord{
		AtlasKey("A"): AtlasCoord{X: 0, Y: 0},
		AtlasKey("B"): AtlasCoord{X: 1, Y: 0},
		AtlasKey("C"): AtlasCoord{X: 0, Y: 1},
		AtlasKey("D"): AtlasCoord{X: 1, Y: 1},
	},
	Pitch:      4,
	TileWidth:  2,
	TileHeight: 2,
	Bitmap: Bitmap{
		Pitch:             4,
		Transparency:      false,
		TransparentColour: sdl.Color{R: 0, G: 0, B: 0, A: 0},
		Pixels: []int{
			100, 101, 200, 201,
			110, 111, 210, 211,
			300, 301, 400, 401,
			310, 311, 410, 411,
		},
	},
}

func TestBitmapAtlas(t *testing.T) {

	keys := goodAtlas.GetKeys()
	if len(keys) != 4 {
		t.Errorf("expecting 4 keys, got %d", len(keys))
	}

	{
		_, err := goodAtlas.GetTile(AtlasKey("z"))
		if err == nil {
			t.Errorf("expected error getting tile 'z', get nil")
		}
	}
	{
		tile, err := goodAtlas.GetTile(AtlasKey("A"))
		if err != nil {
			t.Errorf("unexpected error getting tile 'A', expected nil, got %v", err)
		}
		if tile.Pitch != 2 {
			t.Errorf("unexpected value geetting tile 'A'; expecting 2, got %d", tile.Pitch)
		}
		if len(tile.Pixels) != 4 {
			t.Errorf("unexpected length getting tile 'A', expected 4, got %v", len(tile.Pixels))
		}
		if tile.Pixels[0] != 100 {
			t.Errorf("unexpected pixel value geetting tile 'A', index 0; expecting 100, got %d", tile.Pixels[0])
		}
		if tile.Pixels[1] != 101 {
			t.Errorf("unexpected pixel value geetting tile 'A', index 1; expecting 101, got %d", tile.Pixels[1])
		}
		if tile.Pixels[2] != 110 {
			t.Errorf("unexpected pixel value geetting tile 'A', index 2; expecting 110, got %d", tile.Pixels[2])
		}
		if tile.Pixels[3] != 111 {
			t.Errorf("unexpected pixel value geetting tile 'A', index 3; expecting 111, got %d", tile.Pixels[3])
		}
	}
	{
		tile, err := goodAtlas.GetTile(AtlasKey("B"))
		if err != nil {
			t.Errorf("unexpected error getting tile 'B', expected nil, got %v", err)
		}
		if tile.Pitch != 2 {
			t.Errorf("unexpected value geetting tile 'B'; expecting 2, got %d", tile.Pitch)
		}
		if len(tile.Pixels) != 4 {
			t.Errorf("unexpected length getting tile 'B', expected 4, got %v", len(tile.Pixels))
		}
		if tile.Pixels[0] != 200 {
			t.Errorf("unexpected pixel value geetting tile 'B', index 0; expecting 200, got %d", tile.Pixels[0])
		}
		if tile.Pixels[1] != 201 {
			t.Errorf("unexpected pixel value geetting tile 'B', index 1; expecting 201, got %d", tile.Pixels[1])
		}
		if tile.Pixels[2] != 210 {
			t.Errorf("unexpected pixel value geetting tile 'B', index 2; expecting 210, got %d", tile.Pixels[2])
		}
		if tile.Pixels[3] != 211 {
			t.Errorf("unexpected pixel value geetting tile 'B', index 3; expecting 211, got %d", tile.Pixels[3])
		}
	}
	{
		tile, err := goodAtlas.GetTile(AtlasKey("C"))
		if err != nil {
			t.Errorf("unexpected error getting tile 'C', expected nil, got %v", err)
		}
		if tile.Pitch != 2 {
			t.Errorf("unexpected value geetting tile 'C'; expecting 2, got %d", tile.Pitch)
		}
		if len(tile.Pixels) != 4 {
			t.Errorf("unexpected length getting tile 'C', expected 4, got %v", len(tile.Pixels))
		}
		if tile.Pixels[0] != 300 {
			t.Errorf("unexpected pixel value geetting tile 'C', index 0; expecting 300, got %d", tile.Pixels[0])
		}
		if tile.Pixels[1] != 301 {
			t.Errorf("unexpected pixel value geetting tile 'C', index 1; expecting 301, got %d", tile.Pixels[1])
		}
		if tile.Pixels[2] != 310 {
			t.Errorf("unexpected pixel value geetting tile 'C', index 2; expecting 310, got %d", tile.Pixels[2])
		}
		if tile.Pixels[3] != 311 {
			t.Errorf("unexpected pixel value geetting tile 'C', index 3; expecting 311, got %d", tile.Pixels[3])
		}
	}
	{
		tile, err := goodAtlas.GetTile(AtlasKey("D"))
		if err != nil {
			t.Errorf("unexpected error getting tile 'D', expected nil, got %v", err)
		}
		if tile.Pitch != 2 {
			t.Errorf("unexpected value geetting tile 'D'; expecting 2, got %d", tile.Pitch)
		}
		if len(tile.Pixels) != 4 {
			t.Errorf("unexpected length getting tile 'D', expected 4, got %v", len(tile.Pixels))
		}
		if tile.Pixels[0] != 400 {
			t.Errorf("unexpected pixel value geetting tile 'D', index 0; expecting 400, got %d", tile.Pixels[0])
		}
		if tile.Pixels[1] != 401 {
			t.Errorf("unexpected pixel value geetting tile 'D', index 1; expecting 401, got %d", tile.Pixels[1])
		}
		if tile.Pixels[2] != 410 {
			t.Errorf("unexpected pixel value geetting tile 'D', index 2; expecting 410, got %d", tile.Pixels[2])
		}
		if tile.Pixels[3] != 411 {
			t.Errorf("unexpected pixel value geetting tile 'D', index 3; expecting 411, got %d", tile.Pixels[3])
		}
	}
}

var badLenghtAtlas = BitmapAtlas{
	Keys: map[AtlasKey]AtlasCoord{
		AtlasKey("A"): AtlasCoord{X: 0, Y: 0},
		AtlasKey("B"): AtlasCoord{X: 1, Y: 0},
		AtlasKey("C"): AtlasCoord{X: 0, Y: 1},
		AtlasKey("D"): AtlasCoord{X: 1, Y: 1},
	},
	Pitch:      4,
	TileWidth:  2,
	TileHeight: 2,
	Bitmap: Bitmap{
		Pitch:             4,
		Transparency:      false,
		TransparentColour: sdl.Color{R: 0, G: 0, B: 0, A: 0},
		Pixels: []int{
			100, 101, 200, 201,
			110, 111, 210, 211,
			300, 301, 400, 401,
			310, 311, 410,
		},
	},
}

func TestBitmapAtlasBadLnegth(t *testing.T) {

	keys := badLenghtAtlas.GetKeys()
	if len(keys) != 4 {
		t.Errorf("expecting 4 keys, got %d", len(keys))
	}

	{
		_, err := badLenghtAtlas.GetTile(AtlasKey("z"))
		if err == nil {
			t.Errorf("expected error getting tile 'z', get nil")
		}
	}
	{
		_, err := badLenghtAtlas.GetTile(AtlasKey("A"))
		if err == nil {
			t.Errorf("unexpected success getting tile 'A', expected err, got nil")
		}
	}
	{
		_, err := badLenghtAtlas.GetTile(AtlasKey("D"))
		if err == nil {
			t.Errorf("unexpected success getting tile 'D', expected err, got nil")
		}
	}
}
