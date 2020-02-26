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

type state struct {
	running         bool
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
	aliens []*gfx.Asset
}

// IsRunning returns true if the game is running
func (s *state) IsRunning() bool {
	return s.running
}

func main() {
	runtime.LockOSThread()
	vp, err := gfx.NewViewPort("Space Invaders", 50, 200, 600, 768)
	if err != nil {
		fmt.Printf("error creating window: %v", err)
	}
	defer vp.Destroy()

	s := &state{
		running:         true,
		vp:              vp,
		backgroundColor: sdl.Color{R: 0, G: 0, B: 0, A: 0},
	}

	vp.KeyboardHandler = s.keyb
	vp.UpdateHandler = s.update
	s.scale = calcScale(vp)
	s.players = make([]*playerState, 2)
	for p := 0; p < 2; p++ {
		ps := resetPlayer(vp, s.scale)
		ps.ship.Name = "Player " + strconv.Itoa(p+1)
		ps.aliens = resetAlienGrid(vp, s.scale)
		s.players[p] = ps
	}
	s.currentPlayer = s.players[0]
	s.alphabet = resetAlphabet(vp, s.scale)
	fmt.Println("Finished loading assets")
	vp.Run(s)
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

func resetAlienGrid(vp *gfx.ViewPort, scale float32) []*gfx.Asset {
	fmt.Println("resetAlienGrid")
	rows := 5
	cols := 11
	aliens := make([]*gfx.Asset, rows*cols)
	i := 0
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			var alien *gfx.Asset
			switch row {
			case 0, 1:
				if a, err := gfx.AssetFromBitmaps(vp, alienSprC0, alienSprC1, alienExplode); err != nil {
					panic(err)
				} else {
					a.Name = "AlienC" + strconv.Itoa(i)
					alien = a
				}
			case 2, 3:
				if a, err := gfx.AssetFromBitmaps(vp, alienSprB0, alienSprB1, alienExplode); err != nil {
					panic(err)
				} else {
					a.Name = "AlienB" + strconv.Itoa(i)
					alien = a
				}
			case 4:
				if a, err := gfx.AssetFromBitmaps(vp, alienSprA0, alienSprA1, alienExplode); err != nil {
					panic(err)
				} else {
					a.Name = "AlienA" + strconv.Itoa(i)
					alien = a
				}
			}
			alien.SetScale(scale)
			aliens[i] = alien
			vp.AddAsset(alien)
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

func (s *state) update(vp *gfx.ViewPort, ticks uint32) {
	vp.SetBackgroundColor(s.backgroundColor)

	gridSize := 20 * s.scale
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

	if ticks-s.ticks > 500 {
		s.ticks = ticks
		for _, player := range s.players {
			visible := player == s.currentPlayer

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

func (s *state) keyb(e *sdl.KeyboardEvent) {

	if e.Type == sdl.KEYUP {
		switch e.Keysym.Scancode {
		case sdl.SCANCODE_Q:
			s.running = false
			return
		case sdl.SCANCODE_D:
			s.dumpAssetNames()
		}

	}
	if e.Type == sdl.KEYDOWN {
		switch e.Keysym.Scancode {
		case sdl.SCANCODE_R:
			if (e.Keysym.Mod & sdl.KMOD_SHIFT) == 0 {
				if s.backgroundColor.R < 254 {
					s.backgroundColor.R++
				}
			} else {
				if s.backgroundColor.R > 0 {
					s.backgroundColor.R--
				}
			}
		case sdl.SCANCODE_G:
			if (e.Keysym.Mod & sdl.KMOD_SHIFT) == 0 {
				if s.backgroundColor.G < 254 {
					s.backgroundColor.G++
				}
			} else {
				if s.backgroundColor.G > 0 {
					s.backgroundColor.G--
				}
			}
		case sdl.SCANCODE_B:
			if (e.Keysym.Mod & sdl.KMOD_SHIFT) == 0 {
				if s.backgroundColor.B < 254 {
					s.backgroundColor.B++
				}
			} else {
				if s.backgroundColor.B > 0 {
					s.backgroundColor.B--
				}
			}
		}
	}
}

func calcScale(vp *gfx.ViewPort) float32 {

	w, h := vp.WindowSize()

	rW := float32(w / originalWidth)
	rH := float32(h / originalHeight)

	if rW > rH {
		return rH
	}
	return rW
}

func (s *state) dumpAssetNames() {
	fmt.Println("index  name                     x     y visible")
	fmt.Println("=====  ====================  ====  ==== =======")
	for i, v := range s.vp.Assets {
		x, y, _ := v.Pos()
		if v.IsVisible() {
			fmt.Printf(" %3d   %-20v  %4d  %4d Yes\n", i, v.Name, x, y)
		} else {
			fmt.Printf(" %3d   %-20v  %4d  %4d No\n", i, v.Name, x, y)
		}
	}
}
