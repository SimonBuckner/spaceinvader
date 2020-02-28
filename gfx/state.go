package gfx

import (
	"github.com/veandco/go-sdl2/sdl"
)

// KeyboardHandler registers a function to handle keyboard event
type KeyboardHandler func(e *sdl.KeyboardEvent)

// MouseButtonHandler is called each time a mouse button event is triggered
type MouseButtonHandler func(event *sdl.MouseButtonEvent)

// MouseMotionHandler is called each time a mouse motion event is triggered
type MouseMotionHandler func(event *sdl.MouseMotionEvent)

// MouseWheelHandler is called each time a mouse wheel event is triggered
type MouseWheelHandler func(event *sdl.MouseWheelEvent)

// UpdateHandler is called once each game loop to update game assets before rendering
type UpdateHandler func(vp *ViewPort, ticks uint32)

// StateController defines the methods of the game state controller
type StateController interface {
	Running() bool
	Quit()
	// KeyboardEvent handles keyboard events
	KeyboardEvent(e *sdl.KeyboardEvent)
	MouseButtonEvent(e *sdl.MouseButtonEvent)
	MouseMotionEvent(e *sdl.MouseMotionEvent)
	MouseWheelEvent(e *sdl.MouseWheelEvent)
	UpdateEvent(ticks uint32)
}

// StateControl controls the state of the game
type StateControl struct {
	running bool
	// states  []*State
	// current *State

	keyboardEvent    func(e *sdl.KeyboardEvent)
	mouseButtonEvent func(e *sdl.MouseButtonEvent)
	mouseMotionEvent func(e *sdl.MouseMotionEvent)
	mouseWheelEvent  func(e *sdl.MouseWheelEvent)
	updateEvent      func(ticks uint32)
}

// NewGlobalState factory
func NewGlobalState() *StateControl {
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

// AddState adds a state to the controller
// func (sc *StateControl) AddState(state *State) error {
// 	for _, s := range sc.states {
// 		if s.name == state.name {
// 			return fmt.Errorf("state %v already exists", state.ame)
// 		}
// 	}
// 	sc.states = append(sc.states, state)
// 	return nil
// }

// // Enter makes the named state the current one
// func (sc *StateControl) Enter(stateName StateName) {
// 	for _, s := range sc.states {
// 		if s.name == stateName {
// 			sc.current = s
// 		}
// 	}
// }

// StateName indicates a state in a game
// type StateName string

// // State represents a specific state in a game
// type State struct {
// 	name StateName
// 	// keyboardHandler KeyboardHandler
// 	// kouseButtonHandler KouseButtonHandler
// 	// kouseMotionHandler
// 	// kouseWheelHandler
// 	// kpdateHandler
// }

// // NewState factory
// func NewState(name StateName) *State {
// 	return &State{
// 		name: name,
// 	}
// }
