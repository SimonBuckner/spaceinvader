package gfx

// Actor represents a specific Actor in a game
type Actor struct {
	Name  string
	Scene *Scene
	// Props   []*Prop
	Visible bool
	Scale   float32
	Pos     Vec3
	Speed   Vec3

	// StartEventHandler  func()
	// StopEventHandler   func()
	// UpdateEventHandler func(uint32)
}

// NewActor factory
func NewActor(name string) *Actor {
	return &Actor{
		Name: name,
		// Props:   make([]*Prop, 0),
		Visible: false,
		Scale:   1,
		Pos:     Vec3{},
		Speed:   Vec3{},
	}
}

// SetScale ..
func (a *Actor) SetScale(scale float32) {
	a.Scale = scale
}

// Start the actor
// func (a *Actor) Start(scene *Scene) {
// 	a.Scene = scene
// 	if a.StartEventHandler != nil {
// 		a.StartEventHandler()
// 	}
// }

// Stop the actor
// func (a *Actor) Stop() {
// 	if a.StopEventHandler != nil {
// 		a.StopEventHandler()
// 	}
// 	a.Scene = nil
// }

// Update the actor
// func (a *Actor) Update(ticks uint32) {
// 	if a.UpdateEventHandler != nil {
// 		a.UpdateEventHandler(ticks)
// 	}
// }

// Draw draws all props the actor is holding
// func (a *Actor) Draw() {
// 	if !a.Visible {
// 		return
// 	}

// 	for _, prop := range a.Props {

// 	}
// }

// AddProp adds a prop to an actor
// func (a *Actor) AddProp(prop *Prop) {
// 	if prop == nil {
// 		panic("AddProp is nil")
// 	}
// 	prop.Scale = a.Scale
// 	a.Props = append(a.Props, prop)
// }

// RemoveProp removes a prop from an actor
// func (a *Actor) RemoveProp(prop *Prop) {
// 	for i, p := range a.Props {
// 		if p == prop {
// 			a.Props = append(a.Props[:i], a.Props[i+1:]...)
// 			return
// 		}
// 	}
// }

// ClearProps removes all props
// func (a *Actor) ClearProps() {
// 	for _, p := range a.Props {
// 		p.Texture.Destroy()
// 	}
// 	a.Props = make([]*Prop, 0)
// }

// Destory all props
// func (a *Actor) Destory() {
// 	a.ClearProps()
// }

// SetScene ..
// func (a *Actor) SetScene(scene *Scene) {
// 	a.Scene = scene
// }

// ClearScene ..
// func (a *Actor) ClearScene() {
// 	a.Scene = nil
// }

// // Update ..
// func (a *Actor) Update(ticks uint32) {

// }

// // Draw ..
// func (a *Actor) Draw() {

// }

// // Unload ..
// func (a *Actor) Unload() {

// }
