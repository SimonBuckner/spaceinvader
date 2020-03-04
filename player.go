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
	*gfx.Actor
	prop        *gfx.Prop
	game        *game
	aliveTex    *sdl.Texture
	explode1Tex *sdl.Texture
	explode2Tex *sdl.Texture

	score        int
	lives        int
	exploding    bool
	explodeCount int
	ticks        uint32
}

func newPlayer(game *game) (*player, error) {
	p := &player{
		Actor:        gfx.NewActor("player"),
		game:         game,
		score:        0,
		lives:        3,
		exploding:    false,
		explodeCount: 0,
	}

	err := p.loadTextures(game.stage, playerSprite, plrBlowupSprite0, plrBlowupSprite1)
	if err != nil {
		return nil, err
	}
	p.prop = gfx.NewProp(game.stage, "player", p.aliveTex)
	x, y := game.convertXY(playerStartX, playerStartY)
	p.Pos.SetInt32(x, y, 0)
	p.prop.SetInt32(x, y, 0)
	p.Speed.Set(60, 0, 0)

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
		p.prop.SetVisible(false)
		return
	}
	if p.exploding == false {
		p.prop.Set(p.X, p.Y, p.Z)
		return
	}

	if p.exploding && p.explodeCount >= 10 {
		p.lives--
		p.prop.SetVisible(false)
		return
	}

	if ticks-p.ticks > (16 * 4) {
		p.ticks = ticks
		p.explodeCount++
	}

	if p.explodeCount%2 == 0 {
		p.prop.SetTexture(p.explode1Tex)
	} else {
		p.prop.SetTexture(p.explode2Tex)
	}
}

// Reset the player to a
func (p *player) Reset() {
	p.score = 0
	p.lives = 3
	p.exploding = false
	p.explodeCount = 0
	p.ticks = 0
	x, _ := p.game.convertXY(playerStartX, 0)
	p.SetX(float32(x))

	p.prop.SetTexture(p.aliveTex)
	p.prop.SetVisible(true)
}

// Hit indicates the player has been hit
func (p *player) Hit() {
	p.exploding = true
	p.ticks = sdl.GetTicks()
}

// MoveLeft moves the player left
func (p *player) MoveLeft() {
	if p.lives == 0 || p.exploding == true {
		return
	}
	if p.X > 0 {
		p.X = p.X - float32(p.Speed.X*p.game.stage.ElapsedTime())
	}
}

// MoveRight moves the player right
func (p *player) MoveRight() {
	if p.lives == 0 || p.exploding == true {
		return
	}
	if p.X < originalWidth {
		p.X = p.X + float32(p.Speed.X*p.game.stage.ElapsedTime())
	}
}
