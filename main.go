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
	alienRows      = 5
	alienCols      = 11
)

type gameState struct {
	*gfx.Director
	vp              *gfx.ViewPort
	scale           float32
	backgroundColor sdl.Color
	// players         []*playerState
	// currentPlayer   *playerState
	ticks uint32
}

type playerState struct {
	lives  int
	score  int
	ship   *playerShip
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
		Director:        gfx.NewDirector(),
		vp:              vp,
		backgroundColor: sdl.Color{R: 0, G: 0, B: 0, A: 0},
		scale:           calcScale(vp),
		// players:         make([]*playerState, 2),
	}
	state.SetKeyboardEvent(state.keyb)
	state.SetUpdateEvent(state.update)

	newTestState(state)

	// for p := 0; p < 2; p++ {
	// 	ps := loadPlayerState(state, state.scale, p+1)
	// 	state.players = append(state.players, ps)
	// }
	// state.currentPlayer = state.players[0]
	// state.alphabet = resetAlphabet(vp, state.scale)
	fmt.Println("Finished loading assets")
	vp.Run(state.Director)
}

func loadPlayerState(gs *gameState, scale float32, number int) *playerState {
	fmt.Println("loading Player " + strconv.Itoa(number))

	ps := &playerState{
		lives: 3,
		score: 0,
	}

	ship, err := newPlayerShip(gs, number)
	if err != nil {
		fmt.Printf("error creating ship asset: %v", err)
		panic(err)
	}
	gs.vp.AddAsset(ship.Asset)
	ps.ship = ship
	return ps
}

func loadAlienGrid(gs *gameState) ([]*enemyShip, error) {
	fmt.Println("Loading Alien Grid")

	aliens := make([]*enemyShip, alienRows*alienCols)
	i := 0
	for row := 0; row < alienRows; row++ {
		for col := 0; col < alienCols; col++ {
			var class enemyClass
			switch row {
			case 0, 1:
				class = enemyClassC
			case 2, 3:
				class = enemyClassB
			case 4:
				class = enemyClassA
			}
			alien, err := newEnemyShip(gs, class)
			if err != nil {
				return nil, err
			}
			aliens[i] = alien
			gs.vp.AddAsset(alien.Asset)
			i++
		}
	}
	return aliens, nil
}

// func resetAlphabet(vp *gfx.ViewPort, scale float32) *gfx.AssetMap {
// 	fmt.Println("resetAlphabet")

// 	atlas, err := gfx.NewAssetMapFromBitMapAtlas(vp, alphabetAtlas)
// 	if err != nil {
// 		panic(err)
// 	}
// 	for _, asset := range atlas.GetAssets() {
// 		asset.SetScale(scale)
// 		asset.Show()
// 		vp.AddAsset(asset)
// 	}
// 	return atlas
// }

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
		// for _, player := range gs.players {
		// 	player.ship.update(ticks)
		// for _, alien := range player.aliens {
		// 	if alien == nil {
		// 		panic("nil alien")
		// 	}
		// 	if visible {
		// 		alien.Show()
		// 		if alien.CurrentIndex() >= 2 {
		// 			alien.SetCurrent(0)
		// 		} else {
		// 			alien.SetCurrent(alien.CurrentIndex() + 1)
		// 		}
		// 	} else {
		// 		alien.Hide()
		// 	}
		// }
		// ship := player.ship
		// if visible {
		// 	ship.Show()
		// 	// if ship.CurrentIndex() >= 2 {
		// 	// 	ship.SetCurrent(0)
		// 	// } else {
		// 	// 	ship.SetCurrent(ship.CurrentIndex() + 1)
		// 	// }
		// } else {
		// 	ship.Hide()
		// }
		// }
	}
}

func (gs *gameState) keyb(e *sdl.KeyboardEvent) {

	if e.Type == sdl.KEYUP {
		switch e.Keysym.Scancode {
		case sdl.SCANCODE_Q:
			gs.Close()
			return
		case sdl.SCANCODE_D:
			gs.dumpAssetNames()
		case sdl.SCANCODE_F1:
			gs.StartActor(testStateName)
		case sdl.SCANCODE_1:
			// gs.players[0].ship.Show()
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
