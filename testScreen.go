package main

import (
	"github.com/SimonBuckner/spaceinvader/gfx"
	"github.com/veandco/go-sdl2/sdl"
)

const testScreenName = "Test Scene"

type testScreen struct {
	*gfx.Scene
	game *game
	p1   *player
	t    *text
}

func newTestScene(game *game) *testScreen {

	s := &testScreen{
		Scene: gfx.NewScene(testScreenName),
		game:  game,
		p1:    newPlayer(game),
		t:     newText(game, "ABCDEFGHIJKLMNOPQRSTUVYXWZ01234567899 "),
	}
	s.StartEventHandler = s.onStart
	s.StopEventHandler = s.onStop
	s.UpdateEventHandler = s.onUpdate
	s.KeyboardEventHandler = s.onKeyboard
	return s
}

func (s *testScreen) onStart() {
	s.AddActor(s.p1.Actor)
	s.t.value = "SCORE P1"
	s.AddActor(s.t.Actor)
}

func (s *testScreen) onStop() {
	s.RemoveActor(s.p1.Actor)
	s.RemoveActor(s.t.Actor)
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
	kb := s.KBState()

	if kb.IsKeyDown(sdl.SCANCODE_LEFT) != kb.IsKeyDown(sdl.SCANCODE_RIGHT) {
		if kb.IsKeyDown(sdl.SCANCODE_LEFT) {
			s.p1.moveLeft()
		} else if kb.IsKeyDown(sdl.SCANCODE_RIGHT) {
			s.p1.moveRight()
		}
	}
	s.p1.update(ticks)
	s.t.update(ticks)
}
