package gfx

import "github.com/veandco/go-sdl2/sdl"

const (
	keyUp   uint8 = 0
	keyDown uint8 = 1
)

// Keyboard represents the state of the keyboard
type Keyboard struct {
	state  []uint8
	state1 []uint8
}

// NewKeyboard fsceney
func NewKeyboard() *Keyboard {
	kb := &Keyboard{}

	kb.state = sdl.GetKeyboardState()
	kb.state1 = make([]uint8, len(kb.state))

	return kb
}

// Refresh the keyboard state
func (kb *Keyboard) Refresh() {
	for i, v := range kb.state {
		kb.state1[i] = v
	}
	kb.state = sdl.GetKeyboardState()
}

// OnKeyDown returns true when the key state changes from up to down
func (kb *Keyboard) OnKeyDown(key uint8) bool {
	return kb.state[key] == keyDown && kb.state1[key] == keyUp
}

// OnKeyUp returns true when the key state changes from up to down
func (kb *Keyboard) OnKeyUp(key uint8) bool {
	return kb.state[key] == keyUp && kb.state1[key] == keyDown
}

// IsKeyDown returns true if the key state is down
func (kb *Keyboard) IsKeyDown(key uint8) bool {
	return kb.state[key] == keyDown
}

// IsKeyUp returns true if the key state is up
func (kb *Keyboard) IsKeyUp(key uint8) bool {
	return kb.state[key] == keyUp
}
