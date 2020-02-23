package gfx

import (
	"github.com/veandco/go-sdl2/sdl"
)

// ViewPort represents a single Window
type ViewPort struct {
	title    string
	top      int32
	left     int32
	width    int32
	height   int32
	window   *sdl.Window
	renderer *sdl.Renderer
	KeyboardHandler
	MouseButtonHandler
	MouseMotionHandler
	MouseWheelHandler
	UpdateHandler
	Assets []*Asset
}

// NewViewPort factory
func NewViewPort(title string, top, left, width, height int) (*ViewPort, error) {
	vp := &ViewPort{
		title:  title,
		top:    int32(top),
		left:   int32(left),
		width:  int32(width),
		height: int32(height),
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

	vp.Assets = make([]*Asset, 0)

	return vp, nil
}

// Run the main event loop for the ViewPort..
func (vp *ViewPort) Run(state State) {
	for {
		if !state.IsRunning() {
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
				if vp.KeyboardHandler != nil {
					vp.KeyboardHandler(e)
				}
			case *sdl.MouseButtonEvent:
				if vp.MouseButtonHandler != nil {
					vp.MouseButtonHandler(e)
				}
			case *sdl.MouseMotionEvent:
				if vp.MouseMotionHandler != nil {
					vp.MouseMotionHandler(e)
				}
			case *sdl.MouseWheelEvent:
				if vp.MouseWheelHandler != nil {
					vp.MouseWheelHandler(e)
				}
			}
		}

		vp.renderer.Clear()
		if vp.UpdateHandler != nil {
			vp.UpdateHandler(vp)
		}

		vp.DrawAssets()

		vp.renderer.Present()
		sdl.Delay(1)
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
	tex, err := vp.renderer.CreateTexture(sdl.PIXELFORMAT_ABGR8888, sdl.TEXTUREACCESS_STATIC, 1, 1)
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

// DrawAsset ..
func (vp *ViewPort) DrawAssets() {
	for _, asset := range vp.Assets {
		tex := asset.Texture()
		_, _, w, h, _ := tex.Query()
		dstRect := sdl.Rect{
			X: asset.pos.X,
			Y: asset.pos.Y,
			W: w,
			H: h,
		}
		vp.renderer.Copy(tex, nil, &dstRect)
	}
}

// AddAsset adds an asset
func (vp *ViewPort) AddAsset(asset *Asset) {
	vp.Assets = append(vp.Assets, asset)
}
