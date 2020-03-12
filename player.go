package main

import (
	"github.com/SimonBuckner/spaceinvader/gfx"
	"github.com/veandco/go-sdl2/sdl"
)

const (
	playerHeight = 8
	playerwidth  = 16
	playerSpeed  = 60

	// Start positions for player props
	playerX = 1
	playerY = originalHeight - (playerHeight * 4)
	// bankX   = 1
	// bankY   = originalHeight - playerHeight

	startLives = 3
)

type player struct {
	*gfx.Actor
	ship        *gfx.Prop
	aliveTex    *sdl.Texture
	explode1Tex *sdl.Texture
	explode2Tex *sdl.Texture

	score        int
	extraAvail   bool
	hit          bool
	explodeCount int
	lives        int
	ticks        uint32
	width        int32
	height       int32
}

func newPlayer(game *game) *player {
	p := &player{
		Actor: gfx.NewActor("player"),
		ship:  gfx.NewProp("player ship", nil),
	}
	p.StartEventHandler = p.onStart
	p.StopEventHandler = p.onStop
	p.UpdateEventHandler = p.onUpdate
	return p
}

func (p *player) onStart() {
	stage := p.Scene.Stage

	p.aliveTex, _ = playerSprite.ToTexture(stage)
	p.explode1Tex, _ = plrBlowupSprite0.ToTexture(stage)
	p.explode2Tex, _ = plrBlowupSprite1.ToTexture(stage)

	_, _, w, h, _ := p.aliveTex.Query()

	p.width = w
	p.height = h

	p.reset()

}

func (p *player) onStop() {
	p.RemoveProp(p.ship)
}

func (p *player) reset() {
	p.Pos.SetInt32(playerX, playerY, 0)
	p.score = 0
	p.lives = 3
	p.extraAvail = true
	p.hit = false
	p.explodeCount = 0
	p.ticks = 0
	p.ship.Texture = p.aliveTex
	p.Visible = true
	p.AddProp(p.ship)
}

func (p *player) onUpdate(ticks uint32) {
	if !p.Visible {
		return
	}
	x, y, _ := p.Pos.Int32()
	x1, y1 := convertXY(p.Scene, x, y)
	p.ship.Pos.SetInt32(x1, y1, 0)

	if !p.hit {
		return
	}

	if p.explodeCount > 10 {
		p.lives--
		if p.lives > 0 {
			p.hit = false
			p.ship.Texture = p.aliveTex
			return
		}
		p.setDead()
		return
	}

	if ticks-p.ticks > 32 {
		p.explodeCount++
		p.ticks = ticks
	}

	if p.explodeCount%2 == 0 {
		p.ship.Texture = p.explode1Tex
	} else {
		p.ship.Texture = p.explode2Tex
	}
}

func (p *player) setHit() {
	p.hit = true
	p.ticks = sdl.GetTicks()
	p.explodeCount = 0
}

func (p *player) setDead() {
	p.Visible = false
	p.RemoveProp(p.ship)
}
func (p *player) moveLeft() {

	if p.lives == 0 || p.hit == true {
		return
	}
	if p.Pos.X > 0 {
		p.Pos.X = p.Pos.X - float32(playerSpeed*p.Scene.ElapsedTime())
		return
	}
	p.Pos.X = 0
}

func (p *player) moveRight() {
	if p.lives == 0 || p.hit == true {
		return
	}
	if int32(p.Pos.X)+p.width < originalWidth {
		p.Pos.X = p.Pos.X + float32(playerSpeed*p.Scene.ElapsedTime())
		return
	}
	p.Pos.X = float32(originalWidth - p.width)
}
