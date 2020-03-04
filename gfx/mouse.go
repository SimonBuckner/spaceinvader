package gfx

// // Mouse represents the state of the mouse
// type Mouse struct {
// 	left  bool
// 	right bool
// 	pos   Pos

// 	left1  bool
// 	right1 bool
// 	pos1   Pos
// }

// // NewMouse fsceney
// func NewMouse() *Mouse {
// 	m := &Mouse{}
// 	m.Refresh()
// 	m.Refresh()
// 	return m
// }

// // Refresh the mouse state..
// func (m *Mouse) Refresh() {
// 	m.left1 = m.left
// 	m.right1 = m.right
// 	m.pos1 = m.pos

// 	x, y, state := sdl.GetMouseState()
// 	m.left = !((state & sdl.ButtonLMask()) == 0)
// 	m.right = !((state & sdl.ButtonRMask()) == 0)
// 	m.pos.X = x
// 	m.pos.Y = y
// }

// // LeftUp returns true if left button is not held down
// func (m *Mouse) LeftUp() bool {
// 	return m.left == false
// }

// // LeftDown returns true if left button is held down
// func (m *Mouse) LeftDown() bool {
// 	return m.left
// }

// // RightUp returns true if right button is not held down
// func (m *Mouse) RightUp() bool {
// 	return m.right == false
// }

// // RightDown returns true if right button not held down
// func (m *Mouse) RightDown() bool {
// 	return m.right
// }
