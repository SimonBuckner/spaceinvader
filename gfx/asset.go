package gfx

import (
	"github.com/veandco/go-sdl2/sdl"
)

// Asset represents an on-screen asset
type Asset struct {
	pos      Pos
	w, h     int
	scale    float32
	index    int
	visible  bool
	vp       *ViewPort
	textures []*sdl.Texture
}

var _ Drawable = &Asset{}

// SurfaceFromBitMap returns an array of surfaces from the bitmap data
func SurfaceFromBitMap(bitmaps ...Bitmap) []*sdl.Surface {
	pf := sdl.PixelFormat{
		Format:        sdl.PIXELFORMAT_RGBA8888,
		Palette:       nil,
		BitsPerPixel:  32,
		BytesPerPixel: 4,
		Rmask:         rMask,
		Gmask:         gMask,
		Bmask:         bMask,
		Amask:         aMask,
	}
	surfaces := make([]*sdl.Surface, len(bitmaps))
	for i := 0; i < len(bitmaps); i++ {
		bm := bitmaps[i]
		tc := bm.TransparentColour
		ck := sdl.MapRGBA(&pf, tc.R, tc.G, tc.B, tc.A)

		w := int32(bm.Width)
		h := int32(len(bm.Pixels)) / w

		surface, _ := sdl.CreateRGBSurface(0, w, h, 32, rMask, gMask, bMask, aMask)
		surface.Lock()
		pixels := surface.Pixels()
		j := 0
		for _, pixel := range bm.Pixels {
			if j >= len(pixels) {
				break
			}
			rgba := HexColorToRGBA(pixel)
			pixels[j] = rgba.R
			j++
			pixels[j] = rgba.G
			j++
			pixels[j] = rgba.B
			j++
			pixels[j] = rgba.A
			j++
		}
		surface.Unlock()
		if bm.Transparency {
			surface.SetColorKey(true, ck)
			surface.SetBlendMode(sdl.BLENDMODE_BLEND)
		} else {
			surface.SetBlendMode(sdl.BLENDMODE_NONE)
		}
		surfaces[i] = surface
	}
	return surfaces
}

// AssetFromBitmap converts and array of integer color values into a texture of the specified width.
func AssetFromBitmap(vp *ViewPort, bitmaps ...Bitmap) *Asset {

	asset := &Asset{
		pos:      Pos{},
		scale:    1.0,
		index:    0,
		vp:       vp,
		textures: make([]*sdl.Texture, len(bitmaps)),
	}

	surfaces := SurfaceFromBitMap(bitmaps...)
	for i := 0; i < len(surfaces); i++ {
		tex, _ := vp.renderer.CreateTextureFromSurface(surfaces[i])
		bm, _ := surfaces[i].GetBlendMode()
		tex.SetBlendMode(bm)
		asset.textures[i] = tex
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
