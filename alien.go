package main

import (
	"github.com/SimonBuckner/screen2d"
	"github.com/veandco/go-sdl2/sdl"
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
	game       *game
	state      alienState
	breed      alienBreed
	frame      int
	frameStart uint32
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
	case alienCrab:
		a.SetSprite(a.game.sprites.GetSprite(keyAlienSprA0))
	case alienOctopus:
		a.SetSprite(a.game.sprites.GetSprite(keyAlienSprB0))
	case alienSquid:
		a.SetSprite(a.game.sprites.GetSprite(keyAlienSprC0))
	}
}

func (a *alien) reset() {
	a.state = alienAlive
	a.frame = 0
	a.frameStart = sdl.GetTicks()
}

func (a *alien) update(ticks uint32, elapsed float32) {
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

	if ticks-a.frameStart > alienFrameTimer {
		a.frameStart = ticks
		a.frame++
		if a.frame > 1 {
			a.frame = 0
		}
	}
}

func (a *alien) setHit() {
	a.state = alienExploding
}
