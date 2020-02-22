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
	mouse    *Mouse
	keyboard *Keyboard
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
	vp.keyboard = NewKeyboard()
	return vp, nil
}

// Run the main event loop for the ViewPort..
func (vp *ViewPort) Run() {
	r := uint8(0)
	g := uint8(0)
	b := uint8(0)

	for {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch e := event.(type) {
			case *sdl.QuitEvent:
				return
			case *sdl.WindowEvent:
				if e.Event == sdl.WINDOWEVENT_CLOSE {
					return
				}
			case *sdl.KeyboardEvent:
				if e.Type == sdl.KEYDOWN {
					if e.Keysym.Scancode == sdl.SCANCODE_R {
						if (e.Keysym.Mod & sdl.KMOD_SHIFT) == 0 {
							if r < 254 {
								r++
							}
						} else {
							if r > 0 {
								r--
							}
						}
					}
					if e.Keysym.Scancode == sdl.SCANCODE_G {
						if (e.Keysym.Mod & sdl.KMOD_SHIFT) == 0 {
							if g < 254 {
								g++
							}
						} else {
							if g > 0 {
								g--
							}
						}
					}
					if e.Keysym.Scancode == sdl.SCANCODE_B {
						if (e.Keysym.Mod & sdl.KMOD_SHIFT) == 0 {
							if b < 254 {
								b++
							}
						} else {
							if b > 0 {
								b--
							}
						}
					}
				}
			}
		}
		vp.mouse.Refresh()
		vp.keyboard.Refresh()

		vp.renderer.Clear()

		vp.SetBackgroundColor(r, g, b, 0)
		vp.renderer.Present()
		sdl.Delay(1)
	}
}

// SetBackgroundColor sets the background color
func (vp *ViewPort) SetBackgroundColor(r, g, b, a uint8) {
	vp.renderer.SetDrawColor(r, g, b, a)
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
