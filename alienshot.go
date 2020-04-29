package main

import (
	"github.com/SimonBuckner/screen2d"
)

type alienShotState int

const (
	asReady alienShotState = iota
	asFired
	asMissed
	asHit
)

type alienShotKind int

const (
	askRolling alienShotKind = iota
	askSquigly
	askPlunger
)

type alienShot struct {
	*screen2d.Entity
	game      *game
	direction float32
	state     alienShotState
	width     int32
	kind      alienShotKind
	timer     int
	frame     int
	frames    []*screen2d.Sprite
	deltaY    float32
	// explodeTimer uint32
	// explodeCount int
}

func newAlienShot(game *game) *alienShot {
	as := &alienShot{
		Entity: game.es.NewEntity(),
		game:   game,
		state:  asReady,
		frames: make([]*screen2d.Sprite, 4),
	}
	as.setReady(playerX)
	return as
}

func (as *alienShot) setKind(kind alienShotKind) {
	as.kind = kind
	as.deltaY = 4.0
	switch kind {
	case askRolling:
		as.frames[0] = as.game.sprites.GetSprite(keyRollShot0)
		as.frames[1] = as.game.sprites.GetSprite(keyRollShot1)
		as.frames[2] = as.game.sprites.GetSprite(keyRollShot2)
		as.frames[3] = as.game.sprites.GetSprite(keyRollShot3)
	case askSquigly:
		as.frames[0] = as.game.sprites.GetSprite(keySquiglyShot0)
		as.frames[1] = as.game.sprites.GetSprite(keySquiglyShot1)
		as.frames[2] = as.game.sprites.GetSprite(keySquiglyShot2)
		as.frames[3] = as.game.sprites.GetSprite(keySquiglyShot3)
	case askPlunger:
		as.frames[0] = as.game.sprites.GetSprite(keyPlungerShot0)
		as.frames[1] = as.game.sprites.GetSprite(keyPlungerShot1)
		as.frames[2] = as.game.sprites.GetSprite(keyPlungerShot2)
		as.frames[3] = as.game.sprites.GetSprite(keyPlungerShot3)
	}
}

func (as *alienShot) reset() {
	as.state = asFired
	as.timer = 4
	as.frame = 0
	as.SetSprite(as.frames[0])
}

func (as *alienShot) update(ticks uint32, elapsed float32, shipX float32) {
	switch as.state {
	case asReady:

	case asFired:
		if as.timer == 0 {
			as.timer = 4
			as.frame++
			if as.frame > 3 {
				as.frame = 0
			}
			as.SetSprite(as.frames[as.frame])
		}
		as.timer--
		as.Y = as.Y + 4
		// if as.Y > originalHeight {
		// 	as.Y = alienShotMissedY
		// 	as.setMissed()
		// }
	case asMissed, asHit:
		// if (ticks - as.explodeTimer) > alienShotMissedTTL {
		// 	as.setAvailable(shipX)
		// }
	}
}

func (as *alienShot) fire() {
	if as.state != asReady {
		return
	}
	as.state = asFired
}

func (as *alienShot) setReady(alienX float32) {
	as.state = asReady
	// as.width = as.Sprite.GetPitch()
	// as.X = alienX + float32(playerwidth-as.width+1)/2
	// as.Y = float32(playerY - 2)
	as.Visible = true
}

func (as *alienShot) setMissed() {
	// as.state = alienShotMissed
	// as.explodeTimer = sdl.GetTicks()
	// as.SetSprite(as.game.sprites.GetSprite(keyShotExploding))
	// as.X = as.X - float32(as.Sprite.GetPitch()/2)
}

func (as *alienShot) setHit() {
	// as.state = alienShotHit
	// as.explodeTimer = sdl.GetTicks()
	// as.Visible = false
	// as.Y = 0
}
