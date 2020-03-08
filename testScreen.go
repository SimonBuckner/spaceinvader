package main

import (
	"github.com/SimonBuckner/spaceinvader/gfx"
	"github.com/veandco/go-sdl2/sdl"
)

const testScreenName = "Test Scene"

type testScreen struct {
	scene *gfx.Scene
	game  *game
	p1    *player
}

func newTestScene(game *game) *testScreen {

	s := &testScreen{
		game:  game,
		scene: gfx.NewScene(testScreenName),
		p1:    newPlayer(game),
	}
	s.scene.StartEventHandler = s.onStart
	s.scene.StopEventHandler = s.onStop
	s.scene.UpdateEventHandler = s.onUpdate
	s.scene.KeyboardEventHandler = s.onKeyboard
	return s
}

func (s *testScreen) onStart() {
	s.scene.AddActor(s.p1.Actor)
}

func (s *testScreen) onStop() {
	s.scene.RemoveActor(s.p1.Actor)

}

func (s *testScreen) onKeyboard(e *sdl.KeyboardEvent) {
	if e.Type == sdl.KEYUP {
		switch e.Keysym.Scancode {
		case sdl.SCANCODE_V:
			if s.p1.Visible {
				s.p1.Visible = false
			} else {
				s.p1.reset()
			}
		case sdl.SCANCODE_H:
			if s.p1.Visible && !s.p1.hit {
				s.p1.setHit()
			}
		}
	}
}

func (s *testScreen) onUpdate(ticks uint32) {
	kb := s.scene.KBState()

	if kb.IsKeyDown(sdl.SCANCODE_LEFT) != kb.IsKeyDown(sdl.SCANCODE_RIGHT) {
		if kb.IsKeyDown(sdl.SCANCODE_LEFT) {
			s.p1.moveLeft()
		} else if kb.IsKeyDown(sdl.SCANCODE_RIGHT) {
			s.p1.moveRight()
		}
	}
	s.p1.update(ticks)
}
