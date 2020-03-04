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

// Pos represents the position of an item
type Pos struct {
	X, Y, Z float32
}

// Int32 returns the position as int32
func (p *Pos) Int32() (x, y, z int32) {
	return int32(p.X), int32(p.Y), int32(p.Z)
}

// Float32 returns the position as float32
func (p *Pos) Float32() (x, y, z float32) {
	return p.X, p.Y, p.Z
}

// PosInt32 returns the pos as int32
func (p *Pos) PosInt32() (x, y, z int32) {
	return int32(p.X), int32(p.Y), int32(p.Z)
}

// PosFloat32 returns the pos as float32
func (p *Pos) PosFloat32() (x, y, z float32) {
	return p.X, p.Y, p.Z
}

// MovePos moves the position by +/- x, y, z
func (p *Pos) MovePos(x, y, z float32) {
	p.X += x
	p.Y += y
	p.Z += z
}

// MoveX moves the X pos by +/- x
func (p *Pos) MoveX(x float32) {
	p.X += x
}

// MoveY moves the Y pos by +/- x
func (p *Pos) MoveY(y float32) {
	p.Y += y
}

// MoveZ moves the Z pos by +/- x
func (p *Pos) MoveZ(z float32) {
	p.Z += z
}

// SetX sets the X pos
func (p *Pos) SetX(x float32) {
	p.X = x
}

// SetY sets the Y pos
func (p *Pos) SetY(y float32) {
	p.Y = y
}

// SetZ sets the Z pos
func (p *Pos) SetZ(z float32) {
	p.Z = z
}

// Set sets the pos
func (p *Pos) Set(x, y, z float32) {
	p.X = x
	p.Y = y
	p.Z = z
}

// SetInt32 sets the pos
func (p *Pos) SetInt32(x, y, z int32) {
	p.X = float32(x)
	p.Y = float32(y)
	p.Z = float32(z)
}
