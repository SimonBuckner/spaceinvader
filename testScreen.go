package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

const testScreenName = "Test Scene"

type testScreen struct {
	game *game
	p1   *player
	// score     *score
	// alienRack *alienRack
}

func newTestScene(game *game) *testScreen {

	s := &testScreen{
		game: game,
		p1:   newPlayer(game),
		// alienRack: newAlienRack(game),
	}
	return s
}

func (s *testScreen) activate() {
	s.game.screen.ClearFuncs()
	s.game.screen.SetKeyDownFunc(s.onKeyDown)
	s.game.screen.SetUpdateFunc(s.onUpdate)
	s.game.screen.SetDrawFunc(s.onDraw)
}

func (s *testScreen) onKeyDown(e *sdl.KeyboardEvent) {
	switch e.Keysym.Scancode {
	case sdl.SCANCODE_ESCAPE:
		s.game.activate()
	case sdl.SCANCODE_Q:
		s.game.screen.Close()
	case sdl.SCANCODE_LEFT:
		//
	case sdl.SCANCODE_RIGHT:
		//
	case sdl.SCANCODE_V:
		//
	case sdl.SCANCODE_H:
		//
	}
}

func (s *testScreen) onUpdate(ticks uint32, elapsed float32) {
	s.p1.update(ticks, elapsed)
	// kb := s.KBState()

	// if kb.IsKeyDown(sdl.SCANCODE_LEFT) != kb.IsKeyDown(sdl.SCANCODE_RIGHT) {
	// 	if kb.IsKeyDown(sdl.SCANCODE_LEFT) {
	// 		s.p1.moveLeft()
	// 	} else if kb.IsKeyDown(sdl.SCANCODE_RIGHT) {
	// 		s.p1.moveRight()
	// 	}
	// }
	// if kb.OnKeyDown(sdl.SCANCODE_SPACE) {
	// 	s.p1.fireShot()
	// }
	// if s.p1.shotState == shotInFlight {
	// 	if s.alienRack.checkForHit(s.p1) {
	// 		s.p1.shotState = shotHit
	// 	}
	// }
}

func (s *testScreen) onDraw() {
	s.p1.Draw()
}
