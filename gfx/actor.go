package gfx

// Actor represents a specific Actor in a game
type Actor struct {
	Pos
	Speed         Pos
	name          string
	scene         *Scene
	startHandler  func()
	stopHandler   func()
	updateHandler func(ticks uint32)
}

// NewActor returns a new game Actor
func NewActor(name string) *Actor {
	return &Actor{
		name: name,
	}
}

// Name ..
func (a *Actor) Name() string {
	return a.name
}

// Scene returns the scene for the scene
func (a *Actor) Scene() *Scene {
	return a.scene
}

// SetScene sets the scene for the scene
func (a *Actor) SetScene(scene *Scene) {
	a.scene = scene
}

// SetStartEvent sets the start handler
func (a *Actor) SetStartEvent(handler func()) {
	a.startHandler = handler
}

// StartEvent triggers the start event handler for the scene and the current scene
func (a *Actor) StartEvent() {
	if a.startHandler != nil {
		a.startHandler()
	}
}

// SetStopEvent sets the start handler
func (a *Actor) SetStopEvent(handler func()) {
	a.startHandler = handler
}

// StopEvent triggers the start event handler for the scene and the current scene
func (a *Actor) StopEvent() {
	if a.startHandler != nil {
		a.startHandler()
	}
}

// SetUpdateEvent sets the update handler
func (a *Actor) SetUpdateEvent(handler func(ticks uint32)) {
	a.updateHandler = handler
}

// UpdateEvent triggers the update event handler for the scene and the current scene
func (a *Actor) UpdateEvent(ticks uint32) {
	if a.updateHandler != nil {
		a.updateHandler(ticks)
	}
}
