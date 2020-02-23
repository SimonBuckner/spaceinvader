package gfx

import "github.com/veandco/go-sdl2/sdl"

// Asset represents an on-screen asset
type Asset struct {
	Pos
	Tex *sdl.Texture
}

// BitmapToTexture converts and array of integer color values into a texture of the specified width.
func BitmapToTexture(r *sdl.Renderer, bm []int, width int) *sdl.Texture {

	pixels := make([]byte, len(bm)*4)
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

	w := int32(width)
	h := int32(len(bm) / width)
	tex, _ := r.CreateTexture(sdl.PIXELFORMAT_RGBA8888, sdl.TEXTUREACCESS_STATIC, w, h)

	tex.Update(nil, pixels, 4*width)

	return tex
}
