package gfx

import (
	"github.com/veandco/go-sdl2/sdl"
)

// Scene represents a specific Scene in a game
type Scene struct {
	Name   string
	Stage  *Stage
	Actors []*Actor

	KeyboardEventHandler    func(e *sdl.KeyboardEvent)
	MouseButtonEventHandler func(e *sdl.MouseButtonEvent)
	MouseMotionEventHandler func(e *sdl.MouseMotionEvent)
	MouseWheelEventHandler  func(e *sdl.MouseWheelEvent)

	StartEventHandler  func()
	UpdateEventHandler func(ticks uint32)
	StopEventHandler   func()
}

// NewScene returns a new game Scene
func NewScene(name string) *Scene {
	return &Scene{
		Name:   name,
		Actors: make([]*Actor, 0),
	}
}

// KBState returns the keyboard
func (s *Scene) KBState() *KBState {
	return s.Stage.KBState
}

// Renderer returns the current renderer
func (s *Scene) Renderer() *sdl.Renderer {
	return s.Stage.Renderer
}

// ElapsedTime returns the legnth of the last from in MS
func (s *Scene) ElapsedTime() float32 {
	return s.Stage.ElapsedTime
}

// Scale returns the default scale factor
func (s *Scene) Scale() float32 {
	return s.Stage.Scale
}

// Start starts the scene on the specified stage
func (s *Scene) Start(stage *Stage) {
	s.Stage = stage
	if s.StartEventHandler != nil {
		s.StartEventHandler()
	}
	for _, a := range s.Actors {
		a.Start(s)
	}
}

// Stop stops the running scene
func (s *Scene) Stop() {
	for _, a := range s.Actors {
		a.Stop()
	}
	if s.StopEventHandler != nil {
		s.StopEventHandler()
	}
	s.Stage = nil
}

// AddActor adds a actor to the scene
func (s *Scene) AddActor(actor *Actor) {
	actor.Scene = s
	actor.Scale = s.Scale()
	s.Actors = append(s.Actors, actor)
}

// RemoveActor removes a actor from the scene
func (s *Scene) RemoveActor(actor *Actor) {
	for i, p := range s.Actors {
		if p == actor {
			s.Actors = append(s.Actors[:i], s.Actors[i+1:]...)
			actor.Scene = nil
			return
		}
	}
}

// ClearActors removes all actors
func (s *Scene) ClearActors() {
	for _, a := range s.Actors {
		a.Scene = nil
	}
	s.Actors = make([]*Actor, 0)
}

// Draw draws the supplied props into the specified Stage
func (s *Scene) Draw() {
	for _, a := range s.Actors {
		a.Draw()
	}
}

// FireKeyboardEvent triggers the keyboard event handler for the director and the current scene
func (s *Scene) FireKeyboardEvent(e *sdl.KeyboardEvent) {
	if s.KeyboardEventHandler != nil {
		s.KeyboardEventHandler(e)
	}
}

// FireUpdateEvent triggers the update event handler for the director and the current scene
func (s *Scene) FireUpdateEvent(ticks uint32) {
	if s.UpdateEventHandler != nil {
		s.UpdateEventHandler(ticks)
	}
	for _, a := range s.Actors {
		a.Update(ticks)
	}
}