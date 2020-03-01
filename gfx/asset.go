package gfx

import (
	"github.com/veandco/go-sdl2/sdl"
)

// Asset represents an on-screen asset
type Asset struct {
	Name    string
	pos     Pos
	w, h    int
	scale   float32
	visible bool
	vp      *ViewPort
	texture *sdl.Texture
}

// NewAsset factory
func NewAsset(vp *ViewPort, name string) *Asset {
	return &Asset{
		Name:  name,
		pos:   Pos{},
		w:     0,
		h:     0,
		scale: 1.0,
		vp:    vp,
	}
}

// NewAssetFromTexture returns a new asset from a texture
func NewAssetFromTexture(vp *ViewPort, name string, tex *sdl.Texture) *Asset {

	_, _, w, h, _ := tex.Query()

	a := NewAsset(vp, name)
	a.w = int(w)
	a.h = int(h)
	a.texture = tex

	return a
}

// NewAssetFromBitmap converts and array of integer color values into a texture of the specified width.
func NewAssetFromBitmap(vp *ViewPort, name string, bitmap *Bitmap) (*Asset, error) {

	tex, err := bitmap.ToTexture(vp)
	if err != nil {
		return nil, err
	}

	a := NewAssetFromTexture(vp, name, tex)

	return a, nil
}

// Pos returns the position of the asset..
func (a *Asset) Pos() (x, y, z int32) {
	x, y, z = a.pos.X, a.pos.Y, a.pos.Z
	return
}

// SetPos sets the asset position ..
func (a *Asset) SetPos(x, y, z int32) {
	a.pos.X = x
	a.pos.Y = y
	a.pos.Z = z
}

// Scale returns the scale factor to use when drawing the asset
func (a *Asset) Scale() float32 {
	return a.scale
}

// SetScale sets the scale factor to use when drawing the asset
func (a *Asset) SetScale(scale float32) {
	a.scale = scale
}

// Texture an asset onto a rednerer ..
func (a *Asset) Texture() *sdl.Texture {
	return a.texture
}

// SetTexture an asset onto a rednerer ..
func (a *Asset) SetTexture(tex *sdl.Texture) {
	a.texture = tex
}

// IsVisible returns true if the asset should be visible on the screen
func (a *Asset) IsVisible() bool {
	return a.visible
}

// Show sets the asset to be visible on the screen
func (a *Asset) Show() {
	a.visible = true
}

// Hide set the asset to be hidden on the screen
func (a *Asset) Hide() {
	a.visible = false
}
