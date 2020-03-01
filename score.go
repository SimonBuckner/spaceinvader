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
	assets []*gfx.Asset
	digits []*sdl.Texture

	gs     *gameState
	points int
}

func newScore(gs *gameState, pos scorePos) (*score, error) {
	s := &score{
		gs:     gs,
		points: 0,
		assets: make([]*gfx.Asset, 4),
		digits: make([]*sdl.Texture, 10),
	}

	keys := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	for i, k := range keys {
		tex, err := alphabetAtlas.GetTexture(gs.vp, k)
		if err != nil {
			return nil, fmt.Errorf("error getting '%v' texture for score; %v", k, err)
		}
		s.digits[i] = tex
	}

	x := int32(pos)
	y := int32(2)
	for i, asset := range s.assets {
		name := "score " + strconv.Itoa(int(i))
		asset = gfx.NewAssetFromTexture(gs.vp, name, s.digits[0])

		newX, newY := gs.convertXY(x, y)
		asset.SetPos(newX, newY, 0)
		s.assets[i] = asset
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
	// 	x, y := s.gs.convertXY(int32(s.x), int32(s.y))
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
	// 	s.Asset.SetTexture(s.explode1Tex)
	// } else {
	// 	s.Asset.SetTexture(s.explode2Tex)
	// }
}

// Reset the score to a
func (s *score) Reset() {
	s.points = 0
	for _, v := range s.assets {

		v.Show()
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
// 		s.x = s.x - float32(s.speed*s.gs.vs.ElapsedTime())
// 	}
// }

// // MoveRight moves the score right
// func (p *score) MoveRight() {
// 	if s.lives == 0 || s.exploding == true {
// 		return
// 	}
// 	if s.x < originalWidth {
// 		s.x = s.x + float32(s.speed*s.gs.vs.ElapsedTime())
// 	}
// }
