package gfx

import (
	"fmt"
	"sort"

	"github.com/veandco/go-sdl2/sdl"
)

// Bitmap reqpresents a basic bitmap
type Bitmap struct {
	Pitch  int
	Pixels []int
}

// ToTexture returnes a texture from a bitmap
func (bm *Bitmap) ToTexture(vp *ViewPort) (*sdl.Texture, error) {

	w := bm.Pitch
	h := len(bm.Pixels) / bm.Pitch
	if w*h != len(bm.Pixels) {
		return nil, fmt.Errorf("bitmap has the wrong number of pixels; expected %4d, got %4d", w*h, len(bm.Pixels))
	}

	pixels := make([]byte, len(bm.Pixels)*4)
	i := 0
	for _, p := range bm.Pixels {
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
	tex.Update(nil, pixels, w*4)
	return tex, nil
}

// AtlasCoord indicates the location in bitmap/texture atlas
type AtlasCoord struct {
	X, Y int
}

// stringNotFoundError inndicates the key was not found in the atlas
func stringNotFoundError(key string) error {
	return fmt.Errorf("key '%v'not found in Altas", string(key))
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
func (bma *BitmapAtlas) GetKeys() []string {
	keys := make([]string, len(bma.Keys))
	i := 0
	for k := range bma.Keys {
		keys[i] = k
		i++
	}
	sort.Strings(keys)
	return keys
}

// GetBitmap returns the bitmap tile for key speified
func (bma *BitmapAtlas) GetBitmap(key string) (*Bitmap, error) {
	if len(bma.Bitmap.Pixels) != len(bma.Keys)*bma.TileHeight*bma.TileWidth {
		return nil, fmt.Errorf("bitmap atlas, wrong number of pixels")
	}
	coord, ok := bma.Keys[key]
	if !ok {
		return nil, stringNotFoundError(key)
	}
	bm := &Bitmap{
		Pitch:  bma.TileWidth,
		Pixels: make([]int, bma.TileWidth*bma.TileHeight),
	}

	i := 0
	rowStart := (coord.Y * bma.TileHeight * bma.Pitch) + (coord.X * bma.TileWidth)
	for y := 0; y < bma.TileHeight; y++ {
		j := rowStart
		for x := 0; x < bma.TileWidth; x++ {
			if j >= len(bma.Bitmap.Pixels) {
				return nil, fmt.Errorf("error getting bitmap with key '%v'; out of pixels (j=%d/len=%d)", key, j, len(bma.Bitmap.Pixels))
			}

			bm.Pixels[i] = bma.Bitmap.Pixels[j]
			i++
			j++
		}
		rowStart = rowStart + bma.Pitch
	}
	return bm, nil
}

// GetTexture returns a texuter for the bitmap at the position associated witht he key
func (bma *BitmapAtlas) GetTexture(vp *ViewPort, key string) (*sdl.Texture, error) {
	bm, err := bma.GetBitmap(key)
	if err != nil {
		return nil, err
	}
	return bm.ToTexture(vp)
}
