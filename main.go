package main

/*
	TODO: Alphabet props are not drawing

*/

import (
	"fmt"
	"runtime"

	"github.com/SimonBuckner/spaceinvader/gfx"
	"github.com/veandco/go-sdl2/sdl"
)

const (
	originalWidth  = 224
	originalHeight = 256
	width          = 1024
	height         = 768
	alienRows      = 5
	alienCols      = 11
)

type game struct {
	stage           *gfx.Stage
	backgroundColor sdl.Color
}

func main() {
	runtime.LockOSThread()
	scale := calcScale(width, height)
	stage, err := gfx.NewStage("Space Invaders", 50, 200, width, height, scale)
	if err != nil {
		fmt.Printf("error creating window: %v", err)
	}
	defer stage.Destroy()

	if _, err := newTestScene(stage); err != nil {
		panic(err)
	}

	game := &game{
		stage:           stage,
		backgroundColor: sdl.Color{R: 0, G: 0, B: 0, A: 0},
	}
	stage.SetKeyboardEvent(game.keyb)
	stage.Run()
}

func (g *game) keyb(e *sdl.KeyboardEvent) {
	fmt.Println("game.keyb")
	if e.Type == sdl.KEYUP {
		switch e.Keysym.Scancode {
		case sdl.SCANCODE_Q:
			g.stage.Close()
			return
		case sdl.SCANCODE_D:
			g.dumpPropNames()
		case sdl.SCANCODE_F1:
			g.stage.StartScene(testSceneName)
		case sdl.SCANCODE_1:
			// stage.players[0].ship.Show()
		}
	}
}

// func loadAlienGrid(g *gamee) ([]*enemyShip, error) {
// 	fmt.Println("Loading Alien Grid")

// 	aliens := make([]*enemyShip, alienRows*alienCols)
// 	i := 0
// 	for row := 0; row < alienRows; row++ {
// 		for col := 0; col < alienCols; col++ {
// 			var class enemyClass
// 			switch row {
// 			case 0, 1:
// 				class = enemyClassC
// 			case 2, 3:
// 				class = enemyClassB
// 			case 4:
// 				class = enemyClassA
// 			}
// 			alien, err := newEnemyShip(stage, class)
// 			if err != nil {
// 				return nil, err
// 			}
// 			aliens[i] = alien
// 			stage.stage.AddProp(alien.Prop)
// 			i++
// 		}
// 	}
// 	return aliens, nil
// }

// func resetAlphabet(stage *gfx.Stage, scale float32) *gfx.PropMap {
// 	fmt.Println("resetAlphabet")

// 	atlas, err := gfx.NewPropMapFromBitMapAtlas(stage, alphabetAtlas)
// 	if err != nil {
// 		panic(err)
// 	}
// 	for _, prop := range atlas.GetProps() {
// 		prop.SetScale(scale)
// 		prop.Show()
// 		stage.AddProp(prop)
// 	}
// 	return atlas
// }

func (g *game) update(ticks uint32) {

	g.stage.SetBackgroundColor(g.backgroundColor)

	// gridSize := 20 * stage.scale
	// x := gridSize
	// y := gridSize
	// i := 0
	// for _, prop := range stage.Props {
	// 	if prop.IsVisible() {
	// 		prop.SetPos(int32(x), int32(y), 0)
	// 		x = x + gridSize
	// 		i++
	// 		if i%10 == 0 {
	// 			x = gridSize
	// 			y = y + gridSize
	// 		}
	// 	}
	// }

	// if ticks-stage.ticks > 500 {
	// 	stage.ticks = ticks
	// for _, player := range stage.players {
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
	// }
}

func calcScale(w, h int32) float32 {

	rW := w / originalWidth
	rH := h / originalHeight
	if rW > rH {
		return float32(rH)
	}
	return float32(rW)
}

func (g *game) dumpPropNames() {
	fmt.Println("index  name                     x     y visible")
	fmt.Println("=====  ====================  ====  ==== =======")
	for i, v := range g.stage.Props() {
		x, y, _ := v.PosInt32()
		if v.Visible() {
			fmt.Printf(" %3d   %-20v  %4d  %4d Yes\n", i, v.Name, x, y)
		} else {
			fmt.Printf(" %3d   %-20v  %4d  %4d No\n", i, v.Name, x, y)
		}
	}
}

func (g *game) convertXY(x, y int32) (int32, int32) {
	w, h := g.stage.WindowSize()

	ow := float32(originalWidth) * g.stage.Scale()
	oh := float32(originalHeight) * g.stage.Scale()

	offsetX := (float32(w) - ow) / 2
	offsetY := (float32(h) - oh) / 2

	newX := int32(offsetX + (float32(x) * g.stage.Scale()))
	newY := int32(offsetY + (float32(y) * g.stage.Scale()))

	return newX, newY
}
