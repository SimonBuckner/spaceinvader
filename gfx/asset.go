package gfx

import "github.com/veandco/go-sdl2/sdl"

// Asset represents an on-screen asset
type Asset struct {
	pos      Pos
	w, h     int
	scale    float32
	index    int
	vp       *ViewPort
	textures []*sdl.Texture
}

var _ Drawable = &Asset{}

// AssetFromBitmap converts and array of integer color values into a texture of the specified width.
func AssetFromBitmap(vp *ViewPort, bitmaps ...Bitmap) *Asset {

	asset := &Asset{
		pos:      Pos{},
		scale:    1.0,
		index:    0,
		vp:       vp,
		textures: make([]*sdl.Texture, len(bitmaps)),
	}

	for bmIndex, bm := range bitmaps {
		pixels := make([]byte, bm.Width*bm.Height*4)
		w := int32(bm.Width)
		h := int32(bm.Height)

		i := 0
		for _, pixel := range bm.Pixels {
			if i >= len(pixels) {
				break
			}
			rgba := HexColorToRGBA(pixel)
			pixels[i] = rgba.R
			i++
			pixels[i] = rgba.G
			i++
			pixels[i] = rgba.B
			i++
			pixels[i] = rgba.A
			i++
		}

		asset.textures[bmIndex], _ = vp.renderer.CreateTexture(sdl.PIXELFORMAT_RGBA8888, sdl.TEXTUREACCESS_STATIC, w, h)
		asset.textures[bmIndex].Update(nil, pixels, 4*bm.Width)
	}
	return asset
}

// Pos returns the position of the asset..
func (a *Asset) Pos() (x, y, z int32) {
	return a.pos.X, a.pos.Y, a.pos.Z
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
	return a.textures[a.index]
}

// CurrentIndex returns the index of the current texture
func (a *Asset) CurrentIndex() int {
	return a.index
}

// SetCurrent sets the current index of the texture to be displayed
func (a *Asset) SetCurrent(index int) {
	if index < 0 || index >= len(a.textures) {
		a.index = 0
		return
	}
	a.index = index
}
