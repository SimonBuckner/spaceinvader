package main

import (
	"fmt"
	"strconv"

	"github.com/SimonBuckner/spaceinvader/gfx"
	"github.com/veandco/go-sdl2/sdl"
)

type playerClass int

type playerStatus int

const (
	showPlayer playerStatus = iota
	explodePlayer
	hidePlayer
)

type playerShip struct {
	*gfx.Asset

	liveTex *sdl.Texture
	hitTex1 *sdl.Texture
	hitTex2 *sdl.Texture

	status       playerStatus
	explodeCount int
	lastTick     uint32
}

func newPlayerShip(gs *gameState, number int) (*playerShip, error) {
	ship := &playerShip{
		status: hidePlayer,
	}

	ship.Asset = gfx.NewAsset(gs.vp, "player_"+strconv.Itoa(number))
	ship.SetScale(gs.scale)

	err := ship.loadTextures(gs.vp, playerSprite, plrBlowupSprite0, plrBlowupSprite1)
	if err != nil {
		return nil, err
	}

	ship.SetTexture(ship.liveTex)
	return ship, nil
}

func (ship *playerShip) loadTextures(vp *gfx.ViewPort, live, hit1, hit2 *gfx.Bitmap) error {

	var err error
	ship.liveTex, err = live.ToTexture(vp)
	if err != nil {
		return fmt.Errorf("unable to load live1 bitmap")
	}
	ship.hitTex1, err = hit1.ToTexture(vp)
	if err != nil {
		return fmt.Errorf("unable to load live2 bitmap")
	}
	ship.hitTex2, err = hit1.ToTexture(vp)
	if err != nil {
		return fmt.Errorf("unable to load hit bitmap")
	}
	return nil
}

func (ship *playerShip) update(ticks uint32) {

	switch ship.status {
	case showPlayer:
		ship.Show()
	case hidePlayer:
		ship.Hide()
	case explodePlayer:
		if ticks-ship.lastTick > 16 {
			ship.lastTick = ticks
			ship.explodeCount--
		}
		if ship.explodeCount == 0 {
			ship.Hide()
			return
		}
		if ship.explodeCount%2 == 0 {
			ship.Asset.SetTexture(ship.hitTex1)
		} else {
			ship.Asset.SetTexture(ship.hitTex2)
		}
	}
}

func (ship *playerShip) Show() {
	ship.explodeCount = 10
	ship.lastTick = sdl.GetTicks()
	ship.status = explodePlayer
}
func (ship *playerShip) Hide() {
	ship.explodeCount = 10
	ship.lastTick = sdl.GetTicks()
	ship.status = explodePlayer
}

func (ship *playerShip) Explode() {
	ship.explodeCount = 10
	ship.lastTick = sdl.GetTicks()
	ship.status = explodePlayer
}
