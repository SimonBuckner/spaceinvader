package gfx

import (
	"github.com/veandco/go-sdl2/mix"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

// Pos represents the position of an item
type Pos struct {
	X, Y, Z int32
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
type UpdateHandler func(vp *ViewPort)

// State represents the current state of the game
type State interface {
	IsRunning() bool
}

// HexColorToRGB converts a colour stored in an int to RGBA values
func HexColorToRGBA(color int) *sdl.Color {

	r := uint8((color & 0b11111111000000000000000000000000) >> 24)
	g := uint8((color & 0b00000000111111110000000000000000) >> 16)
	b := uint8((color & 0b00000000000000001111111100000000) >> 8)
	a := uint8(color & 0b00000000000000000000000011111111)

	return &sdl.Color{
		R: r,
		G: g,
		B: b,
		A: a,
	}
}
