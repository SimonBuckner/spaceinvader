package gfx

import (
	"fmt"

	"github.com/veandco/go-sdl2/mix"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

const (
	rMask = 0xFF000000
	gMask = 0x00FF0000
	bMask = 0x0000FF00
	aMask = 0x000000FF
)

// Drawable represents nd object that can be drawn by a renderer
type Drawable interface {
	Texture() *sdl.Texture
	Pos() (x, y, z int32)
	Scale() float32
	IsVisible() bool
}

// Pos represents the position of an item
type Pos struct {
	X, Y, Z int32
}

// Bitmap reqpresents a basic bitmap
type Bitmap struct {
	Pitch             int
	Transparency      bool
	Pixels            []int
	TransparentColour sdl.Color
}

func init() {
	sdl.LogSetAllPriority(sdl.LOG_PRIORITY_VERBOSE)

	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}

	if err := ttf.Init(); err != nil {
		panic(err)
	}

	if err := mix.Init(mix.INIT_OGG); err != nil {
		panic(err)
	}
}

// KeyboardHandler registers a function to handle keyboard event
type KeyboardHandler func(e *sdl.KeyboardEvent)

// MouseButtonHandler is called each time a mouse button event is triggered
type MouseButtonHandler func(event *sdl.MouseButtonEvent)

// MouseMotionHandler is called each time a mouse motion event is triggered
type MouseMotionHandler func(event *sdl.MouseMotionEvent)

// MouseWheelHandler is called each time a mouse wheel event is triggered
type MouseWheelHandler func(event *sdl.MouseWheelEvent)

// UpdateHandler is called once each game loop to update game assets before rendering
type UpdateHandler func(vp *ViewPort, ticks uint32)

// State represents the current state of the game
type State interface {
	IsRunning() bool
}

// HexColorToRGBA converts a colour stored in an int to RGBA values
func HexColorToRGBA(color int) *sdl.Color {

	r := uint8((color & rMask) >> 24)
	g := uint8((color & gMask) >> 16)
	b := uint8((color & bMask) >> 8)
	a := uint8(color & aMask)
	// fmt.Printf("R:%2x G: %2x B: %2x A: %2x\n", r, g, b, a)
	return &sdl.Color{
		R: r,
		G: g,
		B: b,
		A: a,
	}
}

// BitmapAtlas represents an array of bitmap addressable by a key
type BitmapAtlas struct {
	Bitmap
	Key        map[rune]AtlasCoord
	TileWidth  int
	TileHeight int
}

// AtlasCoord indicates the location in bitmap/texture atlas
type AtlasCoord struct {
	X, Y int
}

// GetKeys returns a []rune of keys in the atlas
func (ba *BitmapAtlas) GetKeys() []rune {
	keys := make([]rune, len(ba.Key))
	i := 0
	for k := range ba.Key {
		keys[i] = k
	}
	return keys
}

// GetTile returns the bitmap tile for key speified
func (ba *BitmapAtlas) GetTile(key rune) (*Bitmap, error) {
	coord, ok := ba.Key[key]
	if !ok {
		return nil, keyNotFoundError(key)
	}
	tile := &Bitmap{
		Pitch:             ba.TileWidth,
		Transparency:      ba.Bitmap.Transparency,
		TransparentColour: ba.Bitmap.TransparentColour,
		Pixels:            make([]int, ba.TileWidth*ba.TileHeight),
	}

	i := 0
	rowStart := ((coord.Y * ba.TileHeight) * ba.Pitch) + (coord.X * ba.TileWidth)
	for y := 0; y < ba.TileHeight; y++ {
		j := rowStart
		for x := 0; x < ba.TileWidth; x++ {
			tile.Pixels[i] = ba.Bitmap.Pixels[j]
			i++
			j++
		}
		rowStart = rowStart + ba.Pitch
	}

	return tile, nil
}

func keyNotFoundError(key rune) error {
	return fmt.Errorf("key '%v'not found in Altas", key)
}
