package main

import (
	"github.com/SimonBuckner/spaceinvader/gfx"
)

const title = "SCORE-1   HI-SCORE   SCORE-2"

type score struct {
	*gfx.Actor
	banner  *banner
	p1Score int
	hiScore int
	p2Score int

	titleText   *banner
	p1ScoreText *banner
	hiScoreText *banner
	p2ScoreText *banner
}

func newScore(game *game) *score {
	s := &score{
		Actor:       gfx.NewActor("score"),
		p1Score:     0,
		hiScore:     0,
		p2Score:     0,
		titleText:   newBanner(game, "CEHIORS12- ", title, len(title)),
		p1ScoreText: newBanner(game, "0123456789", "0000", 4),
		hiScoreText: newBanner(game, "0123456789", "0000", 4),
		p2ScoreText: newBanner(game, "0123456789", "0000", 4),
	}

	s.titleText.Pos.SetInt32(7*2, 0, 0)
	s.p1ScoreText.Pos.SetInt32(7*4, 12, 0)
	s.hiScoreText.Pos.SetInt32(7*14, 12, 0)
	s.p2ScoreText.Pos.SetInt32(7*25, 12, 0)

	s.titleText.Visible = true
	s.p1ScoreText.Visible = true
	s.hiScoreText.Visible = true
	s.p2ScoreText.Visible = true

	return s
}

func (s *score) Start(scene *gfx.Scene) {
	s.Scene = scene
	s.Scale = scene.Scale()

	s.titleText.Start(scene)
	s.p1ScoreText.Start(scene)
	s.hiScoreText.Start(scene)
	s.p2ScoreText.Start(scene)

	s.Scene.AddActor(s.titleText)
	s.Scene.AddActor(s.p1ScoreText)
	s.Scene.AddActor(s.hiScoreText)
	s.Scene.AddActor(s.p2ScoreText)
}

func (s *score) Draw() {
	s.titleText.Draw()
	s.p1ScoreText.Draw()
	s.hiScoreText.Draw()
	s.p2ScoreText.Draw()
}

func (s *score) Update(ticks uint32) {
	// s.p1ScoreText.value = fmt.Sprintf("%04d", s.p1Score)
	// s.hiScoreText.value = fmt.Sprintf("%04d", s.hiScore)
	// s.p2ScoreText.value = fmt.Sprintf("%04d", s.p2Score)
}
