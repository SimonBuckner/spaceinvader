package gfx

import "github.com/veandco/go-sdl2/sdl"

// Actor represents a specific Actor in a game
type Actor struct {
	name     string
	director *Director

	keyboardHandler    func(e *sdl.KeyboardEvent)
	mouseButtonHandler func(e *sdl.MouseButtonEvent)
	mouseMotionHandler func(e *sdl.MouseMotionEvent)
	mouseWheelHandler  func(e *sdl.MouseWheelEvent)
	updateHandler      func(ticks uint32)

	startHandler func()
	stopHandler  func()
}

// SetKeyboardEvent sets the keyboard handler
func (a *Actor) SetKeyboardEvent(handler func(e *sdl.KeyboardEvent)) {
	a.keyboardHandler = handler
}

// KeyboardEvent triggers the keyboard event handler for the director and the current actor
func (a *Actor) KeyboardEvent(e *sdl.KeyboardEvent) {
	if a.keyboardHandler != nil {
		a.keyboardHandler(e)
	}
}

// SetUpdateEvent sets the update handler
func (d *Director) SetUpdateEvent(handler func(ticks uint32)) {
	d.updateHandler = handler
}

// UpdateEvent triggers the update event handler for the director and the current actor
func (d *Director) UpdateEvent(ticks uint32) {
	if d.updateHandler != nil {
		d.updateHandler(ticks)
	}
	if d.IsRunning() {
		if d.current != nil {
			d.current.UpdateEvent(ticks)
		}
	}
}

// SetUpdateEvent sets the update handler
func (a *Actor) SetUpdateEvent(handler func(ticks uint32)) {
	a.updateHandler = handler
}

// UpdateEvent triggers the update event handler for the director and the current actor
func (a *Actor) UpdateEvent(ticks uint32) {
	if a.updateHandler != nil {
		a.updateHandler(ticks)
	}
}

// SetStartEvent sets the start handler
func (a *Actor) SetStartEvent(handler func()) {
	a.startHandler = handler
}

// StartEvent triggers the start event handler for the director and the current actor
func (a *Actor) StartEvent() {
	if a.startHandler != nil {
		a.startHandler()
	}
}
