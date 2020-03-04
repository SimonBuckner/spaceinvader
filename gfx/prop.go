package gfx

import (
	"github.com/veandco/go-sdl2/sdl"
)

// Prop represents an on-screen item.
type Prop struct {
	name string
	Pos
	speed   Pos
	tex     *sdl.Texture
	stage   *Stage
	visible bool
	scale   float32

	w, h int
}

// NewProp returns a new prop
func NewProp(stage *Stage, name string, tex *sdl.Texture) *Prop {

	p := &Prop{
		name:  name,
		Pos:   Pos{},
		speed: Pos{},
		w:     0,
		h:     0,
		scale: stage.Scale(),
		stage: stage,
	}
	if tex != nil {
		_, _, w, h, _ := tex.Query()
		p.w = int(w)
		p.h = int(h)
		p.tex = tex
	}
	return p
}

// Name returns the name of the prop
func (p *Prop) Name() string {
	return p.name
}

// Scale returns the scale fscene to use when drawing the p
func (p *Prop) Scale() float32 {
	return p.scale
}

// SetScale sets the scale fscene to use when drawing the p
func (p *Prop) SetScale(scale float32) {
	p.scale = scale
}

// Texture an p onto a rednerer ..
func (p *Prop) Texture() *sdl.Texture {
	return p.tex
}

// SetTexture an p onto a rednerer ..
func (p *Prop) SetTexture(tex *sdl.Texture) {
	p.tex = tex
}

// Visible returns true if the p should be visible on the screen
func (p *Prop) Visible() bool {
	return p.visible
}

// SetVisible set the visibility of a Prop
func (p *Prop) SetVisible(visible bool) {
	p.visible = visible
}
