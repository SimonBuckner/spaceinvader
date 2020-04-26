package main

import (
	"github.com/SimonBuckner/screen2d"
	"github.com/veandco/go-sdl2/sdl"
)

type playModeState int

const (
	pmReady playModeState = iota
	pmPlaying
	pmLevelComplete
	pmDead
)

const playModeName = "Test Scene"

type playMode struct {
	game        *game
	keyb        *screen2d.KBState
	p1          *player
	p1Shot      *playerShot
	p1AlienRack *alienRack
	title       *text
	state       playModeState
	frameStart  uint32
	frame       int
	frameJog    uint32
	timer       int
	soundDelay  int
	soundTimer  int
}

func newPlayMode(game *game) *playMode {

	pm := &playMode{
		game:        game,
		keyb:        game.screen.GetKBState(),
		p1:          newPlayer(game),
		p1Shot:      newPlayerShot(game),
		p1AlienRack: newAlienRack(game),
		title:       newText(game),
	}

	return pm
}

func (pm *playMode) activate() {
	pm.game.screen.ClearFuncs()

	pm.p1.reset()
	pm.p1Shot.reset()
	pm.p1AlienRack.reset(1)
	pm.game.screen.SetKeyDownFunc(pm.onKeyDown)
	pm.game.screen.SetUpdateFunc(pm.onUpdate)
	pm.game.screen.SetDrawFunc(pm.onDraw)
	pm.title.setText(scoreTitle)
	pm.title.X = 0
	pm.title.Y = 0

	pm.frame = 0
	pm.frameStart = sdl.GetTicks()

}

func (pm *playMode) onKeyDown(e *sdl.KeyboardEvent) {
	switch e.Keysym.Scancode {
	case sdl.SCANCODE_ESCAPE:
		pm.game.activate()
	case sdl.SCANCODE_Q:
		pm.game.screen.Close()
	case sdl.SCANCODE_LEFT:
		pm.p1.moveLeft()
	case sdl.SCANCODE_RIGHT:
		pm.p1.moveRight()
	}
}

func (pm *playMode) onUpdate(ticks uint32, elapsed float32) {

	if pm.keyb.IsKeyDown(sdl.SCANCODE_LEFT) != pm.keyb.IsKeyDown(sdl.SCANCODE_RIGHT) {
		if pm.keyb.IsKeyDown(sdl.SCANCODE_LEFT) {
			pm.p1.moveLeft()
		} else {
			pm.p1.moveRight()
		}
	} else {
		pm.p1.stopMoving()
	}
	if pm.keyb.IsKeyDown(sdl.SCANCODE_SPACE) {
		pm.p1Shot.fire()
	}

	// Move cursor to next alien
	pm.p1AlienRack.advanceCursor()

	// decrement general purpose time

	// count remaining aliens
	aliveCount := pm.p1AlienRack.remainCount()

	// Change rate of step sound
	switch {
	case aliveCount > 50:
		// 52 frames between sounds
	case aliveCount > 43:
		// 46 frames between sounds
	case aliveCount > 36:
		// 39 frames between sounds
	case aliveCount > 28:
		// 34 frames between sounds
	case aliveCount > 22:
		// 28 frames between sounds
	case aliveCount > 17:
		// 24 frames between sounds
	case aliveCount > 13:
		// 21 frames between sounds
	case aliveCount > 10:
		// 19 frames between sounds
	case aliveCount > 8:
		// 16 frames between sounds
	case aliveCount > 7:
		// 14 frames between sounds
	case aliveCount > 6:
		// 13 frames between sounds
	case aliveCount > 5:
		// 12 frames between sounds
	case aliveCount > 4:
		// 11 frames between sounds
	case aliveCount > 3:
		// 9 frames between sounds
	case aliveCount > 2:
		// 7 frames between sounds
	case aliveCount > 1:
		// 5 frames between sounds
	case aliveCount > 0:
		pm.p1AlienRack.stepR = 3.0
	default:
		// Level complete
		pm.state = pmLevelComplete
	}

	// Sync the three alien shots so only one is processed by screen

	// Execute game objects

	// Move player
	pm.p1.update(ticks, elapsed)

	// Move player shot
	pm.p1Shot.update(ticks, elapsed, pm.p1.X)

	// Move rolling shot & plunger shot & either squiggly shot or saucer
	// Saucer appears every 600 frames whilst > 8 aliens and no squiggly shot on the screen
	// One shot moves 4 picels each frame
	if aliveCount < 8 {
		// change shot step to 5 pixels
	}
	if aliveCount == 1 {
		// disable plunger shot
	}

	// Move alien rack
	pm.p1AlienRack.update(ticks, elapsed, pm.p1, pm.p1Shot)

}

func (pm *playMode) movePlayer(ticks uint32, elapsed float32) {

}

func (pm *playMode) movePlayerShot(ticks uint32, elapsed float32) {

}

func (pm *playMode) moveRollingShot(ticks uint32, elapsed float32) {

}

func (pm *playMode) movePlungerShot(ticks uint32, elapsed float32) {

}

func (pm *playMode) moveSquigglyShot(ticks uint32, elapsed float32) {

}

func (pm *playMode) moveAliens(ticks uint32, elapsed float32) {

}

func (pm *playMode) onDraw() {
	pm.title.drawText()
	pm.p1.Draw()
	pm.p1Shot.Draw()
	pm.p1AlienRack.draw()
}
