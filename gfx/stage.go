package gfx

import (
	"fmt"
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

// Stage represents a single Window
type Stage struct {
	title    string
	window   *sdl.Window
	renderer *sdl.Renderer
	keyboard *Keyboard
	scenes   []*Scene
	current  *Scene
	props    []*Prop

	top         int32
	left        int32
	width       int32
	height      int32
	scale       float32
	frameStart  time.Time
	elapsedTime float32

	closing bool
	// startHandler  func()
	// stopHandler   func()
	updateHandler func(ticks uint32)

	keyboardHandler    func(e *sdl.KeyboardEvent)
	mouseButtonHandler func(e *sdl.MouseButtonEvent)
	mouseMotionHandler func(e *sdl.MouseMotionEvent)
	mouseWheelHandler  func(e *sdl.MouseWheelEvent)
}

// NewStage fsceney
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

// AttachScene attaches a scene to the stage
func (s *Stage) AttachScene(scene *Scene) error {
	for _, v := range s.scenes {
		if v.Name() == scene.Name() {
			return fmt.Errorf("the scene '%v' already exists in the stage '%v'", scene.Name(), v.Name())
		}
	}
	s.scenes = append(s.scenes, scene)
	scene.SetStage(s)
	return nil
}

// Run the main event loop for the Stage..
func (s *Stage) Run() {
	var frameStart time.Time
	for {
		frameStart = time.Now()
		s.keyboard.Refresh()

		if s.closing {
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
				fmt.Println("stage.keyb")
				s.KeyboardEvent(e)
			case *sdl.MouseButtonEvent:
				// director.MouseButtonEvent(e)
			case *sdl.MouseMotionEvent:
				// director.MouseMotionEvent(e)
			case *sdl.MouseWheelEvent:
				// director.MouseWheelEvent(e)
			}
		}

		s.renderer.Clear()
		s.UpdateEvent(sdl.GetTicks())
		s.DrawProps()
		s.renderer.Present()
		sdl.Delay(1)
		s.elapsedTime = float32(time.Since(frameStart).Seconds())
	}
}

// Destroy cleans u resources
func (s *Stage) Destroy() {
	if s.renderer != nil {
		s.renderer.Destroy()
	}

	if s.window != nil {
		s.window.Destroy()
	}
}

// SetBackgroundColor sets the background color
func (s *Stage) SetBackgroundColor(color sdl.Color) {
	s.renderer.SetDrawColor(color.R, color.G, color.B, color.A)
}

// HasMouse returnes true if the Stage has the mouse
func (s *Stage) HasMouse() bool {
	return sdl.GetMouseFocus() == s.window
}

// HasKeyboard returnes true if the Stage has the keyboard
func (s *Stage) HasKeyboard() bool {
	return sdl.GetKeyboardFocus() == s.window
}

// HasBoth returnes true if the Stage has the mouse and the keyboard
func (s *Stage) HasBoth() bool {
	return sdl.GetKeyboardFocus() == s.window || sdl.GetMouseFocus() == s.window
}

// NewSinglePixelTexture returns a new texture comprising a single pixel
func (s *Stage) NewSinglePixelTexture(r, g, b, a uint8) (*sdl.Texture, error) {
	tex, err := s.renderer.CreateTexture(sdl.PIXELFORMAT_RGBA8888, sdl.TEXTUREACCESS_STATIC, 1, 1)
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
func (s *Stage) AddProp(prop *Prop) {
	fmt.Printf("prop '%v' added\n", prop.name)
	s.props = append(s.props, prop)
}

// RemoveProp removes a screen prop from the render queue
func (s *Stage) RemoveProp(prop *Prop) {
	for i := 0; i < len(s.props); i++ {
		if s.props[i] == prop {
			s.props = nil
		}
	}
}

// Props returns a slice of Props
func (s *Stage) Props() []*Prop {
	return s.props
}

// ClearProps clears all props from the render queue
func (s *Stage) ClearProps() {
	s.props = make([]*Prop, 0)
}

// WindowSize returns the width and height of the window
func (s *Stage) WindowSize() (w, h int32) {
	w, h = s.window.GetSize()
	return
}

// DrawProps draws the supplied props into the specified Stage
func (s *Stage) DrawProps() {
	for _, prop := range s.props {
		if prop != nil && prop.Visible() {
			texture := prop.Texture()
			_, _, w, h, _ := texture.Query()
			x, y, _ := prop.Int32()
			dstRect := sdl.Rect{
				X: x,
				Y: y,
				W: int32(float32(w) * prop.Scale()),
				H: int32(float32(h) * prop.Scale()),
			}
			s.renderer.Copy(texture, nil, &dstRect)
		}
	}
}

// Renderer gets the stage renderer
func (s *Stage) Renderer() *sdl.Renderer {
	return s.renderer
}

// Scale returns the default prop scale
func (s *Stage) Scale() float32 {
	return s.scale
}

// SetScale sets the default scale
func (s *Stage) SetScale(scale float32) {
	s.scale = scale
}

// ElapsedTime returns the amount of time since the last frame start
func (s *Stage) ElapsedTime() float32 {
	return s.elapsedTime
}

// Keyboard returns the keyboard state
func (s *Stage) Keyboard() *Keyboard {
	return s.keyboard
}

// SetKeyboardEvent sets the keyboard handler
func (s *Stage) SetKeyboardEvent(handler func(e *sdl.KeyboardEvent)) {
	s.keyboardHandler = handler
}

// KeyboardEvent triggers the keyboard event handler for the director and the current scene
func (s *Stage) KeyboardEvent(e *sdl.KeyboardEvent) {
	if !s.closing {
		if s.keyboardHandler != nil {
			s.keyboardHandler(e)
		}
		if s.current != nil {
			s.current.KeyboardEvent(e)
		}
	}
}

// SetUpdateEvent sets the update handler
func (s *Stage) SetUpdateEvent(handler func(ticks uint32)) {
	s.updateHandler = handler
}

// UpdateEvent triggers the update event handler for the director and the current scene
func (s *Stage) UpdateEvent(ticks uint32) {
	if !s.closing {
		if s.updateHandler != nil {
			s.updateHandler(ticks)
		}
		if s.current != nil {
			s.current.UpdateEvent(ticks)
		}
	}
}

// Close closes the stage
func (s *Stage) Close() {
	s.closing = true
}
