package gfx

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

// StateController defines the methods of the game state controller
type StateController interface {
	Running() bool
	Quit()
	KeyboardEvent(e *sdl.KeyboardEvent)
	MouseButtonEvent(e *sdl.MouseButtonEvent)
	MouseMotionEvent(e *sdl.MouseMotionEvent)
	MouseWheelEvent(e *sdl.MouseWheelEvent)
	UpdateEvent(ticks uint32)
}

// StateControl controls the state of the game
type StateControl struct {
	running bool
	states  []*State
	current *State

	keyboardEvent    func(e *sdl.KeyboardEvent)
	mouseButtonEvent func(e *sdl.MouseButtonEvent)
	mouseMotionEvent func(e *sdl.MouseMotionEvent)
	mouseWheelEvent  func(e *sdl.MouseWheelEvent)
	updateEvent      func(ticks uint32)
}

// StateName indicates a state in a game
type StateName string

// State represents a specific state in a game
type State struct {
	name             StateName
	sc               *StateControl
	keyboardEvent    func(e *sdl.KeyboardEvent)
	mouseButtonEvent func(e *sdl.MouseButtonEvent)
	mouseMotionEvent func(e *sdl.MouseMotionEvent)
	mouseWheelEvent  func(e *sdl.MouseWheelEvent)
	updateEvent      func(ticks uint32)
}

// NewStateControl factory
func NewStateControl() *StateControl {
	return &StateControl{
		running: true,
		// states:  make([]*State, 0),
	}
}

// Running returns true is running
func (sc *StateControl) Running() bool {
	return sc.running
}

// Quit signals the viewport to quit
func (sc *StateControl) Quit() {
	sc.running = false
}

// SetKeyboardEvent sets the keyboard event handler
func (sc *StateControl) SetKeyboardEvent(handler func(e *sdl.KeyboardEvent)) {
	sc.keyboardEvent = handler
}

// KeyboardEvent handles keyboard events
func (sc *StateControl) KeyboardEvent(e *sdl.KeyboardEvent) {
	if sc.keyboardEvent != nil {
		sc.keyboardEvent(e)
	}
}

// SetMouseButtonEvent sets the MouseButton event handler
func (sc *StateControl) SetMouseButtonEvent(handler func(e *sdl.MouseButtonEvent)) {
	sc.mouseButtonEvent = handler
}

// MouseButtonEvent handles MouseButton events
func (sc *StateControl) MouseButtonEvent(e *sdl.MouseButtonEvent) {
	if sc.mouseButtonEvent != nil {
		sc.mouseButtonEvent(e)
	}
}

// SetMouseMotionEvent sets the MouseMotion event handler
func (sc *StateControl) SetMouseMotionEvent(handler func(e *sdl.MouseMotionEvent)) {
	sc.mouseMotionEvent = handler
}

// MouseMotionEvent handles MouseMotion events
func (sc *StateControl) MouseMotionEvent(e *sdl.MouseMotionEvent) {
	if sc.mouseMotionEvent != nil {
		sc.mouseMotionEvent(e)
	}
}

// SetMouseWheelEvent sets the MouseWheel event handler
func (sc *StateControl) SetMouseWheelEvent(handler func(e *sdl.MouseWheelEvent)) {
	sc.mouseWheelEvent = handler
}

// MouseWheelEvent handles MouseWheel events
func (sc *StateControl) MouseWheelEvent(e *sdl.MouseWheelEvent) {
	if sc.MouseWheelEvent != nil {
		sc.MouseWheelEvent(e)
	}
}

// SetUpdateEvent ..
func (sc *StateControl) SetUpdateEvent(handler func(ticks uint32)) {
	sc.updateEvent = handler
}

// UpdateEvent ..
func (sc *StateControl) UpdateEvent(ticks uint32) {
	if sc.updateEvent != nil {
		sc.updateEvent(ticks)
	}
}

// NewState creates a new state and adds a state to the controller
func (sc *StateControl) NewState(name StateName) (*State, error) {
	for _, s := range sc.states {
		if s.name == name {
			return nil, fmt.Errorf("state %v already exists", string(name))
		}
	}
	state := &State{
		name: name,
		sc:   sc,
	}
	sc.states = append(sc.states, state)
	return state, nil
}

// EnterState makes the named state the current one
func (sc *StateControl) EnterState(name StateName) error {
	for _, s := range sc.states {
		if s.name == name {
			sc.current = s
			return nil
		}
	}
	return fmt.Errorf("Unable to finn the state %v", string(name))
}
