package main

import (
	"fmt"

	"github.com/SimonBuckner/spaceinvader/gfx"
	"github.com/veandco/go-sdl2/sdl"
)

const (
	playerStartX = 1
	playerStartY = 250
	playerSpeed  = 60
)

type player struct {
	*gfx.Prop

	aliveTex    *sdl.Texture
	explode1Tex *sdl.Texture
	explode2Tex *sdl.Texture

	gs           *gameState
	score        int
	lives        int
	exploding    bool
	explodeCount int
	ticks        uint32
	x            float32
	y            float32
	speed        float32
}

func newPlayer(gs *gameState) (*player, error) {
	p := &player{
		gs:           gs,
		score:        0,
		lives:        3,
		exploding:    false,
		explodeCount: 0,
		ticks:        0,
		x:            playerStartX,
		y:            playerStartY,
		speed:        playerSpeed,
	}

	err := p.loadTextures(gs.stage, playerSprite, plrBlowupSprite0, plrBlowupSprite1)
	if err != nil {
		return nil, err
	}
	p.Prop = gfx.NewProp(gs.stage, "player", p.aliveTex)
	x, y := gs.convertXY(int32(p.x), int32(p.y))
	p.SetPos(x, y, 0)

	return p, nil
}

func (p *player) loadTextures(stage *gfx.Stage, alive, explode1, explode2 *gfx.Bitmap) error {

	var err error
	p.aliveTex, err = alive.ToTexture(stage)
	if err != nil {
		return fmt.Errorf("unable to load live1 bitmap")
	}
	p.explode1Tex, err = explode1.ToTexture(stage)
	if err != nil {
		return fmt.Errorf("unable to load live2 bitmap")
	}
	p.explode2Tex, err = explode2.ToTexture(stage)
	if err != nil {
		return fmt.Errorf("unable to load hit bitmap")
	}
	return nil
}

func (p *player) update(ticks uint32) {

	if p.lives == 0 {
		p.SetVisible(false)
		return
	}
	if p.exploding == false {
		x, y := p.gs.convertXY(int32(p.x), int32(p.y))
		p.SetPos(int32(x), int32(y), 0)
		return
	}

	if p.exploding && p.explodeCount >= 10 {
		p.lives--
		p.SetVisible(false)
		return
	}

	if ticks-p.ticks > (16 * 4) {
		p.ticks = ticks
		p.explodeCount++
	}

	if p.explodeCount%2 == 0 {
		p.Prop.SetTexture(p.explode1Tex)
	} else {
		p.Prop.SetTexture(p.explode2Tex)
	}
}

// Reset the player to a
func (p *player) Reset() {
	p.score = 0
	p.lives = 3
	p.exploding = false
	p.explodeCount = 0
	p.ticks = 0
	p.x = playerStartX

	p.SetTexture(p.aliveTex)
	p.SetVisible(true)
}

// Hit indicates the player has been hit
func (p *player) Hit() {
	p.exploding = true
	p.ticks = sdl.GetTicks()
}

// MoveLeft moves the player left
func (p *player) MoveLeft() {
	// paddle.y += paddle.speed * pct * elapsedTime //
	if p.lives == 0 || p.exploding == true {
		return
	}
	if p.x > 0 {
		p.x = p.x - float32(p.speed*p.gs.stage.ElapsedTime())
	}
}

// MoveRight moves the player right
func (p *player) MoveRight() {
	if p.lives == 0 || p.exploding == true {
		return
	}
	if p.x < originalWidth {
		p.x = p.x + float32(p.speed*p.gs.stage.ElapsedTime())
	}
}
