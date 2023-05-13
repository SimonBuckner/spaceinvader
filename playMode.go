package main

import (
	"fmt"
	"strconv"

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
	title            *screen2d.TextEntity
	p1score          *screen2d.TextEntity
	p2score          *screen2d.TextEntity
	highScore        int
	hiscore          *screen2d.TextEntity
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
	alienShots       []*alienShot
	syncShot         int
	rollShotTimer    int
	squiglyShotTimer int
	plungerShotTimer int
}

func newPlayMode(game *game) *playMode {

	pm := &playMode{
		game:       game,
		keyb:       game.screen.GetKBState(),
		p1:         newPlayer(game),
		p2:         newPlayer(game),
		title:      game.es.NewTextEntity(),
		highScore:  0,
		p1score:    game.es.NewTextEntity(),
		p2score:    game.es.NewTextEntity(),
		hiscore:    game.es.NewTextEntity(),
		alienShots: make([]*alienShot, 3),
	}
	pm.title.LoadAtlas(game.font, game.fontKeys, 8)
	pm.p1score.LoadAtlas(game.font, game.fontKeys, 8)
	pm.hiscore.LoadAtlas(game.font, game.fontKeys, 8)
	pm.p2score.LoadAtlas(game.font, game.fontKeys, 8)

	rollShot := newAlienShot(game)
	squiglyShot := newAlienShot(game)
	plungerShot := newAlienShot(game)

	rollShot.setKind(askRolling)
	squiglyShot.setKind(askSquigly)
	plungerShot.setKind(askPlunger)

	pm.alienShots[0] = rollShot
	pm.alienShots[1] = squiglyShot
	pm.alienShots[2] = plungerShot
	return pm
}

func (pm *playMode) activate() {
	pm.game.screen.ClearFuncs()
	pm.game.screen.SetKeyDownFunc(pm.onKeyDown)
	pm.game.screen.SetUpdateFunc(pm.onUpdate)
	pm.game.screen.SetDrawFunc(pm.onDraw)

	pm.p1.reset()
	pm.player = pm.p1
	pm.title.SetText(scoreTitle)
	pm.title.X = scoreTitleX
	pm.title.Y = scoreTitleY
	pm.title.Visible = true

	pm.p1score.SetText(pm.p1.getScore())
	pm.p1score.X = scoreP1X
	pm.p1score.Y = scoreY
	pm.p1score.Visible = true

	pm.hiscore.SetText(fmt.Sprintf("%04d", pm.highScore))
	pm.hiscore.X = scoreHiX
	pm.hiscore.Y = scoreY
	pm.hiscore.Visible = true

	pm.p2score.SetText(pm.p2.getScore())
	pm.p2score.X = scoreP2X
	pm.p2score.Y = scoreY
	pm.p2score.Visible = true

	pm.frame = 0
	pm.frameStart = sdl.GetTicks()

	pm.state = pmReady
	pm.timer = pmReadyTTL
	pm.timer2 = pmReadyDelayTTL

	x := float32(30)
	for i := range pm.alienShots {
		pm.alienShots[i].X = x
		pm.alienShots[i].Y = 40
		pm.alienShots[i].Visible = true
		pm.alienShots[i].reset()
		x = x + 20.0
	}

	pm.syncShot = 0

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
		sdl.Delay(5)
	case pmDead:
		pm.updateDead(ticks, elapsed)
	}

}

func (pm *playMode) updateReady(ticks uint32, elapsed float32) {
	pm.timer2--
	if pm.timer2 == 0 {
		if pm.player == pm.p1 {
			pm.p1score.Visible = !pm.p1score.Visible
		} else {
			pm.p2score.Visible = !pm.p1score.Visible
		}
		pm.timer2 = pmReadyDelayTTL
		pm.timer--

	}
	if pm.timer == 0 {
		pm.state = pmPlaying
	}
}

func (pm *playMode) updatePlaying(ticks uint32, elapsed float32) {
	if pm.keyb.IsKeyDown(uint8(sdl.SCANCODE_LEFT)) != pm.keyb.IsKeyDown(uint8(sdl.SCANCODE_RIGHT)) {
		if pm.keyb.IsKeyDown(uint8(sdl.SCANCODE_LEFT)) {
			pm.player.moveLeft()
		} else {
			pm.player.moveRight()
		}
	}
	if pm.keyb.IsKeyDown(uint8(sdl.SCANCODE_SPACE)) {
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
		for i := range pm.alienShots {
			pm.alienShots[i].deltaY = 5.0
		}
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
		fmt.Println("Level complete")
		pm.state = pmLevelComplete
	}

	// Execute game objects

	// Move player
	pm.player.update(ticks, elapsed)

	// Move player shot
	pm.player.shot.update(ticks, elapsed, pm.p1.X)

	// Move rolling shot & plunger shot & either squiggly shot or saucer
	// Saucer appears every 600 frames whilst > 8 aliens and no squiggly shot on the screen
	// One shot moves 4 picels each frame
	pm.alienShots[pm.syncShot].update(ticks, elapsed, pm.player.X)

	// Sync the three alien shots so only one is processed by screen
	pm.syncShot++
	if pm.syncShot >= 3 {
		pm.syncShot = 0
	}

	for i := range pm.alienShots {

		aShot := pm.alienShots[i]
		if screen2d.CheckBoxHit(aShot, pm.player.shot) {
			if screen2d.CheckPixelHit(aShot, pm.player.shot) {
				fmt.Println("Shot " + strconv.Itoa(i) + " hit player shot")
			}
		}

		if screen2d.CheckBoxHit(aShot, pm.player) {
			if screen2d.CheckPixelHit(aShot, pm.player) {
				fmt.Println("Shot " + strconv.Itoa(i) + " hit player")
			}
		}
	}

	// Move alien rack
	pm.player.alienRack.update(ticks, elapsed, pm.p1, pm.player.shot)

	pm.p1score.SetText(pm.p1.getScore())
	pm.p2score.SetText(pm.p2.getScore())
}

func (pm *playMode) updateLevelComplete(ticks uint32, elapsed float32) {
}

func (pm *playMode) updateDead(ticks uint32, elapsed float32) {
}

func (pm *playMode) onDraw() {
	pm.title.Draw()
	pm.p1score.Draw()
	pm.hiscore.Draw()
	pm.p2score.Draw()
	pm.player.Draw()
	pm.player.shot.Draw()
	for i := range pm.alienShots {
		pm.alienShots[i].Draw()
	}
	pm.player.alienRack.drawRack()
}
