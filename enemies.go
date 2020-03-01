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
	*gfx.Asset

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
	ship.SetScale(gs.scale)

	switch class {
	case enemyClassA:
		ship.Asset = gfx.NewAsset(gs.vp, "alien_a")
		if err := ship.loadTextures(gs.vp, alienSprA0, alienSprA1, alienExplode); err != nil {
			return nil, err
		}
	case enemyClassB:
		ship.Asset = gfx.NewAsset(gs.vp, "alien_b")
		if err := ship.loadTextures(gs.vp, alienSprB0, alienSprB1, alienExplode); err != nil {
			return nil, err
		}
	case enemyClassC:
		ship.Asset = gfx.NewAsset(gs.vp, "alien_c")
		if err := ship.loadTextures(gs.vp, alienSprC0, alienSprC1, alienExplode); err != nil {
			return nil, err
		}
	}

	ship.SetTexture(ship.liveTex1)
	return ship, nil
}

func (ship *enemyShip) loadTextures(vp *gfx.ViewPort, live1, live2, hit *gfx.Bitmap) error {

	var err error
	ship.liveTex1, err = alienSprA0.ToTexture(vp)
	if err != nil {
		return fmt.Errorf("unable to load live1 bitmap")
	}
	ship.liveTex1, err = alienSprA0.ToTexture(vp)
	if err != nil {
		return fmt.Errorf("unable to load live2 bitmap")
	}
	ship.liveTex1, err = alienSprA0.ToTexture(vp)
	if err != nil {
		return fmt.Errorf("unable to load hit bitmap")
	}
	return nil
}
