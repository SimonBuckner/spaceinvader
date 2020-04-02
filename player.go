package main

import (
	"github.com/SimonBuckner/screen2d"
)

type playerState int

const (
	playerReady playerState = iota
	playerAlive
	playerHit
	plyaerDead
)

type player struct {
	*screen2d.Entity
	game      *game
	direction float32
	score     int
	lives     int
	extraUsed bool
	state     playerState
}

func newPlayer(game *game) *player {
	p := &player{
		Entity:    screen2d.NewEntity(),
		game:      game,
		lives:     3,
		extraUsed: false,
		state:     playerAlive,
	}
	p.SetCalcScreenXYFunc(translatePos)
	p.Scale = game.scale

	return p
}

func (p *player) reset() {
	p.X = playerX
	p.Y = playerY
	p.SetSprite(p.game.sprites.GetSprite(keyPlayerSprite))
	p.SetPos(playerX, playerY, 0)
}

func (p *player) update(ticks uint32, elapsed float32) {
	newX := p.X + (p.direction * shipSpeed * elapsed)
	if newX > 0 && int(newX) < originalWidth-playerwidth {
		p.X = newX
	}
}

func (p *player) moveLeft() {
	p.direction = -1
}

func (p *player) moveRight() {
	p.direction = +1
}

func (p *player) stopMoving() {
	p.direction = 0
}

// type shipState int

// const (
// 	shipReady shipState = iota
// 	shipAlive
// 	shipHit
// 	shipDead
// )

// type shotState int

// const (
// 	shotAvailable shotState = iota
// 	shotInFlight
// 	shotHit
// 	shotMissed
// )

// type player struct {
// *gfx.Actor
// ship *gfx.Entity
// shot *gfx.Entity
// ship             *gfx.Prop
// shipAlive        *sdl.Texture
// shipExplode1     *sdl.Texture
// shipExplode2     *sdl.Texture
// shipState        shipState
// shipTimer        uint32
// shipExplodeCount int

// shot             *gfx.Prop
// shotAvailable    *sdl.Texture
// shotExploding    *sdl.Texture
// shotState        shotState
// shotTimer        uint32
// shotExplodeCount int

// 	score      int
// 	lives      int
// 	extraAvail bool

// 	width  int32
// 	height int32
// }

// func (p *player) Start(scene *gfx.Scene) {
// 	p.Scene = scene
// 	p.Scale = scene.Scale()
// 	p.ship.Scale = scene.Scale()
// 	p.shot.Scale = scene.Scale()

// 	stage := scene.Stage

// 	p.shipAlive, _ = playerSprite.ToTexture(stage)
// 	p.shipExplode1, _ = plrBlowupSprite0.ToTexture(stage)
// 	p.shipExplode2, _ = plrBlowupSprite1.ToTexture(stage)
// 	p.shotAvailable, _ = playerShotSpr.ToTexture(stage)
// 	p.shotExploding, _ = shotExploding.ToTexture(stage)

// 	_, _, w, h, _ := p.shipAlive.Query()

// 	p.width = w
// 	p.height = h

// 	p.resetShip()
// }

// // Update ..
// func (p *player) Update(ticks uint32) {
// 	p.updateShip(ticks)
// 	p.updateShot(ticks)
// }

// func (p *player) updateShip(ticks uint32) {
// 	if !p.Visible {
// 		return
// 	}

// 	x, y, _ := p.Pos.Int32()
// 	p.ship.Pos.SetInt32(x, y, 0)

// 	if p.shipState == shipAlive {
// 		return
// 	}

// 	if p.shipState == shipHit {
// 		if p.shipExplodeCount == 0 {
// 			p.lives--
// 			if p.lives > 0 {
// 				p.shipState = shipAlive
// 				p.ship.Texture = p.shipAlive
// 				return
// 			}
// 			p.setDead()
// 			return
// 		}
// 		if ticks-p.shipTimer > 16*2 {
// 			p.shipExplodeCount--
// 			p.shipTimer = ticks
// 			if p.shipExplodeCount%2 == 0 {
// 				p.ship.Texture = p.shipExplode1
// 			} else {
// 				p.ship.Texture = p.shipExplode2
// 			}
// 		}
// 	}
// }

// func (p *player) updateShot(ticks uint32) {

// 	switch p.shotState {
// 	case shotAvailable:
// 		w, _ := p.width, p.height
// 		x, y, _ := p.Pos.Int32()
// 		x = x + int32(w/2)
// 		y = y - 2
// 		p.shot.Pos.SetInt32(x, y, 0)
// 	case shotInFlight:
// 		if p.shot.Pos.Y > shotMissedY {
// 			p.shot.Pos.Y = p.shot.Pos.Y - float32(shotSpeed*p.Scene.ElapsedTime())
// 		} else {
// 			p.setShotMissed()
// 		}
// 	case shotHit, shotMissed:
// 		if p.shotExplodeCount == 0 {
// 			p.resetShot()
// 			return
// 		}
// 		if ticks-p.shotTimer > 16 {
// 			p.shotExplodeCount--
// 		}
// 	}
// }

// func (p *player) resetShip() {
// 	p.Pos.SetInt32(playerX, playerY, 0)
// 	p.score = 0
// 	p.lives = 3
// 	p.extraAvail = true

// 	p.ship.Texture = p.shipAlive
// 	p.shipState = shipAlive

// 	p.Visible = true
// 	p.resetShot()
// }

// func (p *player) resetShot() {
// 	p.shot.Texture = p.shotAvailable
// 	p.shotState = shotAvailable

// }

// func (p *player) setDead() {
// 	p.Visible = false
// }

// func (p *player) fireShot() {
// 	if p.shotState == shotAvailable {
// 		p.shotState = shotInFlight
// 	}
// }

// func (p *player) setShipHit() {
// 	p.shipState = shipHit
// 	p.shipTimer = sdl.GetTicks()
// 	p.shipExplodeCount = shipExplodeTTL
// }

// func (p *player) setShotMissed() {
// 	p.shotState = shotMissed
// 	p.shotTimer = sdl.GetTicks()
// 	p.shotExplodeCount = shotExplodeTTL
// 	p.shot.Texture = p.shotExploding
// 	_, _, w, _, _ := p.shotExploding.Query()
// 	p.shot.Pos.X = p.shot.Pos.X - float32(w/2)
// }

// func (p *player) setShotHit() {
// 	p.shotState = shotHit
// 	p.shotTimer = sdl.GetTicks()
// 	p.shotExplodeCount = shotExplodeTTL
// 	p.shot.Texture = nil
// }
