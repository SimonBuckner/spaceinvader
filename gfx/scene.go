package gfx

import "github.com/veandco/go-sdl2/sdl"

// Scene represents a specific Scene in a game
type Scene struct {
	name     string
	stage    *Stage
	renderer *sdl.Renderer

	keyboardHandler func(e *sdl.KeyboardEvent)
	// mouseButtonHandler func(e *sdl.MouseButtonEvent)
	// mouseMotionHandler func(e *sdl.MouseMotionEvent)
	// mouseWheelHandler  func(e *sdl.MouseWheelEvent)
	updateHandler func(ticks uint32)

	startHandler func()
	stopHandler  func()
}

// NewScene returns a new game Scene
func NewScene(name string) *Scene {
	return &Scene{
		name: name,
	}
}

// Name ..
func (s *Scene) Name() string {
	return s.name
}

// Renderer ..
func (s *Scene) Renderer() *sdl.Renderer {
	return s.renderer
}

// SetStage set the stage the scene will unfold on
func (s *Scene) SetStage(stage *Stage) {
	s.stage = stage
}

// SetKeyboardEvent sets the keyboard handler
func (s *Scene) SetKeyboardEvent(handler func(e *sdl.KeyboardEvent)) {
	s.keyboardHandler = handler
}

// KeyboardEvent triggers the keyboard event handler for the director and the current scene
func (s *Scene) KeyboardEvent(e *sdl.KeyboardEvent) {
	if s.keyboardHandler != nil {
		s.keyboardHandler(e)
	}
}

// SetUpdateEvent sets the update handler
func (s *Scene) SetUpdateEvent(handler func(ticks uint32)) {
	s.updateHandler = handler
}

// UpdateEvent triggers the update event handler for the director and the current scene
func (s *Scene) UpdateEvent(ticks uint32) {
	if s.updateHandler != nil {
		s.updateHandler(ticks)
	}
}

// SetStartEvent sets the start handler
func (s *Scene) SetStartEvent(handler func()) {
	s.startHandler = handler
}

// StartEvent triggers the start event handler for the director and the current scene
func (s *Scene) StartEvent() {
	if s.startHandler != nil {
		s.startHandler()
	}
}

// SetStopEvent sets the start handler
func (s *Scene) SetStopEvent(handler func()) {
	s.startHandler = handler
}

// StopEvent triggers the start event handler for the director and the current scene
func (s *Scene) StopEvent() {
	if s.startHandler != nil {
		s.startHandler()
	}
}
