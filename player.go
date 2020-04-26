package main

import (
	"fmt"

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
	shot      *playerShot
	alienRack *alienRack

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
		shot:      newPlayerShot(game),
		alienRack: newAlienRack(game),
		lives:     3,
		extraUsed: false,
		state:     playerAlive,
	}
	p.SetCalcScreenXYFunc(translatePos)
	p.Scale = game.scale
	p.shot.Scale = game.scale
	p.alienRack.scale = game.scale
	return p
}

func (p *player) reset() {
	p.X = playerX
	p.Y = playerY
	p.SetSprite(p.game.sprites.GetSprite(keyPlayerSprite))
	p.SetPos(playerX, playerY, 0)
	p.score = 0
	p.shot.reset()
	p.alienRack.reset(1)
}

func (p *player) update(ticks uint32, elapsed float32) {
	delta := (p.direction * shipSpeed * elapsed)
	newX := p.X + delta
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

func (p *player) getScore() string {
	return fmt.Sprintf("%04d", p.score)
}
