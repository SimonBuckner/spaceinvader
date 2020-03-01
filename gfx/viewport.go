package gfx

import (
	"fmt"
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

// ViewPort represents a single Window
type ViewPort struct {
	title       string
	top         int32
	left        int32
	width       int32
	height      int32
	scale       float32
	frameStart  time.Time
	elapsedTime float32

	window   *sdl.Window
	renderer *sdl.Renderer
	assets   []*Asset
	keyboard *Keyboard
}

// NewViewPort factory
func NewViewPort(title string, top, left, width, height int, scale float32) (*ViewPort, error) {
	vp := &ViewPort{
		title:    title,
		top:      int32(top),
		left:     int32(left),
		width:    int32(width),
		height:   int32(height),
		scale:    scale,
		keyboard: NewKeyboard(),
	}

	{
		window, err := sdl.CreateWindow(vp.title, vp.left, vp.top, vp.width, vp.height, sdl.WINDOW_SHOWN)
		if err != nil {
			return nil, err
		}
		vp.window = window
	}

	{
		renderer, err := sdl.CreateRenderer(vp.window, -1, sdl.RENDERER_ACCELERATED)
		if err != nil {
			return nil, err
		}
		vp.renderer = renderer
	}

	vp.assets = make([]*Asset, 0)

	return vp, nil
}

// Run the main event loop for the ViewPort..
func (vp *ViewPort) Run(director *Director) {
	var frameStart time.Time
	for {
		frameStart = time.Now()
		vp.keyboard.Refresh()

		if director.IsClosing() {
			fmt.Println("Quit event")
			return
		}
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch e := event.(type) {
			case *sdl.QuitEvent:
				return
			case *sdl.WindowEvent:
				if e.Event == sdl.WINDOWEVENT_CLOSE {
					return
				}
			case *sdl.KeyboardEvent:
				director.KeyboardEvent(e)
			case *sdl.MouseButtonEvent:
				// director.MouseButtonEvent(e)
			case *sdl.MouseMotionEvent:
				// director.MouseMotionEvent(e)
			case *sdl.MouseWheelEvent:
				// director.MouseWheelEvent(e)
			}
		}

		vp.renderer.Clear()
		director.UpdateEvent(sdl.GetTicks())
		vp.DrawAssets()
		vp.renderer.Present()
		sdl.Delay(1)
		vp.elapsedTime = float32(time.Since(frameStart).Seconds())
	}
}

// Destroy cleans u resources
func (vp *ViewPort) Destroy() {
	if vp.renderer != nil {
		vp.renderer.Destroy()
	}

	if vp.window != nil {
		vp.window.Destroy()
	}
}

// SetBackgroundColor sets the background color
func (vp *ViewPort) SetBackgroundColor(color sdl.Color) {
	vp.renderer.SetDrawColor(color.R, color.G, color.B, color.A)
}

// HasMouse returnes true if the ViewPort has the mouse
func (vp *ViewPort) HasMouse() bool {
	return sdl.GetMouseFocus() == vp.window
}

// HasKeyboard returnes true if the ViewPort has the keyboard
func (vp *ViewPort) HasKeyboard() bool {
	return sdl.GetKeyboardFocus() == vp.window
}

// HasBoth returnes true if the ViewPort has the mouse and the keyboard
func (vp *ViewPort) HasBoth() bool {
	return sdl.GetKeyboardFocus() == vp.window || sdl.GetMouseFocus() == vp.window
}

// NewSinglePixelTexture returns a new texture comprising a single pixel
func (vp *ViewPort) NewSinglePixelTexture(r, g, b, a uint8) (*sdl.Texture, error) {
	tex, err := vp.renderer.CreateTexture(sdl.PIXELFORMAT_RGBA8888, sdl.TEXTUREACCESS_STATIC, 1, 1)
	tex.SetBlendMode(sdl.BLENDMODE_ADD)
	if err != nil {
		return nil, err
	}
	pixels := make([]byte, 4)
	pixels[0] = r
	pixels[1] = g
	pixels[2] = b
	pixels[3] = a

	tex.Update(nil, pixels, 4)
	return tex, nil
}

// AddAsset adds an asset
func (vp *ViewPort) AddAsset(asset *Asset) {
	fmt.Printf("asset '%v' added\n", asset.Name)
	vp.assets = append(vp.assets, asset)
}

// RemoveAsset removes a screen asset from the render queue
func (vp *ViewPort) RemoveAsset(asset *Asset) {
	for i := 0; i < len(vp.assets); i++ {
		if vp.assets[i] == asset {
			vp.assets = nil
		}
	}
}

// Assets returns a slice of Assets
func (vp *ViewPort) Assets() []*Asset {
	return vp.assets
}

// ClearAssets clears all assets from the render queue
func (vp *ViewPort) ClearAssets() {
	vp.assets = make([]*Asset, 0)
}

// WindowSize returns the width and height of the window
func (vp *ViewPort) WindowSize() (w, h int32) {
	w, h = vp.window.GetSize()
	return
}

// DrawAssets draws the supplied assets into the specified ViewPort
func (vp *ViewPort) DrawAssets() {
	for _, asset := range vp.assets {
		if asset != nil && asset.IsVisible() {
			texture := asset.Texture()
			_, _, w, h, _ := texture.Query()
			x, y, _ := asset.Pos()
			dstRect := sdl.Rect{
				X: x,
				Y: y,
				W: int32(float32(w) * asset.Scale()),
				H: int32(float32(h) * asset.Scale()),
			}
			vp.renderer.Copy(texture, nil, &dstRect)
		}
	}
}

// Renderer gets the viewport renderer
func (vp *ViewPort) Renderer() *sdl.Renderer {
	return vp.renderer
}

// Scale returns the default asset scale
func (vp *ViewPort) Scale() float32 {
	return vp.scale
}

// SetScale sets the default scale
func (vp *ViewPort) SetScale(scale float32) {
	vp.scale = scale
}

// ElapsedTime returns the amount of time since the last frame start
func (vp *ViewPort) ElapsedTime() float32 {
	return vp.elapsedTime
}

// Keyboard returns the keyboard state
func (vp *ViewPort) Keyboard() *Keyboard {
	return vp.keyboard
}
