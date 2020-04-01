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
	s.game.screen.SetKeyUpFunc(s.onKeyUp)
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
		s.p1.moveLeft()
	case sdl.SCANCODE_RIGHT:
		s.p1.moveRight()
	case sdl.SCANCODE_V:
		//
	case sdl.SCANCODE_H:
		//
	}
}

func (s *testScreen) onKeyUp(e *sdl.KeyboardEvent) {
	switch e.Keysym.Scancode {
	case sdl.SCANCODE_LEFT, sdl.SCANCODE_RIGHT:
		s.p1.stopMoving()
	}
}

func (s *testScreen) onUpdate(ticks uint32, elapsed float32) {
	s.p1.update(ticks, elapsed)
}

func (s *testScreen) onDraw() {
	s.p1.Draw()
}
