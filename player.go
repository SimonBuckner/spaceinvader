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
	delta := (p.direction * shipSpeed * elapsed)
	newX := p.X + delta
	if delta > 0.6 || delta < -0.6 {
		fmt.Printf("Distance/Elapsed/Ticks: %f / %f / %d \n", delta, elapsed, ticks)
	}
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
