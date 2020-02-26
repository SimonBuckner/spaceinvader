package gfx

import (
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

func init() {
	sdl.LogSetAllPriority(sdl.LOG_PRIORITY_WARN)

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

	return &sdl.Color{
		R: r,
		G: g,
		B: b,
		A: a,
	}
}
