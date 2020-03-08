package gfx

import (
	"fmt"
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

// Stage represents a single Window
type Stage struct {
	Title    string
	Window   *sdl.Window
	Renderer *sdl.Renderer
	KBState  *KBState

	Top    int32
	Left   int32
	Width  int32
	Height int32
	Scale  float32

	FrameStart  time.Time
	ElapsedTime float32
	Stopping    bool

	KeyboardEventHandler func(e *sdl.KeyboardEvent)

	Scene *Scene
}

// NewStage fsceney
func NewStage(title string, top, left, width, height int, scale float32) (*Stage, error) {
	stage := &Stage{
		Title:    title,
		Top:      int32(top),
		Left:     int32(left),
		Width:    int32(width),
		Height:   int32(height),
		Scale:    scale,
		KBState:  NewKBState(),
		Stopping: false,
	}

	{
		window, err := sdl.CreateWindow(title, stage.Left, stage.Top, stage.Width, stage.Height, sdl.WINDOW_SHOWN)
		if err != nil {
			return nil, err
		}
		stage.Window = window
	}

	{
		renderer, err := sdl.CreateRenderer(stage.Window, -1, sdl.RENDERER_ACCELERATED)
		if err != nil {
			return nil, err
		}
		stage.Renderer = renderer
	}

	return stage, nil
}

// Destroy cleans u resources
func (s *Stage) Destroy() {
	if s.Renderer != nil {
		s.Renderer.Destroy()
	}

	if s.Window != nil {
		s.Window.Destroy()
	}
}

// Start the main event loop for the Stage..
func (s *Stage) Start() {
	var frameStart time.Time
	for {
		if s.Stopping {
			return
		}

		frameStart = time.Now()
		s.KBState.Refresh()

		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch e := event.(type) {
			case *sdl.QuitEvent:
				return
			case *sdl.WindowEvent:
				if e.Event == sdl.WINDOWEVENT_CLOSE {
					return
				}
			case *sdl.KeyboardEvent:
				s.FireKeyboardEvent(e)
			}
		}

		s.Renderer.Clear()
		if s.Scene != nil {
			s.Scene.FireUpdateEvent(sdl.GetTicks())
			s.Scene.Draw()
		}
		s.Renderer.Present()
		sdl.Delay(1)
		s.ElapsedTime = float32(time.Since(frameStart).Seconds())
	}
}

// Stop the stage
func (s *Stage) Stop() {
	s.StopScene()
	s.Stopping = true
}

// StartScene starts the supplied scene
func (s *Stage) StartScene(scene *Scene) {
	scene.Start(s)
	s.Scene = scene
}

// StopScene stops the running scene
func (s *Stage) StopScene() {
	if s.Scene != nil {
		s.Scene.Stop()
		s.Scene = nil
	}
}

// FireKeyboardEvent triggers the keyboard event handler for the director and the current scene
func (s *Stage) FireKeyboardEvent(e *sdl.KeyboardEvent) {
	if !s.Stopping {
		if s.KeyboardEventHandler != nil {
			s.KeyboardEventHandler(e)
		}
		if s.Scene != nil {
			s.Scene.FireKeyboardEvent(e)
		}
	}
}

// WindowSize returns the width and height of the widow
func (s *Stage) WindowSize() (w, h int32) {
	return s.Window.GetSize()
}

// SetBackgroundColor sets the background color
func (s *Stage) SetBackgroundColor(color sdl.Color) {
	s.Renderer.SetDrawColor(color.R, color.G, color.B, color.A)
}

// NewSinglePixelTexture returns a new texture comprising a single pixel
func (s *Stage) NewSinglePixelTexture(r, g, b, a uint8) (*sdl.Texture, error) {
	tex, err := s.Renderer.CreateTexture(sdl.PIXELFORMAT_RGBA8888, sdl.TEXTUREACCESS_STATIC, 1, 1)
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

// DumpActors dumps out basic details about loaded props
func (s *Stage) DumpActors() {
	fmt.Println("index  name                     x     y visible")
	fmt.Println("=====  ====================  ====  ==== =======")
	for i, a := range s.Scene.Actors {
		x, y, _ := a.Pos.Int32()
		fmt.Printf(" %3d   %-20v  %4d  %4d %v\n", i, a.Name, x, y, a.Visible)
	}
}
