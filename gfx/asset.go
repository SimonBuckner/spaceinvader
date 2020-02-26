package gfx

import (
	"fmt"
	"sort"

	"github.com/veandco/go-sdl2/sdl"
)

// Asset represents an on-screen asset
type Asset struct {
	Name     string
	pos      Pos
	w, h     int
	scale    float32
	index    int
	visible  bool
	vp       *ViewPort
	textures []*sdl.Texture
}

// NewAsset factory
func NewAsset(vp *ViewPort) *Asset {
	return &Asset{
		Name:  "",
		pos:   Pos{},
		w:     0,
		h:     0,
		scale: 1.0,
		index: 0,
		vp:    vp,
	}
}

// NewAssetFromTexture returns a new asset from a texture
func NewAssetFromTexture(vp *ViewPort, tex *sdl.Texture) *Asset {
	_, _, w, h, _ := tex.Query()
	a := &Asset{
		Name:     "",
		pos:      Pos{},
		w:        int(w),
		h:        int(h),
		scale:    1.0,
		index:    0,
		vp:       vp,
		textures: make([]*sdl.Texture, 1),
	}
	a.textures[0] = tex
	return a
}

// AssetFromBitmaps converts and array of integer color values into a texture of the specified width.
func AssetFromBitmaps(vp *ViewPort, bitmaps ...*Bitmap) (*Asset, error) {

	asset := &Asset{
		pos:      Pos{},
		scale:    1.0,
		index:    0,
		vp:       vp,
		textures: make([]*sdl.Texture, len(bitmaps)),
	}

	for i, bm := range bitmaps {
		tex, err := bm.ToTexture(vp)
		if err != nil {
			return nil, err
		}
		asset.textures[i] = tex
	}
	return asset, nil
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

// ***************************************************************************
// AssetMap
// ***************************************************************************

// AssetMap represents an atlas of assets accessible by a key
type AssetMap struct {
	assets map[string]*Asset
}

// NewAssetMap factory
func NewAssetMap() *AssetMap {
	return &AssetMap{
		assets: make(map[string]*Asset),
	}
}

// NewAssetMapFromBitMapAtlas returns a newasset map from a btimap atlas
func NewAssetMapFromBitMapAtlas(vp *ViewPort, atlas *BitmapAtlas) (*AssetMap, error) {

	keys := atlas.GetKeys()

	am := NewAssetMap()

	for _, key := range keys {
		tex, err := atlas.GetTileTexture(vp, key)
		if err != nil {
			return nil, fmt.Errorf("error getting tile '%v' from BitmapAtlas; %v", key, err)
		}
		asset := NewAssetFromTexture(vp, tex)
		asset.Name = key
		am.Add(key, asset)
	}
	return am, nil
}

// Add a asset to the asset map
func (am *AssetMap) Add(key string, asset *Asset) {
	am.assets[key] = asset
}

// GetKeys returns a the keys on the AssetMap
func (am *AssetMap) GetKeys() []string {
	keys := make([]string, len(am.assets))
	i := 0
	for k := range am.assets {
		keys[i] = k
		i++
	}
	sort.Strings(keys)
	return keys
}

// GetAssets returns a the keys on the AssetMap
func (am *AssetMap) GetAssets() []*Asset {
	keys := am.GetKeys()
	assets := make([]*Asset, len(keys))
	for i, k := range keys {
		assets[i] = am.assets[k]
	}
	return assets
}

// SurfaceFromBitMap returns an array of surfaces from the bitmap data
// func SurfaceFromBitMap(bitmaps ...*Bitmap) ([]*sdl.Surface, error) {
// 	pf := sdl.PixelFormat{
// 		Format:        sdl.PIXELFORMAT_RGBA8888,
// 		Palette:       nil,
// 		BitsPerPixel:  32,
// 		BytesPerPixel: 4,
// 		Rmask:         rMask,
// 		Gmask:         gMask,
// 		Bmask:         bMask,
// 		Amask:         aMask,
// 	}
// 	surfaces := make([]*sdl.Surface, len(bitmaps))
// 	for i := 0; i < len(bitmaps); i++ {
// 		bm := bitmaps[i]
// 		tc := bm.TransparentColour
// 		ck := sdl.MapRGBA(&pf, tc.R, tc.G, tc.B, tc.A)

// 		w := int32(bm.Pitch)
// 		h := int32(len(bm.Pixels)) / w
// 		if len(bm.Pixels) != int(w*h) {
// 			return nil, fmt.Errorf("unexpected length of bitmap index %d; expected %d, got %d", i, w*h, len(bm.Pixels))
// 		}

// 		surface, _ := sdl.CreateRGBSurface(0, w, h, 32, rMask, gMask, bMask, aMask)
// 		surface.Lock()
// 		pixels := surface.Pixels()
// 		j := 0
// 		for _, pixel := range bm.Pixels {
// 			if j >= len(pixels) {
// 				break
// 			}
// 			rgba := HexColorToRGBA(pixel)
// 			pixels[j] = rgba.R
// 			j++
// 			pixels[j] = rgba.G
// 			j++
// 			pixels[j] = rgba.B
// 			j++
// 			pixels[j] = rgba.A
// 			j++
// 		}
// 		surface.Unlock()
// 		if bm.Transparency {
// 			surface.SetColorKey(true, ck)
// 			surface.SetBlendMode(sdl.BLENDMODE_BLEND)
// 		} else {
// 			surface.SetBlendMode(sdl.BLENDMODE_NONE)
// 		}
// 		surfaces[i] = surface
// 	}
// 	return surfaces, nil
// }
