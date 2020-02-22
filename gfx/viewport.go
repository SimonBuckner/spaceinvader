package gfx

import "github.com/veandco/go-sdl2/sdl"

// ViewPort represents a single Window
type ViewPort struct {
	title    string
	top      int32
	left     int32
	width    int32
	height   int32
	window   *sdl.Window
	renderer *sdl.Renderer
	mouse    *Mouse
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

	vp.mouse = NewMouse()
	return vp, nil
}

// Run the main event loop for the ViewPort..
func (vp *ViewPort) Run() {
	i := uint8(0)
	for {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch e := event.(type) {
			case *sdl.QuitEvent:
				return
			case *sdl.WindowEvent:
				if e.Event == sdl.WINDOWEVENT_CLOSE {
					return
				}
			}
		}
		vp.mouse.Refresh()
		vp.renderer.Clear()

		// Mouse events
		if vp.mouse.LeftDown() {
			if i < 254 {
				i++
			}
		} else if vp.mouse.RightDown() {
			if i > 0 {
				i--
			}
		}
		vp.SetBackgroundColor(i, 0, 0, 0)
		vp.renderer.Present()
		sdl.Delay(1)
	}
}

// SetBackgroundColor sets the background color
func (vp *ViewPort) SetBackgroundColor(r, g, b, a uint8) {
	vp.renderer.SetDrawColor(r, g, b, a)
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
