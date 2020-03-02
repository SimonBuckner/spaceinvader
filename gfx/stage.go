package gfx

import (
	"fmt"
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

// Stage represents a single Window
type Stage struct {
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
	props    []*Prop
	keyboard *Keyboard
}

// NewStage factory
func NewStage(title string, top, left, width, height int, scale float32) (*Stage, error) {
	stage := &Stage{
		title:    title,
		top:      int32(top),
		left:     int32(left),
		width:    int32(width),
		height:   int32(height),
		scale:    scale,
		keyboard: NewKeyboard(),
	}

	{
		window, err := sdl.CreateWindow(stage.title, stage.left, stage.top, stage.width, stage.height, sdl.WINDOW_SHOWN)
		if err != nil {
			return nil, err
		}
		stage.window = window
	}

	{
		renderer, err := sdl.CreateRenderer(stage.window, -1, sdl.RENDERER_ACCELERATED)
		if err != nil {
			return nil, err
		}
		stage.renderer = renderer
	}

	stage.props = make([]*Prop, 0)

	return stage, nil
}

// Run the main event loop for the Stage..
func (stage *Stage) Run(director *Director) {
	var frameStart time.Time
	for {
		frameStart = time.Now()
		stage.keyboard.Refresh()

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

		stage.renderer.Clear()
		director.UpdateEvent(sdl.GetTicks())
		stage.DrawProps()
		stage.renderer.Present()
		sdl.Delay(1)
		stage.elapsedTime = float32(time.Since(frameStart).Seconds())
	}
}

// Destroy cleans u resources
func (stage *Stage) Destroy() {
	if stage.renderer != nil {
		stage.renderer.Destroy()
	}

	if stage.window != nil {
		stage.window.Destroy()
	}
}

// SetBackgroundColor sets the background color
func (stage *Stage) SetBackgroundColor(color sdl.Color) {
	stage.renderer.SetDrawColor(color.R, color.G, color.B, color.A)
}

// HasMouse returnes true if the Stage has the mouse
func (stage *Stage) HasMouse() bool {
	return sdl.GetMouseFocus() == stage.window
}

// HasKeyboard returnes true if the Stage has the keyboard
func (stage *Stage) HasKeyboard() bool {
	return sdl.GetKeyboardFocus() == stage.window
}

// HasBoth returnes true if the Stage has the mouse and the keyboard
func (stage *Stage) HasBoth() bool {
	return sdl.GetKeyboardFocus() == stage.window || sdl.GetMouseFocus() == stage.window
}

// NewSinglePixelTexture returns a new texture comprising a single pixel
func (stage *Stage) NewSinglePixelTexture(r, g, b, a uint8) (*sdl.Texture, error) {
	tex, err := stage.renderer.CreateTexture(sdl.PIXELFORMAT_RGBA8888, sdl.TEXTUREACCESS_STATIC, 1, 1)
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

// AddProp adds an prop
func (stage *Stage) AddProp(prop *Prop) {
	fmt.Printf("prop '%v' added\n", prop.name)
	stage.props = append(stage.props, prop)
}

// RemoveProp removes a screen prop from the render queue
func (stage *Stage) RemoveProp(prop *Prop) {
	for i := 0; i < len(stage.props); i++ {
		if stage.props[i] == prop {
			stage.props = nil
		}
	}
}

// Props returns a slice of Props
func (stage *Stage) Props() []*Prop {
	return stage.props
}

// ClearProps clears all props from the render queue
func (stage *Stage) ClearProps() {
	stage.props = make([]*Prop, 0)
}

// WindowSize returns the width and height of the window
func (stage *Stage) WindowSize() (w, h int32) {
	w, h = stage.window.GetSize()
	return
}

// DrawProps draws the supplied props into the specified Stage
func (stage *Stage) DrawProps() {
	for _, prop := range stage.props {
		if prop != nil && prop.Visible() {
			texture := prop.Texture()
			_, _, w, h, _ := texture.Query()
			x, y, _ := prop.Pos()
			dstRect := sdl.Rect{
				X: x,
				Y: y,
				W: int32(float32(w) * prop.Scale()),
				H: int32(float32(h) * prop.Scale()),
			}
			stage.renderer.Copy(texture, nil, &dstRect)
		}
	}
}

// Renderer gets the stage renderer
func (stage *Stage) Renderer() *sdl.Renderer {
	return stage.renderer
}

// Scale returns the default prop scale
func (stage *Stage) Scale() float32 {
	return stage.scale
}

// SetScale sets the default scale
func (stage *Stage) SetScale(scale float32) {
	stage.scale = scale
}

// ElapsedTime returns the amount of time since the last frame start
func (stage *Stage) ElapsedTime() float32 {
	return stage.elapsedTime
}

// Keyboard returns the keyboard state
func (stage *Stage) Keyboard() *Keyboard {
	return stage.keyboard
}
