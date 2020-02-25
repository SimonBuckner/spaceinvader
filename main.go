package main

import (
	"fmt"

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
	backgroundColor sdl.Color
	players         []*playerState
	currentPlayer   *playerState
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

	vp, err := gfx.NewViewPort("Space Invaders", 50, 200, 600, 768)
	if err != nil {
		fmt.Printf("error creating window: %v", err)
	}
	defer vp.Destroy()

	s := &state{
		running:         true,
		backgroundColor: sdl.Color{R: 0, G: 0, B: 0, A: 0},
	}

	vp.KeyboardHandler = s.keyb
	vp.UpdateHandler = s.update
	scale := calcScale(vp)

	s.players = make([]*playerState, 2)
	for p := 0; p < 2; p++ {
		ps := &playerState{
			lives:  3,
			score:  0,
			ship:   resetShip(vp, scale),
			aliens: resetAlienGrid(vp, scale),
		}
		s.players[p] = ps
	}
	s.currentPlayer = s.players[0]
	vp.Run(s)
}

func resetShip(vp *gfx.ViewPort, scale float32) *gfx.Asset {
	ship := gfx.AssetFromBitmap(vp, playerSprite, plrBlowupSprite0, plrBlowupSprite1)
	ship.SetPos(50, 700, 0)
	ship.SetScale(scale)
	vp.AddAsset(ship)
	return ship
}

func resetAlienGrid(vp *gfx.ViewPort, scale float32) []*gfx.Asset {
	aliens := make([]*gfx.Asset, 55)
	i := 0
	for row := 0; row < 5; row++ {
		for col := 0; col < 11; col++ {
			x := int32(50 + (float32(20*col) * scale))
			y := int32(50 + (float32(20*row) * scale))
			var alien *gfx.Asset
			switch row {
			case 0, 1:
				alien = gfx.AssetFromBitmap(vp, alienSprC0, alienSprC1, alienExplode)
			case 2, 3:
				alien = gfx.AssetFromBitmap(vp, alienSprB0, alienSprB1, alienExplode)
			case 4:
				alien = gfx.AssetFromBitmap(vp, alienSprA0, alienSprA1, alienExplode)
			}
			alien.SetPos(x, y, 0)
			alien.SetScale(scale)
			aliens[i] = alien
			vp.AddAsset(alien)
			i++
		}
	}
	return aliens
}

func (s *state) update(vp *gfx.ViewPort, ticks uint32) {
	vp.SetBackgroundColor(s.backgroundColor)
	if ticks-s.ticks > 500 {
		s.ticks = ticks
		for _, player := range s.players {
			visible := player == s.currentPlayer

			for _, alien := range player.aliens {
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
