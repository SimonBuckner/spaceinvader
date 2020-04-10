package main

import (
	"github.com/SimonBuckner/screen2d"
	"github.com/veandco/go-sdl2/sdl"
)

const testScreenName = "Test Scene"

type testScreen struct {
	game        *game
	keyb        *screen2d.KBState
	p1          *player
	p1Shot      *playerShot
	p1AlienRack *alienRack
}

func newTestScene(game *game) *testScreen {

	s := &testScreen{
		game:        game,
		keyb:        game.screen.GetKBState(),
		p1:          newPlayer(game),
		p1Shot:      newPlayerShot(game),
		p1AlienRack: newAlienRack(game),
	}

	return s
}

func (s *testScreen) activate() {
	s.game.screen.ClearFuncs()

	s.p1.reset()
	s.p1Shot.reset()
	s.p1AlienRack.reset(1)
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
		s.p1.moveLeft()
	case sdl.SCANCODE_RIGHT:
		s.p1.moveRight()
	case sdl.SCANCODE_V:
		//
	case sdl.SCANCODE_H:
		//
	}
}

func (s *testScreen) onUpdate(ticks uint32, elapsed float32) {

	if s.keyb.IsKeyDown(sdl.SCANCODE_LEFT) != s.keyb.IsKeyDown(sdl.SCANCODE_RIGHT) {
		if s.keyb.IsKeyDown(sdl.SCANCODE_LEFT) {
			s.p1.moveLeft()
		} else {
			s.p1.moveRight()
		}
	} else {
		s.p1.stopMoving()
	}

	if s.keyb.IsKeyDown(sdl.SCANCODE_Q) {
		s.game.screen.Close()
	}
	if s.keyb.IsKeyDown(sdl.SCANCODE_SPACE) {
		s.p1Shot.fire()
	}

	s.p1.update(ticks, elapsed)
	s.p1Shot.update(ticks, elapsed, s.p1.X)
	s.p1AlienRack.update(ticks, elapsed, s.p1.X)
}

func (s *testScreen) onDraw() {
	s.p1.Draw()
	s.p1Shot.Draw()
	s.p1AlienRack.draw()
}
