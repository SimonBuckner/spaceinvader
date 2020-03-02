package gfx

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

// Director holds the main Actor of the game. All other Actors are called from this Actor.
type Director struct {
	actors  []*Actor
	current *Actor
	closing bool

	keyboardHandler    func(e *sdl.KeyboardEvent)
	mouseButtonHandler func(e *sdl.MouseButtonEvent)
	mouseMotionHandler func(e *sdl.MouseMotionEvent)
	mouseWheelHandler  func(e *sdl.MouseWheelEvent)
	updateHandler      func(ticks uint32)
}

// NewDirector factory
func NewDirector() *Director {
	return &Director{
		actors:  make([]*Actor, 0),
		closing: false,
	}
}

// NewActor returns a new game Actor linked to a director
func (d *Director) NewActor(name string) (*Actor, error) {
	for _, actor := range d.actors {
		if actor.name == name {
			return nil, fmt.Errorf("there is already a Actor named %v", string(name))
		}
	}
	actor := &Actor{
		director: d,
		name:     name,
	}
	d.actors = append(d.actors, actor)
	return actor, nil
}

// IsLoading indicates no Actor has been entered
func (d *Director) IsLoading() bool {
	return !d.closing && d.current == nil
}

// IsRunning indicates a Actor has been entered
func (d *Director) IsRunning() bool {
	return !d.closing && d.current != nil
}

// IsClosing indicates the game is closing
func (d *Director) IsClosing() bool {
	return d.closing
}

// Close the director
func (d *Director) Close() {
	d.closing = true
}

// StartActor enters the named Actor
func (d *Director) StartActor(name string) error {
	for _, s := range d.actors {
		if s.name == name {
			d.current = s
			s.StartEvent()
			return nil
		}
	}
	return fmt.Errorf("Actor '%v' not found", name)
}

// SetKeyboardEvent sets the keyboard handler
func (d *Director) SetKeyboardEvent(handler func(e *sdl.KeyboardEvent)) {
	d.keyboardHandler = handler
}

// KeyboardEvent triggers the keyboard event handler for the director and the current actor
func (d *Director) KeyboardEvent(e *sdl.KeyboardEvent) {
	if d.keyboardHandler != nil {
		d.keyboardHandler(e)
	}
	if d.IsRunning() {
		if d.current != nil {
			d.current.KeyboardEvent(e)
		}
	}
}
