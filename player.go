package main

import (
	"fmt"

	"github.com/SimonBuckner/spaceinvader/gfx"
	"github.com/veandco/go-sdl2/sdl"
)

const (
	playerHeight     = 8
	playerwidth      = 16
	playerSpeed      = 60
	playerExplodeTTL = 10
	shotSpeed        = 540
	shotExplodeTTL   = 40
	shotMissedY      = 25

	// Start positions for player props
	playerX = 1
	playerY = originalHeight - (playerHeight * 4)

	startLives = 3
)

type shotState int

const (
	shotAvailable shotState = iota + 1
	shotInFlight
	shotExplode
)

type player struct {
	*gfx.Actor
	ship         *gfx.Prop
	shipAlive    *sdl.Texture
	shipExplode1 *sdl.Texture
	shipExplode2 *sdl.Texture

	shot             *gfx.Prop
	shotAvailable    *sdl.Texture
	shotExploding    *sdl.Texture
	shotState        shotState
	shotExplodeCount int

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
		Actor:     gfx.NewActor("player"),
		ship:      gfx.NewProp("player ship", nil, game.transformXYZ),
		shot:      gfx.NewProp("player shot", nil, game.transformXYZ),
		shotState: shotAvailable,
	}
	return p
}

func (p *player) Start(scene *gfx.Scene) {
	p.Scene = scene
	p.Scale = scene.Scale()
	p.ship.Scale = scene.Scale()
	p.shot.Scale = scene.Scale()

	stage := scene.Stage

	p.shipAlive, _ = playerSprite.ToTexture(stage)
	p.shipExplode1, _ = plrBlowupSprite0.ToTexture(stage)
	p.shipExplode2, _ = plrBlowupSprite1.ToTexture(stage)
	p.shotAvailable, _ = playerShotSpr.ToTexture(stage)
	p.shotExploding, _ = shotExploding.ToTexture(stage)

	_, _, w, h, _ := p.shipAlive.Query()

	p.width = w
	p.height = h

	p.reset()

}

// Update ..
func (p *player) Update(ticks uint32) {
	p.updatePlayer(ticks)
	p.updateShot(ticks)
}

func (p *player) updatePlayer(ticks uint32) {
	if !p.Visible {
		return
	}
	x, y, _ := p.Pos.Int32()
	p.ship.Pos.SetInt32(x, y, 0)
	if !p.hit {
		return
	}

	if p.explodeCount > playerExplodeTTL {
		p.lives--
		if p.lives > 0 {
			p.hit = false
			p.ship.Texture = p.shipAlive
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
		p.ship.Texture = p.shipExplode1
	} else {
		p.ship.Texture = p.shipExplode2
	}
}

func (p *player) updateShot(ticks uint32) {

	switch p.shotState {
	case shotAvailable:
		w, _ := p.width, p.height
		x, y, _ := p.Pos.Int32()
		x = x + int32(w/2)
		y = y - 2
		p.shot.Pos.SetInt32(x, y, 0)
	case shotInFlight:
		if p.shot.Pos.Y > shotMissedY {
			p.shot.Pos.Y = p.shot.Pos.Y - float32(shotSpeed*p.Scene.ElapsedTime())
		} else {
			p.shotState = shotExplode
			p.shot.Texture = p.shotExploding
		}
	case shotExplode:
		p.shotExplodeCount++
		if p.shotExplodeCount > shotExplodeTTL {
			p.shotState = shotAvailable
			p.shot.Texture = p.shotAvailable
			p.shotExplodeCount = 0
		}
	}
}

// Draw ..
func (p *player) Draw() {
	if p.Visible {
		p.ship.Draw(p.Scene.Renderer())
		p.shot.Draw(p.Scene.Renderer())
	}
}

func (p *player) reset() {
	p.Pos.SetInt32(playerX, playerY, 0)
	p.score = 0
	p.lives = 3
	p.extraAvail = true
	p.hit = false
	p.explodeCount = 0
	p.ticks = 0
	p.ship.Texture = p.shipAlive
	p.Visible = true

	p.shot.Texture = p.shotAvailable
	p.shotExplodeCount = 0
}

func (p *player) setHit() {
	p.hit = true
	p.ticks = sdl.GetTicks()
	p.explodeCount = 0
}

func (p *player) setDead() {
	p.Visible = false
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

func (p *player) fireShot() {
	if p.shotState == shotAvailable {
		fmt.Println("fireing")
		p.shotState = shotInFlight
	}
}
