package gfx

import (
	"github.com/veandco/go-sdl2/sdl"
)

// Prop represents an on-screen item.
type Prop struct {
	name    string
	tex     *sdl.Texture
	stage   *Stage
	visible bool
	scale   float32
	pos     Pos
	w, h    int
}

// NewProp returns a new prop
func NewProp(stage *Stage, name string, tex *sdl.Texture) *Prop {

	prop := &Prop{
		name:  name,
		pos:   Pos{},
		w:     0,
		h:     0,
		scale: stage.Scale(),
		stage: stage,
	}
	if tex != nil {
		_, _, w, h, _ := tex.Query()
		prop.w = int(w)
		prop.h = int(h)
		prop.tex = tex
	}
	return prop
}

// Name returns the name of the prop
func (prop *Prop) Name() string {
	return prop.name
}

// Pos returns the position of the proprop..
func (prop *Prop) Pos() (x, y, z int32) {
	x, y, z = prop.pos.X, prop.pos.Y, prop.pos.Z
	return
}

// SetPos sets the prop position ..
func (prop *Prop) SetPos(x, y, z int32) {
	prop.pos.X = x
	prop.pos.Y = y
	prop.pos.Z = z
}

// Scale returns the scale factor to use when drawing the prop
func (prop *Prop) Scale() float32 {
	return prop.scale
}

// SetScale sets the scale factor to use when drawing the prop
func (prop *Prop) SetScale(scale float32) {
	prop.scale = scale
}

// Texture an prop onto a rednerer ..
func (prop *Prop) Texture() *sdl.Texture {
	return prop.tex
}

// SetTexture an prop onto a rednerer ..
func (prop *Prop) SetTexture(tex *sdl.Texture) {
	prop.tex = tex
}

// Visible returns true if the prop should be visible on the screen
func (prop *Prop) Visible() bool {
	return prop.visible
}

// SetVisible set the visibility of a Prop
func (prop *Prop) SetVisible(visible bool) {
	prop.visible = visible
}
