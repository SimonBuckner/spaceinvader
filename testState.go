package main

import (
	"fmt"

	"github.com/SimonBuckner/spaceinvader/gfx"
	"github.com/veandco/go-sdl2/sdl"
)

const testStateName = "Test State"

type testState struct {
	*gfx.Actor
	gs *gameState
	p1 *player
}

func newTestState(gs *gameState) (*testState, error) {
	a, _ := gs.NewActor(testStateName)
	t := &testState{
		Actor: a,
		gs:    gs,
	}
	a.SetKeyboardEvent(t.keyb)
	a.SetStartEvent(t.start)
	a.SetUpdateEvent(t.update)

	p1, err := newPlayer(gs)
	if err != nil {
		return nil, fmt.Errorf("eorror getting new testState; %v", err)
	}
	t.p1 = p1
	return t, nil
}

func (s *testState) start() {
	s.p1.Reset()
	s.gs.vp.AddAsset(s.p1.Asset)
}

func (s *testState) keyb(e *sdl.KeyboardEvent) {
	if e.Type == sdl.KEYUP {
		switch e.Keysym.Scancode {
		case sdl.SCANCODE_V:
			if s.p1.IsVisible() {
				s.p1.Hide()
			} else {
				s.p1.Reset()
				s.p1.Show()
			}
		}
	}
	// 	case sdl.SCANCODE_H:
	// 		if s.p1.IsVisible() && s.p1.exploding == false {
	// 			s.p1.Hit()
	// 		}
	// 	case sdl.SCANCODE_LEFT:
	// 		s.p1.x = float32(int32(s.p1.x))
	// 	case sdl.SCANCODE_RIGHT:
	// 		s.p1.x = float32(int32(s.p1.x))
	// 	}
	// }
	// if e.Type == sdl.KEYDOWN {
	// 	switch e.Keysym.Scancode {
	// 	case sdl.SCANCODE_LEFT:
	// 		s.p1.MoveLeft()
	// 	case sdl.SCANCODE_RIGHT:
	// 		s.p1.MoveRight()
	// 	}
	// }
}

func (s *testState) update(ticks uint32) {
	kb := s.gs.vp.Keyboard()
	if kb.IsKeyDown(sdl.SCANCODE_LEFT) {
		s.p1.MoveLeft()
	} else if kb.IsKeyDown(sdl.SCANCODE_RIGHT) {
		s.p1.MoveRight()
	}

	s.p1.update(ticks)
}
