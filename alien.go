package main

import (
	"github.com/SimonBuckner/screen2d"
)

type alienState int

const (
	alienAlive alienState = iota
	alienExploding
	alienDead
)

type alienBreed int

const (
	alienSquid alienBreed = iota
	alienCrab
	alienOctopus
	// alienSpecial
)

type alien struct {
	*screen2d.Entity
	game  *game
	state alienState
	breed alienBreed
	frame int
	score int
	// moved bool
}

func newAlien(game *game) *alien {
	a := &alien{
		Entity: screen2d.NewEntity(),
		game:   game,
		frame:  0,
	}
	a.Scale = game.scale

	return a
}

func (a *alien) setBreed(breed alienBreed) {
	a.breed = breed
	switch a.breed {
	case alienOctopus:
		a.SetSprite(a.game.sprites.GetSprite(keyAlienSprB0))
		a.score = 10
	case alienCrab:
		a.SetSprite(a.game.sprites.GetSprite(keyAlienSprA0))
		a.score = 20
	case alienSquid:
		a.SetSprite(a.game.sprites.GetSprite(keyAlienSprC0))
		a.score = 30
	}
}

func (a *alien) reset() {
	a.state = alienAlive
	a.frame = 0
	a.Visible = true
}

func (a *alien) update(ticks uint32, elapsed float32) {
	switch a.state {
	case alienAlive:
		if a.frame == 0 {
			switch a.breed {
			case alienCrab:
				a.SetSprite(a.game.sprites.GetSprite(keyAlienSprA0))
			case alienOctopus:
				a.SetSprite(a.game.sprites.GetSprite(keyAlienSprB0))
			case alienSquid:
				a.SetSprite(a.game.sprites.GetSprite(keyAlienSprC0))
			}

		} else {
			switch a.breed {
			case alienCrab:
				a.SetSprite(a.game.sprites.GetSprite(keyAlienSprA1))
			case alienOctopus:
				a.SetSprite(a.game.sprites.GetSprite(keyAlienSprB1))
			case alienSquid:
				a.SetSprite(a.game.sprites.GetSprite(keyAlienSprC1))
			}

		}
	case alienExploding:
		a.SetSprite((a.game.sprites.GetSprite(keyAlienExplode)))
	case alienDead:
		a.Visible = false
	}
}

func (a *alien) setHit() {
	a.state = alienExploding
}

func (a *alien) setDead() {
	a.state = alienDead
}
