package main

import (
	"fmt"
	"strconv"

	"github.com/SimonBuckner/spaceinvader/gfx"
	"github.com/veandco/go-sdl2/sdl"
)

const fontWidth = 9

type scorePos int

const (
	scoreP1Pos   scorePos = 10
	scoreHighPos scorePos = (originalWidth - (4 * fontWidth)) / 2
	scoreP2Pos   scorePos = originalWidth - 10 - (4 * fontWidth)
)

type score struct {
	props  []*gfx.Prop
	digits []*sdl.Texture

	game   *game
	points int
}

func newScore(game *game, pos scorePos) (*score, error) {
	s := &score{
		points: 0,
		props:  make([]*gfx.Prop, 4),
		digits: make([]*sdl.Texture, 10),
		game:   game,
	}

	keys := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	for i, k := range keys {
		tex, err := alphabetAtlas.GetTexture(game.stage, k)
		if err != nil {
			return nil, fmt.Errorf("error getting '%v' texture for score; %v", k, err)
		}
		s.digits[i] = tex
	}

	x := int32(pos)
	y := int32(2)
	for i, prop := range s.props {
		name := "score " + strconv.Itoa(int(i))
		prop = gfx.NewProp(game.stage, name, s.digits[0])

		newX, newY := game.convertXY(x, y)
		prop.SetInt32(newX, newY, 0)
		s.props[i] = prop
		x += fontWidth
	}

	return s, nil
}

func (s *score) update(ticks uint32) {

	// if s.lives == 0 {
	// 	s.Hide()
	// 	return
	// }
	// if s.exploding == false {
	// 	x, y := s.stage.convertXY(int32(s.x), int32(s.y))
	// 	s.SetPos(int32(x), int32(y), 0)
	// 	return
	// }

	// if s.exploding && s.explodeCount >= 10 {
	// 	s.lives--
	// 	s.Hide()
	// 	return
	// }

	// if ticks-s.ticks > (16 * 4) {
	// 	s.ticks = ticks
	// 	s.explodeCount++
	// }

	// if s.explodeCount%2 == 0 {
	// 	s.Prop.SetTexture(s.explode1Tex)
	// } else {
	// 	s.Prop.SetTexture(s.explode2Tex)
	// }
}

// Reset the score to a
func (s *score) Reset() {
	s.points = 0
	for _, v := range s.props {

		v.SetVisible(true)
	}
}

// // Hit indicates the score has been hit
// func (p *score) Hit() {
// 	s.exploding = true
// 	s.ticks = sdl.GetTicks()
// }

// // MoveLeft moves the score left
// func (p *score) MoveLeft() {
// 	// paddle.y += paddle.speed * pct * elapsedTime //
// 	if s.lives == 0 || s.exploding == true {
// 		return
// 	}
// 	if s.x > 0 {
// 		s.x = s.x - float32(s.speed*s.stage.vs.ElapsedTime())
// 	}
// }

// // MoveRight moves the score right
// func (p *score) MoveRight() {
// 	if s.lives == 0 || s.exploding == true {
// 		return
// 	}
// 	if s.x < originalWidth {
// 		s.x = s.x + float32(s.speed*s.stage.vs.ElapsedTime())
// 	}
// }
