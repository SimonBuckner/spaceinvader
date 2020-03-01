package main

import (
	"github.com/SimonBuckner/spaceinvader/gfx"
	"github.com/veandco/go-sdl2/sdl"
)

const testStateName = "Test State"

type testState struct {
	*gfx.Actor
	gs *gameState
}

func newTestState(gs *gameState) *testState {
	a, _ := gs.NewActor(testStateName)
	t := &testState{
		Actor: a,
		gs:    gs,
	}
	a.SetKeyboardEvent(t.keyb)
	return t
}

func (s *testState) keyb(e *sdl.KeyboardEvent) {
	if e.Type == sdl.KEYDOWN {
		switch e.Keysym.Scancode {
		case sdl.SCANCODE_R:
			if (e.Keysym.Mod & sdl.KMOD_SHIFT) == 0 {
				if s.gs.backgroundColor.R < 254 {
					s.gs.backgroundColor.R++
				}
			} else {
				if s.gs.backgroundColor.R > 0 {
					s.gs.backgroundColor.R--
				}
			}
		case sdl.SCANCODE_G:
			if (e.Keysym.Mod & sdl.KMOD_SHIFT) == 0 {
				if s.gs.backgroundColor.G < 254 {
					s.gs.backgroundColor.G++
				}
			} else {
				if s.gs.backgroundColor.G > 0 {
					s.gs.backgroundColor.G--
				}
			}
		case sdl.SCANCODE_B:
			if (e.Keysym.Mod & sdl.KMOD_SHIFT) == 0 {
				if s.gs.backgroundColor.B < 254 {
					s.gs.backgroundColor.B++
				}
			} else {
				if s.gs.backgroundColor.B > 0 {
					s.gs.backgroundColor.B--
				}
			}
		}
	}
}
