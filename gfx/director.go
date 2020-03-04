package gfx

// // Director holds the main Scene of the game. All other Scenes are called from this Scene.
// type Director struct {

// }

// // NewDirector fsceney
// func NewDirector() *Director {
// 	return &Director{
// 		scenes:  make([]*Scene, 0),
// 		closing: false,
// 	}
// }

// // IsLoading indicates no Scene has been entered
// func (d *Director) IsLoading() bool {
// 	return !d.closing && d.current == nil
// }

// // IsRunning indicates a Scene has been entered
// func (d *Director) IsRunning() bool {
// 	return !d.closing && d.current != nil
// }

// // IsClosing indicates the game is closing
// func (d *Director) IsClosing() bool {
// 	return d.closing
// }

// Close the director
// func (d *Director) Close() {
// 	d.closing = true
// }

// // SetStartEvent sets the start handler
// func (s *Actor) SetStartEvent(handler func()) {
// 	s.startHandler = handler
// }

// // StartEvent triggers the start event handler for the director and the current scene
// func (s *Actor) StartEvent() {
// 	if s.startHandler != nil {
// 		s.startHandler()
// 	}
// }

// // SetStopEvent sets the start handler
// func (s *Actor) SetStopEvent(handler func()) {
// 	s.startHandler = handler
// }

// // StopEvent triggers the start event handler for the director and the current scene
// func (s *Actor) StopEvent() {
// 	if s.startHandler != nil {
// 		s.startHandler()
// 	}
// }

// // SetUpdateEvent sets the update handler
// func (d *Director) SetUpdateEvent(handler func(ticks uint32)) {
// 	d.updateHandler = handler
// }

// // UpdateEvent triggers the update event handler for the director and the current scene
// func (d *Director) UpdateEvent(ticks uint32) {
// 	if d.updateHandler != nil {
// 		d.updateHandler(ticks)
// 	}
// 	if d.IsRunning() {
// 		if d.current != nil {
// 			d.current.UpdateEvent(ticks)
// 		}
// 	}
// }

// // StartScene enters the named Scene
// func (d *Director) StartScene(name string) error {
// 	for _, s := range d.scenes {
// 		if s.Name() == name {
// 			d.current = s
// 			s.StartEvent()
// 			return nil
// 		}
// 	}
// 	return fmt.Errorf("Scene '%v' not found", name)
// }

// // SetKeyboardEvent sets the keyboard handler
// func (d *Director) SetKeyboardEvent(handler func(e *sdl.KeyboardEvent)) {
// 	d.keyboardHandler = handler
// }

// // KeyboardEvent triggers the keyboard event handler for the director and the current scene
// func (d *Director) KeyboardEvent(e *sdl.KeyboardEvent) {
// 	if d.keyboardHandler != nil {
// 		d.keyboardHandler(e)
// 	}
// 	if d.IsRunning() {
// 		if d.current != nil {
// 			d.current.KeyboardEvent(e)
// 		}
// 	}
// }

// // AttachScene attaches a scene to the director
// func (d *Director) AttachScene(scene *Scene) error {
// 	for _, s := range d.scenes {
// 		if scene.Name() == s.Name() {
// 			return fmt.Errorf("there is already a Scene named %v", scene.Name())
// 		}
// 	}
// 	d.scenes = append(d.scenes, scene)
// 	return nil
// }
