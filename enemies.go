package main

import (
	"fmt"

	"github.com/SimonBuckner/spaceinvader/gfx"
	"github.com/veandco/go-sdl2/sdl"
)

type enemyClass int

const (
	enemyClassA enemyClass = iota
	enemyClassB
	enemyClassC
)

type enemyStatus int

const (
	livingEnemy enemyStatus = iota
	hitEnemy
	deadenemy
)

type enemyShip struct {
	*gfx.Prop

	liveTex1 *sdl.Texture
	liveTex2 *sdl.Texture
	hitTex   *sdl.Texture

	class       enemyClass
	state       enemyStatus
	stateChange uint32
}

func newEnemyShip(gs *gameState, class enemyClass) (*enemyShip, error) {
	ship := &enemyShip{
		class: class,
		state: livingEnemy,
	}

	switch class {
	case enemyClassA:
		ship.Prop = gfx.NewProp(gs.stage, "alien_a", nil)
		if err := ship.loadTextures(gs.stage, alienSprA0, alienSprA1, alienExplode); err != nil {
			return nil, err
		}
	case enemyClassB:
		ship.Prop = gfx.NewProp(gs.stage, "alien_b", nil)
		if err := ship.loadTextures(gs.stage, alienSprB0, alienSprB1, alienExplode); err != nil {
			return nil, err
		}
	case enemyClassC:
		ship.Prop = gfx.NewProp(gs.stage, "alien_c", nil)
		if err := ship.loadTextures(gs.stage, alienSprC0, alienSprC1, alienExplode); err != nil {
			return nil, err
		}
	}

	ship.SetTexture(ship.liveTex1)
	return ship, nil
}

func (ship *enemyShip) loadTextures(stage *gfx.Stage, live1, live2, hit *gfx.Bitmap) error {

	var err error
	ship.liveTex1, err = alienSprA0.ToTexture(stage)
	if err != nil {
		return fmt.Errorf("unable to load live1 bitmap")
	}
	ship.liveTex1, err = alienSprA0.ToTexture(stage)
	if err != nil {
		return fmt.Errorf("unable to load live2 bitmap")
	}
	ship.liveTex1, err = alienSprA0.ToTexture(stage)
	if err != nil {
		return fmt.Errorf("unable to load hit bitmap")
	}
	return nil
}
