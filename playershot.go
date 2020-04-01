package main

import (
	"github.com/SimonBuckner/screen2d"
	"github.com/veandco/go-sdl2/sdl"
)

const (
	playerShotSpeed   float32 = 500
	playerShotMissedY float32 = 25
	// 	shotExplodeTTL = 15
	// 	shotExplodeTTL = 60
)

type playerShotState int

const (
	playerShotReady playerShotState = iota
	playerShotFired
	playerShotMissed
	playerShotHit
)

type playerShot struct {
	*screen2d.Entity
	game         *game
	direction    float32
	state        playerShotState
	width        int32
	explodeTimer uint32
	explodeCount int
}

func newPlayerShot(game *game) *playerShot {
	p := &playerShot{
		Entity: screen2d.NewEntity(),
		game:   game,
		state:  playerShotReady,
	}

	p.Scale = game.scale
	p.SetCalcScreenXYFunc(translatePos)
	p.setAvailable(playerX)
	return p
}

func (p *playerShot) update(ticks uint32, elapsed float32, shipX float32) {
	switch p.state {
	case playerShotReady:
		p.X = shipX + float32(playerwidth-p.width+1)/2
	case playerShotFired:
		p.Y = p.Y - (playerShotSpeed * elapsed)
		if p.Y < playerShotMissedY {
			p.setMissed()
		}
	case playerShotMissed:
		if (ticks - p.explodeTimer) > 500 {
			p.setAvailable(shipX)
		}
	}
}

func (p *playerShot) fire() {
	if p.state != playerShotReady {
		return
	}
	p.state = playerShotFired
}

func (p *playerShot) setAvailable(shipX float32) {
	p.SetSprite(p.game.sprites.GetSprite(keyPlayerShotSpr))
	p.width = p.Sprite.GetPitch()
	p.state = playerShotReady
	p.X = shipX + float32(playerwidth-p.width+1)/2
	p.Y = float32(playerY - 2)
}

func (p *playerShot) setMissed() {
	p.state = playerShotMissed
	p.explodeTimer = sdl.GetTicks()
	p.SetSprite(p.game.sprites.GetSprite(keyShotExploding))
	p.X = p.X - float32(p.Sprite.GetPitch()/2)
}
