package main

import (
	"github.com/SimonBuckner/spaceinvader/gfx"
)

type score struct {
	*gfx.Actor
	text    *text
	p1Score int
	hiScore int
	p2Score int

	titleText *text
	// p1ScoreText *text
	// hiScoreText *text
	// p2ScoreText *text
}

func newScore(game *game) *score {
	s := &score{
		Actor: gfx.NewActor("score"),
		// p1Score:     0,
		// hiScore:     0,
		// p2Score:     0,
		titleText: newText(game, "CEHIORS12- ", "SCORE-1   HI-SCORE   SCORE-2"),
		// p1ScoreText: newText(game, "0123456789", "0000"),
		// hiScoreText: newText(game, "0123456789", "0000"),
		// p2ScoreText: newText(game, "0123456789", "0000"),
	}

	s.titleText.Pos.SetInt32(7*2, 0, 0)
	// s.p1ScoreText.Pos.SetInt32(7*4, 12, 0)
	// s.hiScoreText.Pos.SetInt32(7*14, 12, 0)
	// s.p2ScoreText.Pos.SetInt32(7*25, 12, 0)

	s.titleText.Visible = true
	// s.p1ScoreText.Visible = true
	// s.hiScoreText.Visible = true
	// s.p2ScoreText.Visible = true

	s.StartEventHandler = s.onStart
	s.StopEventHandler = s.onStop
	s.UpdateEventHandler = s.onUpdate

	return s
}

func (s *score) onStart() {
	s.Scene.AddActor(s.titleText.Actor)
	// s.Scene.AddActor(s.p1ScoreText.Actor)
	// s.Scene.AddActor(s.hiScoreText.Actor)
	// s.Scene.AddActor(s.p2ScoreText.Actor)

	s.titleText.onStart()
	// s.p1ScoreText.onStart()
	// s.hiScoreText.onStart()
	// s.p2ScoreText.onStart()
}

func (s *score) onStop() {
	s.Scene.ClearActors()
}

func (s *score) onUpdate(ticks uint32) {
	// s.p1ScoreText.value = fmt.Sprintf("%04d", s.p1Score)
	// s.hiScoreText.value = fmt.Sprintf("%04d", s.hiScore)
	// s.p2ScoreText.value = fmt.Sprintf("%04d", s.p2Score)
}
