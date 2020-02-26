package gfx

import (
	"fmt"
	"sort"

	"github.com/veandco/go-sdl2/sdl"
)

// Bitmap reqpresents a basic bitmap
type Bitmap struct {
	Pitch        int
	Transparency bool
	Pixels       []int
	// TransparentColour sdl.Color
}

// AtlasCoord indicates the location in bitmap/texture atlas
type AtlasCoord struct {
	X, Y int
}

// stringNotFoundError inndicates the key was not found in the atlas
func stringNotFoundError(key string) error {
	return fmt.Errorf("key '%v'not found in Altas", string(key))
}

// ToAsset returns an asset for the texture
func (b *Bitmap) ToAsset(vp *ViewPort) (*Asset, error) {
	asset := &Asset{
		pos:      Pos{},
		scale:    1.0,
		index:    0,
		vp:       vp,
		textures: make([]*sdl.Texture, 1),
	}
	tex, err := b.ToTexture(vp)
	if err != nil {
		return nil, err
	}
	asset.textures[0] = tex
	return asset, nil
}

// ToTexture returnes a texture from a bitmap
func (b *Bitmap) ToTexture(vp *ViewPort) (*sdl.Texture, error) {

	w := b.Pitch
	h := len(b.Pixels) / b.Pitch
	if w*h != len(b.Pixels) {
		return nil, fmt.Errorf("bitmap has the wrong number of pixels; expected %3d, got %3d", w*h, len(b.Pixels))
	}

	pixels := make([]byte, len(b.Pixels)*4)
	i := 0
	for _, p := range b.Pixels {
		c := HexColorToRGBA(p)
		pixels[i] = c.R
		i++
		pixels[i] = c.G
		i++
		pixels[i] = c.B
		i++
		pixels[i] = c.A
		i++
	}
	tex, err := vp.Renderer().CreateTexture(sdl.PIXELFORMAT_RGBA8888, sdl.TEXTUREACCESS_STATIC, int32(w), int32(h))
	if err != nil {
		return nil, err
	}
	if b.Transparency {
		tex.SetBlendMode(sdl.BLENDMODE_BLEND)
	} else {
		tex.SetBlendMode(sdl.BLENDMODE_NONE)
	}
	tex.Update(nil, pixels, w*4)
	return tex, nil
}

// BitmapAtlas represents an array of bitmap addressable by a key
type BitmapAtlas struct {
	Bitmap
	Keys       map[string]AtlasCoord
	Pitch      int // The width of the atlas
	TileWidth  int // The width of a tile
	TileHeight int // The height of a tile
}

// GetKeys returns a []rune of keys in the atlas
func (a *BitmapAtlas) GetKeys() []string {
	keys := make([]string, len(a.Keys))
	i := 0
	for k := range a.Keys {
		keys[i] = k
		i++
	}
	sort.Strings(keys)
	return keys
}

// GetTileBitmap returns the bitmap tile for key speified
func (a *BitmapAtlas) GetTileBitmap(key string) (*Bitmap, error) {
	if len(a.Bitmap.Pixels) != len(a.Keys)*a.TileHeight*a.TileWidth {
		return nil, fmt.Errorf("bitmap atlas, wrong number of pixels")
	}
	coord, ok := a.Keys[key]
	if !ok {
		return nil, stringNotFoundError(key)
	}
	tile := &Bitmap{
		Pitch:        a.TileWidth,
		Transparency: a.Bitmap.Transparency,
		Pixels:       make([]int, a.TileWidth*a.TileHeight),
	}

	i := 0
	rowStart := (coord.Y * a.TileHeight * a.Pitch) + (coord.X * a.TileWidth)
	for y := 0; y < a.TileHeight; y++ {
		j := rowStart
		for x := 0; x < a.TileWidth; x++ {
			if j >= len(a.Bitmap.Pixels) {
				return nil, fmt.Errorf("error getting bitmap with key '%v'; out of pixels (j=%d/len=%d)", key, j, len(a.Bitmap.Pixels))
			}

			tile.Pixels[i] = a.Bitmap.Pixels[j]
			i++
			j++
		}
		rowStart = rowStart + a.Pitch
	}
	return tile, nil
}

// GetTileTexture returns a texuter for the bitmap at the position associated witht he key
func (a *BitmapAtlas) GetTileTexture(vp *ViewPort, key string) (*sdl.Texture, error) {
	bm, err1 := a.GetTileBitmap(key)
	if err1 != nil {
		return nil, err1
	}
	tex, err2 := bm.ToTexture(vp)
	if err2 != nil {
		return nil, err2
	}
	return tex, nil
}
