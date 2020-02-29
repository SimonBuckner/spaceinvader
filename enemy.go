package main

import (
	"fmt"

	"github.com/SimonBuckner/spaceinvader/gfx"
)

type enemyClass int

const (
	enemyClassA enemyClass = iota
	enemyClassB
	enemyClassC
)

type enemyShip struct {
	*gfx.Asset
	enemyClass
	hit  bool
	dead bool
	ttl  uint32
}

func newAlien(s *gameState, class enemyClass) (*enemyShip, error) {
	ship := &enemyShip{
		enemyClass: class,
		hit:        false,
		dead:       false,
		ttl:        0,
	}
	var name string
	var asset *gfx.Asset
	var err error

	switch class {
	case enemyClassA:
		name = "AlienA"
		asset, err = gfx.AssetFromBitmaps(s.vp, alienSprA0, alienSprA1, alienExplode)
	case enemyClassB:
		name = "AlienA"
		asset, err = gfx.AssetFromBitmaps(s.vp, alienSprB0, alienSprB1, alienExplode)
	case enemyClassC:
		name = "AlienA"
		asset, err = gfx.AssetFromBitmaps(s.vp, alienSprC0, alienSprC1, alienExplode)
	}
	if err != nil {
		return nil, fmt.Errorf("error creating alien %v; error %v", name, err)
	}

	asset.SetScale(s.scale)
	asset.Name = name
	ship.Asset = asset
	return ship, nil
}
