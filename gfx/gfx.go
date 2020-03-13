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
	Vec3() (x, y, z int32)
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

// Vec3 represents the position of an item
type Vec3 struct {
	X, Y, Z float32
}

// Int32 returns the position as int32
func (p *Vec3) Int32() (x, y, z int32) {
	return int32(p.X), int32(p.Y), int32(p.Z)
}

// Float32 returns the position as float32
func (p *Vec3) Float32() (x, y, z float32) {
	return p.X, p.Y, p.Z
}

// Vec3Int32 returns the pos as int32
func (p *Vec3) Vec3Int32() (x, y, z int32) {
	return int32(p.X), int32(p.Y), int32(p.Z)
}

// Vec3Float32 returns the pos as float32
func (p *Vec3) Vec3Float32() (x, y, z float32) {
	return p.X, p.Y, p.Z
}

// MoveVec3 moves the position by +/- x, y, z
func (p *Vec3) MoveVec3(x, y, z float32) {
	p.X += x
	p.Y += y
	p.Z += z
}

// MoveX moves the X pos by +/- x
func (p *Vec3) MoveX(x float32) {
	p.X += x
}

// MoveY moves the Y pos by +/- x
func (p *Vec3) MoveY(y float32) {
	p.Y += y
}

// MoveZ moves the Z pos by +/- x
func (p *Vec3) MoveZ(z float32) {
	p.Z += z
}

// SetX sets the X pos
func (p *Vec3) SetX(x float32) {
	p.X = x
}

// SetY sets the Y pos
func (p *Vec3) SetY(y float32) {
	p.Y = y
}

// SetZ sets the Z pos
func (p *Vec3) SetZ(z float32) {
	p.Z = z
}

// Set sets the pos
func (p *Vec3) Set(x, y, z float32) {
	p.X = x
	p.Y = y
	p.Z = z
}

// SetInt32 sets the pos
func (p *Vec3) SetInt32(x, y, z int32) {
	p.X = float32(x)
	p.Y = float32(y)
	p.Z = float32(z)
}

// TransformXYZFunc transforms game X, Y, Z to diplay X, Y, Z
type TransformXYZFunc func(pos Vec3) (tX, tY, tZ int32)

// Prop represents an on-screen item.
type Prop struct {
	Name    string
	Pos     Vec3
	Texture *sdl.Texture
	Scale   float32

	Width  int32
	Height int32
	TransformXYZFunc
}

// NewProp factory
func NewProp(name string, texture *sdl.Texture, transformFunc TransformXYZFunc) *Prop {
	return &Prop{
		Name:             name,
		Texture:          texture,
		Pos:              Vec3{},
		Scale:            1,
		TransformXYZFunc: transformFunc,
	}
}

// Draw renders the props texture to the screen
func (p *Prop) Draw(renderer *sdl.Renderer) {
	if p.Texture == nil {
		return
	}

	_, _, w, h, _ := p.Texture.Query()
	var x, y int32
	if p.TransformXYZFunc != nil {
		x, y, _ = p.TransformXYZFunc(p.Pos)
	} else {
		fmt.Println("no transform")
		x, y, _ = p.Pos.Int32()
	}

	dstRect := &sdl.Rect{
		X: x,
		Y: y,
		W: int32(float32(w) * p.Scale),
		H: int32(float32(h) * p.Scale),
	}
	renderer.Copy(p.Texture, nil, dstRect)
}

const (
	keyUp   uint8 = 0
	keyDown uint8 = 1
)

// KBState represents the state of the keyboard
type KBState struct {
	state  []uint8
	state1 []uint8
}

// NewKBState fsceney
func NewKBState() *KBState {
	kb := &KBState{}

	kb.state = sdl.GetKeyboardState()
	kb.state1 = make([]uint8, len(kb.state))

	return kb
}

// Refresh the keyboard state
func (kb *KBState) Refresh() {
	for i, v := range kb.state {
		kb.state1[i] = v
	}
	kb.state = sdl.GetKeyboardState()
}

// OnKeyDown returns true when the key state changes from up to down
func (kb *KBState) OnKeyDown(key uint8) bool {
	return kb.state[key] == keyDown && kb.state1[key] == keyUp
}

// OnKeyUp returns true when the key state changes from up to down
func (kb *KBState) OnKeyUp(key uint8) bool {
	return kb.state[key] == keyUp && kb.state1[key] == keyDown
}

// IsKeyDown returns true if the key state is down
func (kb *KBState) IsKeyDown(key uint8) bool {
	return kb.state[key] == keyDown
}

// IsKeyUp returns true if the key state is up
func (kb *KBState) IsKeyUp(key uint8) bool {
	return kb.state[key] == keyUp
}
