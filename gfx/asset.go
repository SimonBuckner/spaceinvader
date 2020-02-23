package gfx

import "github.com/veandco/go-sdl2/sdl"

// Asset represents an on-screen asset
type Asset struct {
	pos  Pos
	vp   *ViewPort
	w, h int
	tex  *sdl.Texture
}

// Texture an asset onto a rednerer ..
func (a *Asset) Texture() *sdl.Texture {
	return a.tex
}

// SetPos sets the asset position ..
func (a *Asset) SetPos(x, y, z int32) {
	a.pos.X = x
	a.pos.Y = y
	a.pos.Z = z
}

// AssetFromBitmap converts and array of integer color values into a texture of the specified width.
func AssetFromBitmap(vp *ViewPort, bm []int, width, height int) *Asset {

	pixels := make([]byte, len(bm)*4)
	w := int32(width)
	h := int32(len(bm) / width)

	i := 0
	for _, pixel := range bm {
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

	tex, _ := vp.renderer.CreateTexture(sdl.PIXELFORMAT_RGBA8888, sdl.TEXTUREACCESS_STATIC, w, h)
	tex.Update(nil, pixels, 4*width)

	return &Asset{
		pos: Pos{},
		vp:  vp,
		tex: tex,
	}
}
