package main

import (
	"fmt"

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
	game             *game
	keyb             *screen2d.KBState
	title            *text
	p1score          *text
	p2score          *text
	highScore        int
	hiscore          *text
	p1               *player
	p2               *player
	player           *player
	state            playModeState
	frameStart       uint32
	frame            int
	frameJog         uint32
	timer            int
	timer2           int
	soundDelay       int
	soundTimer       int
	rollShot         *alienShot
	squiglyShot      *alienShot
	plungerShot      *alienShot
	syncShot         int
	rollShotTimer    int
	squiglyShotTimer int
	plungerShotTimer int
}

func newPlayMode(game *game) *playMode {

	pm := &playMode{
		game:        game,
		keyb:        game.screen.GetKBState(),
		p1:          newPlayer(game),
		p2:          newPlayer(game),
		title:       newText(game),
		highScore:   0,
		p1score:     newText(game),
		p2score:     newText(game),
		hiscore:     newText(game),
		rollShot:    newAlienShot(game),
		squiglyShot: newAlienShot(game),
		plungerShot: newAlienShot(game),
	}
	pm.title.load(game.font, game.fontKeys)
	pm.p1score.load(game.font, game.fontKeys)
	pm.hiscore.load(game.font, game.fontKeys)
	pm.p2score.load(game.font, game.fontKeys)

	pm.rollShot.setKind(askRolling)
	pm.squiglyShot.setKind(askSquigly)
	pm.plungerShot.setKind(askPlunger)

	return pm
}

func (pm *playMode) activate() {
	pm.game.screen.ClearFuncs()

	pm.p1.reset()
	pm.player = pm.p1
	pm.game.screen.SetKeyDownFunc(pm.onKeyDown)
	pm.game.screen.SetUpdateFunc(pm.onUpdate)
	pm.game.screen.SetDrawFunc(pm.onDraw)
	pm.title.setText(scoreTitle)
	pm.title.X = 0
	pm.title.Y = 0
	pm.title.visible = true

	pm.p1score.setText(pm.p1.getScore())
	pm.p1score.X = 14
	pm.p1score.Y = 16
	pm.p1score.visible = true

	pm.hiscore.setText(fmt.Sprintf("%04d", pm.highScore))
	pm.hiscore.X = 70
	pm.hiscore.Y = 16
	pm.hiscore.visible = true

	pm.p2score.setText(pm.p2.getScore())
	pm.p2score.X = 140
	pm.p2score.Y = 16
	pm.p2score.visible = true

	pm.frame = 0
	pm.frameStart = sdl.GetTicks()

	pm.state = pmReady
	pm.timer = pmReadyTTL
	pm.timer2 = pmReadyDelayTTL

	pm.rollShot.X = 30
	pm.rollShot.Y = 40
	pm.rollShot.Visible = true
	pm.rollShot.reset()

	pm.squiglyShot.X = 40
	pm.squiglyShot.Y = 40
	pm.squiglyShot.Visible = true
	pm.squiglyShot.reset()

	pm.plungerShot.X = 50
	pm.plungerShot.Y = 40
	pm.plungerShot.Visible = true
	pm.plungerShot.reset()

	pm.syncShot = -1

}

func (pm *playMode) onKeyDown(e *sdl.KeyboardEvent) {
	switch e.Keysym.Scancode {
	case sdl.SCANCODE_ESCAPE:
		pm.game.activate()
	case sdl.SCANCODE_Q:
		pm.game.screen.Close()
	}
}

func (pm *playMode) onUpdate(ticks uint32, elapsed float32) {

	switch pm.state {
	case pmReady:
		pm.updateReady(ticks, elapsed)
	case pmPlaying:
		pm.updatePlaying(ticks, elapsed)
	case pmLevelComplete:
		pm.updateLevelComplete(ticks, elapsed)
	case pmDead:
		pm.updateDead(ticks, elapsed)
	}

}

func (pm *playMode) updateReady(ticks uint32, elapsed float32) {
	pm.timer2--
	if pm.timer2 == 0 {
		if pm.player == pm.p1 {
			pm.p1score.visible = !pm.p1score.visible
		} else {
			pm.p2score.visible = !pm.p1score.visible
		}
		pm.timer2 = pmReadyDelayTTL
		pm.timer--

	}
	if pm.timer == 0 {
		pm.state = pmPlaying
	}
}

func (pm *playMode) updatePlaying(ticks uint32, elapsed float32) {
	if pm.keyb.IsKeyDown(sdl.SCANCODE_LEFT) != pm.keyb.IsKeyDown(sdl.SCANCODE_RIGHT) {
		if pm.keyb.IsKeyDown(sdl.SCANCODE_LEFT) {
			pm.player.moveLeft()
		} else {
			pm.player.moveRight()
		}
	}
	if pm.keyb.IsKeyDown(sdl.SCANCODE_SPACE) {
		pm.player.shot.fire()
	}

	// Move cursor to next alien
	pm.player.alienRack.advanceCursor()

	// decrement general purpose time

	// count remaining aliens
	aliveCount := pm.player.alienRack.remainCount()

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
		pm.rollShot.deltaY = 5.0
		pm.squiglyShot.deltaY = 5.0
		pm.plungerShot.deltaY = 5.0
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
		pm.player.alienRack.stepR = 3.0
	default:
		// Level complete
		pm.state = pmLevelComplete
	}

	// Sync the three alien shots so only one is processed by screen
	pm.syncShot++
	if pm.syncShot >= 3 {
		pm.syncShot = 0
	}
	// Execute game objects

	// Move player
	pm.player.update(ticks, elapsed)

	// Move player shot
	pm.player.shot.update(ticks, elapsed, pm.p1.X)

	// Move rolling shot & plunger shot & either squiggly shot or saucer
	// Saucer appears every 600 frames whilst > 8 aliens and no squiggly shot on the screen
	// One shot moves 4 picels each frame
	switch pm.syncShot {
	case 0:
		pm.rollShot.update(ticks, elapsed, pm.player.X)
	case 1:
		pm.squiglyShot.update(ticks, elapsed, pm.player.X)
	case 2:
		pm.plungerShot.update(ticks, elapsed, pm.player.X)
	}

	if aliveCount < 8 {
		// change shot step to 5 pixels
	}
	if aliveCount == 1 {
		// disable plunger shot
	}

	// Move alien rack
	pm.player.alienRack.update(ticks, elapsed, pm.p1, pm.player.shot)

	pm.p1score.setText(pm.p1.getScore())
	pm.p2score.setText(pm.p2.getScore())
}

func (pm *playMode) updateLevelComplete(ticks uint32, elapsed float32) {
}

func (pm *playMode) updateDead(ticks uint32, elapsed float32) {
}

func (pm *playMode) onDraw() {
	pm.title.drawText()
	pm.p1score.drawText()
	pm.hiscore.drawText()
	pm.p2score.drawText()
	pm.player.Draw()
	pm.player.shot.Draw()
	pm.rollShot.Draw()
	pm.squiglyShot.Draw()
	pm.plungerShot.Draw()
	pm.player.alienRack.drawRack()
}
