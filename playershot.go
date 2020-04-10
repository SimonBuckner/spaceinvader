package main

import (
	"fmt"

	"github.com/SimonBuckner/screen2d"
	"github.com/veandco/go-sdl2/sdl"
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
	ps := &playerShot{
		Entity: screen2d.NewEntity(),
		game:   game,
		state:  playerShotReady,
	}

	ps.Scale = game.scale
	ps.SetCalcScreenXYFunc(translatePos)
	ps.setAvailable(playerX)
	return ps
}

func (ps *playerShot) reset() {
	ps.setAvailable(playerX)
}

func (ps *playerShot) update(ticks uint32, elapsed float32, shipX float32) {
	switch ps.state {
	case playerShotReady:
		ps.X = shipX + float32(playerwidth-ps.width+1)/2
	case playerShotFired:
		ps.Y = ps.Y - (playerShotSpeed * elapsed)
		if ps.Y < playerShotMissedY {
			ps.setMissed()
		}
	case playerShotMissed, playerShotHit:
		fmt.Printf("shot missed/hit update %d\n", (ticks - ps.explodeTimer))
		if (ticks - ps.explodeTimer) > playerShotMissedTTL {
			ps.setAvailable(shipX)
		}
	}
}

func (ps *playerShot) fire() {
	if ps.state != playerShotReady {
		return
	}
	ps.state = playerShotFired
}

func (ps *playerShot) setAvailable(shipX float32) {
	ps.SetSprite(ps.game.sprites.GetSprite(keyPlayerShotSpr))
	ps.width = ps.Sprite.GetPitch()
	ps.state = playerShotReady
	ps.X = shipX + float32(playerwidth-ps.width+1)/2
	ps.Y = float32(playerY - 2)
	ps.Visible = true
}

func (ps *playerShot) setMissed() {
	ps.state = playerShotMissed
	ps.explodeTimer = sdl.GetTicks()
	ps.SetSprite(ps.game.sprites.GetSprite(keyShotExploding))
	ps.X = ps.X - float32(ps.Sprite.GetPitch()/2)
}

func (ps *playerShot) setHit() {
	ps.state = playerShotHit
	ps.explodeTimer = sdl.GetTicks()
	ps.Visible = false
}
