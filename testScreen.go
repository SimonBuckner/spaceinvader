package main

import (
	"github.com/SimonBuckner/spaceinvader/gfx"
	"github.com/veandco/go-sdl2/sdl"
)

const testScreenName = "Test Scene"

type testScreen struct {
	*gfx.Scene
	game      *game
	p1        *player
	score     *score
	alienRack *alienRack
}

func newTestScene(game *game) *testScreen {

	s := &testScreen{
		Scene:     gfx.NewScene(testScreenName),
		game:      game,
		p1:        newPlayer(game),
		score:     newScore(game),
		alienRack: newAlienRack(game),
	}

	s.StartEventHandler = s.onStart
	s.StopEventHandler = s.onStop
	s.UpdateEventHandler = s.onUpdate
	s.KeyboardEventHandler = s.onKeyboard

	return s
}

func (s *testScreen) onStart() {
	s.AddActor(s.p1)
	s.AddActor(s.score)
	s.AddActor(s.alienRack)
}

func (s *testScreen) onStop() {
	s.ClearActors()
}

func (s *testScreen) onKeyboard(e *sdl.KeyboardEvent) {
	if e.Type == sdl.KEYUP {
		switch e.Keysym.Scancode {
		case sdl.SCANCODE_V:
			if s.p1.Visible {
				s.p1.Visible = false
			} else {
				s.p1.resetShip()
			}
		case sdl.SCANCODE_H:
			if s.p1.Visible && s.p1.shipState == shipAlive {
				s.p1.setShipHit()
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
	if kb.OnKeyDown(sdl.SCANCODE_SPACE) {
		s.p1.fireShot()
	}
	if s.p1.shotState == shotInFlight {
		if s.alienRack.checkForHit(s.p1) {
			s.p1.shotState = shotHit
		}
	}
}
