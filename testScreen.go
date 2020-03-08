package main

import (
	"github.com/SimonBuckner/spaceinvader/gfx"
	"github.com/veandco/go-sdl2/sdl"
)

const testScreenName = "Test Scene"

type testScreen struct {
	*gfx.Scene
	game    *game
	p1      *player
	titles  *text
	p1Score *text
	hiScore *text
	p2Score *text
}

func newTestScene(game *game) *testScreen {

	s := &testScreen{
		Scene:   gfx.NewScene(testScreenName),
		game:    game,
		p1:      newPlayer(game),
		titles:  newText(game, "SCORE-1 HI-SCORE SCORE-2"),
		p1Score: newText(game, "1234"),
		hiScore: newText(game, "5678"),
		p2Score: newText(game, "9012"),
	}
	s.titles.Pos.SetInt32(7*1, 0, 0)
	s.p1Score.Pos.SetInt32(7*3, 12, 0)
	s.hiScore.Pos.SetInt32(7*11, 12, 0)
	s.p2Score.Pos.SetInt32(7*20, 12, 0)

	s.StartEventHandler = s.onStart
	s.StopEventHandler = s.onStop
	s.UpdateEventHandler = s.onUpdate
	s.KeyboardEventHandler = s.onKeyboard
	return s
}

func (s *testScreen) onStart() {
	s.AddActor(s.p1.Actor)
	s.AddActor(s.titles.Actor)
	s.AddActor(s.p1Score.Actor)
	s.AddActor(s.hiScore.Actor)
	s.AddActor(s.p2Score.Actor)
}

func (s *testScreen) onStop() {
	s.RemoveActor(s.p1.Actor)
	s.RemoveActor(s.titles.Actor)
	s.RemoveActor(s.p1Score.Actor)
	s.RemoveActor(s.hiScore.Actor)
	s.RemoveActor(s.p2Score.Actor)
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
	s.titles.update(ticks)
	s.p1Score.update(ticks)
	s.hiScore.update(ticks)
	s.p2Score.update(ticks)
}
