package main

/*
	TODO: Alphabet assets are not drawing

*/

import (
	"fmt"
	"runtime"
	"strconv"

	"github.com/SimonBuckner/spaceinvader/gfx"
	"github.com/veandco/go-sdl2/sdl"
)

const (
	originalWidth  = 224
	originalHeight = 256
)

type gameState struct {
	*gfx.StateControl
	ticks           uint32
	scale           float32
	backgroundColor sdl.Color
	vp              *gfx.ViewPort
	players         []*playerState
	currentPlayer   *playerState
	alphabet        *gfx.AssetMap
}

type playerState struct {
	lives  int
	score  int
	ship   *gfx.Asset
	aliens []*enemyShip
}

func main() {
	runtime.LockOSThread()
	vp, err := gfx.NewViewPort("Space Invaders", 50, 200, 600, 768)
	if err != nil {
		fmt.Printf("error creating window: %v", err)
	}
	defer vp.Destroy()

	state := &gameState{
		StateControl:    gfx.NewStateControl(),
		vp:              vp,
		backgroundColor: sdl.Color{R: 0, G: 0, B: 0, A: 0},
	}

	state.SetKeyboardEvent(state.keyb)
	state.SetUpdateEvent(state.update)
	state.scale = calcScale(vp)
	state.players = make([]*playerState, 2)
	for p := 0; p < 2; p++ {
		ps := resetPlayer(vp, state.scale)
		ps.ship.Name = "Player " + strconv.Itoa(p+1)
		ps.aliens = resetAlienGrid(state)
		state.players[p] = ps
	}
	state.currentPlayer = state.players[0]
	state.alphabet = resetAlphabet(vp, state.scale)
	fmt.Println("Finished loading assets")
	vp.Run(state)
}

func resetPlayer(vp *gfx.ViewPort, scale float32) *playerState {
	fmt.Println("resetPlayer")

	ps := &playerState{
		lives: 3,
		score: 0,
	}

	ship, err := gfx.AssetFromBitmaps(vp, playerSprite, plrBlowupSprite0, plrBlowupSprite1)
	if err != nil {
		fmt.Printf("error creating ship asset: %v", err)
		panic(err)
	}
	ship.SetScale(scale)
	vp.AddAsset(ship)
	ps.ship = ship
	return ps
}

func resetAlienGrid(gs *gameState) []*enemyShip {
	fmt.Println("resetAlienGrid")
	rows := 5
	cols := 11
	aliens := make([]*enemyShip, rows*cols)
	i := 0
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			var class enemyClass
			switch row {
			case 0, 1:
				class = enemyClassC
			case 2, 3:
				class = enemyClassB
			case 4:
				class = enemyClassA
			}
			a, err := newAlien(gs, class)
			if err != nil {

			}
			a.SetScale(gs.scale)
			aliens[i] = a
			gs.vp.AddAsset(a.Asset)
			i++
		}
	}
	return aliens
}

func resetAlphabet(vp *gfx.ViewPort, scale float32) *gfx.AssetMap {
	fmt.Println("resetAlphabet")

	atlas, err := gfx.NewAssetMapFromBitMapAtlas(vp, alphabetAtlas)
	if err != nil {
		panic(err)
	}
	for _, asset := range atlas.GetAssets() {
		asset.SetScale(scale)
		asset.Show()
		vp.AddAsset(asset)
	}
	return atlas
}

func (gs *gameState) update(ticks uint32) {
	vp := gs.vp
	vp.SetBackgroundColor(gs.backgroundColor)

	gridSize := 20 * gs.scale
	x := gridSize
	y := gridSize
	i := 0
	for _, asset := range vp.Assets {
		if asset.IsVisible() {
			asset.SetPos(int32(x), int32(y), 0)
			x = x + gridSize
			i++
			if i%10 == 0 {
				x = gridSize
				y = y + gridSize
			}
		}
	}

	if ticks-gs.ticks > 500 {
		gs.ticks = ticks
		for _, player := range gs.players {
			visible := player == gs.currentPlayer

			for _, alien := range player.aliens {
				if alien == nil {
					panic("nil alien")
				}
				if visible {
					alien.Show()
					if alien.CurrentIndex() >= 2 {
						alien.SetCurrent(0)
					} else {
						alien.SetCurrent(alien.CurrentIndex() + 1)
					}
				} else {
					alien.Hide()
				}
			}
			ship := player.ship
			if visible {
				ship.Show()
				if ship.CurrentIndex() >= 2 {
					ship.SetCurrent(0)
				} else {
					ship.SetCurrent(ship.CurrentIndex() + 1)
				}
			} else {
				ship.Hide()
			}
		}
	}

}

func (gs *gameState) keyb(e *sdl.KeyboardEvent) {

	if e.Type == sdl.KEYUP {
		switch e.Keysym.Scancode {
		case sdl.SCANCODE_Q:
			gs.Quit()
			return
		case sdl.SCANCODE_D:
			gs.dumpAssetNames()
		}

	}
	if e.Type == sdl.KEYDOWN {
		switch e.Keysym.Scancode {
		case sdl.SCANCODE_R:
			if (e.Keysym.Mod & sdl.KMOD_SHIFT) == 0 {
				if gs.backgroundColor.R < 254 {
					gs.backgroundColor.R++
				}
			} else {
				if gs.backgroundColor.R > 0 {
					gs.backgroundColor.R--
				}
			}
		case sdl.SCANCODE_G:
			if (e.Keysym.Mod & sdl.KMOD_SHIFT) == 0 {
				if gs.backgroundColor.G < 254 {
					gs.backgroundColor.G++
				}
			} else {
				if gs.backgroundColor.G > 0 {
					gs.backgroundColor.G--
				}
			}
		case sdl.SCANCODE_B:
			if (e.Keysym.Mod & sdl.KMOD_SHIFT) == 0 {
				if gs.backgroundColor.B < 254 {
					gs.backgroundColor.B++
				}
			} else {
				if gs.backgroundColor.B > 0 {
					gs.backgroundColor.B--
				}
			}
		}
	}
}

func calcScale(vp *gfx.ViewPort) float32 {

	w, h := vp.WindowSize()

	rW := int(w / originalWidth)
	rH := int(h / originalHeight)

	if rW > rH {
		return float32(rH)
	}
	return float32(rW)
}

func (gs *gameState) dumpAssetNames() {
	fmt.Println("index  name                     x     y visible")
	fmt.Println("=====  ====================  ====  ==== =======")
	for i, v := range gs.vp.Assets {
		x, y, _ := v.Pos()
		if v.IsVisible() {
			fmt.Printf(" %3d   %-20v  %4d  %4d Yes\n", i, v.Name, x, y)
		} else {
			fmt.Printf(" %3d   %-20v  %4d  %4d No\n", i, v.Name, x, y)
		}
	}
}

func (gs *gameState) convertXY(x, y int32) (int32, int32) {
	w, h := gs.vp.WindowSize()

	newX := (w - int32(originalWidth*gs.scale)) / 2
	newY := (h - int32(originalHeight*gs.scale)) / 2
	return newX, newY
}
