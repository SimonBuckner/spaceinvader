package gfx

import (
	"github.com/veandco/go-sdl2/sdl"
)

// Actor represents a specific Actor in a game
type Actor struct {
	Name    string
	Scene   *Scene
	Props   []*Prop
	Visible bool
	Scale   float32
	Pos     Vec3
	Speed   Vec3

	StartEventHandler func()
	StopEventHandler  func()
}

// NewActor factory
func NewActor(name string) *Actor {
	return &Actor{
		Name:    name,
		Props:   make([]*Prop, 0),
		Visible: false,
		Scale:   1,
		Pos:     Vec3{},
		Speed:   Vec3{},
	}
}

// Start the actor
func (a *Actor) Start(scene *Scene) {
	a.Scene = scene
	if a.StartEventHandler != nil {
		a.StartEventHandler()
	}
}

// Stop the actor
func (a *Actor) Stop() {
	if a.StopEventHandler != nil {
		a.StopEventHandler()
	}
	a.Scene = nil
}

// Draw draws all props the actor is holding
func (a *Actor) Draw() {
	if !a.Visible {
		return
	}

	for _, prop := range a.Props {
		if prop.Texture == nil {
			continue
		}

		tex := prop.Texture
		_, _, w, h, _ := tex.Query()
		x, y, _ := prop.Pos.Int32()

		dstRect := sdl.Rect{
			X: x,
			Y: y,
			W: int32(float32(w) * prop.Scale),
			H: int32(float32(h) * prop.Scale),
		}
		a.Scene.Renderer().Copy(tex, nil, &dstRect)
	}
}

// AddProp adds a prop to an actor
func (a *Actor) AddProp(prop *Prop) {
	if prop == nil {
		panic("AddProp is nil")
	}
	prop.Scale = a.Scale
	a.Props = append(a.Props, prop)
}

// RemoveProp removes a prop from an actor
func (a *Actor) RemoveProp(prop *Prop) {
	for i, p := range a.Props {
		if p == prop {
			a.Props = append(a.Props[:i], a.Props[i+1:]...)
			return
		}
	}
}

// ClearProps removes all props
func (a *Actor) ClearProps() {
	a.Props = make([]*Prop, 0)
}
