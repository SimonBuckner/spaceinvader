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
	player1         *gfx.Asset
	alien1          *gfx.Asset
}

// IsRunning returns true if the game is running
func (s *state) IsRunning() bool {
	return s.running
}

func main() {

	vp, err := gfx.NewViewPort("Space Invaders", 50, 100, 600, 768)
	if err != nil {
		fmt.Printf("error creating window: %v", err)
	}
	defer vp.Destroy()

	s := &state{
		running:         true,
		backgroundColor: sdl.Color{R: 0, G: 0, B: 0, A: 0},
		alien1:          nil,
	}

	vp.KeyboardHandler = s.keyb
	vp.UpdateHandler = s.update
	scale := calcScale(vp)

	s.player1 = gfx.AssetFromBitmap(vp, playerSprite, plrBlowupSprite0, plrBlowupSprite1)
	s.player1.SetPos(50, 50, 0)
	s.player1.SetScale(scale)
	vp.AddAsset(s.player1)

	s.alien1 = gfx.AssetFromBitmap(vp, alienSprCYA, alienSprCYB)
	s.alien1.SetPos(10, 50, 0)
	s.alien1.SetScale(scale)
	vp.AddAsset(s.alien1)

	vp.Run(s)
}

func (s *state) update(vp *gfx.ViewPort, ticks uint32) {
	if ticks-s.ticks > 500 {
		s.ticks = ticks

		if s.alien1.CurrentIndex() == 1 {
			s.alien1.SetCurrent(0)
		} else {
			s.alien1.SetCurrent(1)
		}
		if s.player1.CurrentIndex() >= 2 {
			s.player1.SetCurrent(0)
		} else {
			s.player1.SetCurrent(s.player1.CurrentIndex() + 1)
		}
	}
	vp.SetBackgroundColor(s.backgroundColor)
}

func (s *state) keyb(e *sdl.KeyboardEvent) {

	if e.Type == sdl.KEYUP {
		switch e.Keysym.Scancode {
		case sdl.SCANCODE_Q:
			s.running = false
		}
		if e.Keysym.Scancode == sdl.SCANCODE_R {
			if (e.Keysym.Mod & sdl.KMOD_SHIFT) == 0 {
				if s.backgroundColor.R < 254 {
					s.backgroundColor.R++
				}
			} else {
				if s.backgroundColor.R > 0 {
					s.backgroundColor.R--
				}
			}
		}
		if e.Keysym.Scancode == sdl.SCANCODE_G {
			if (e.Keysym.Mod & sdl.KMOD_SHIFT) == 0 {
				if s.backgroundColor.G < 254 {
					s.backgroundColor.G++
				}
			} else {
				if s.backgroundColor.G > 0 {
					s.backgroundColor.G--
				}
			}
		}
		if e.Keysym.Scancode == sdl.SCANCODE_B {
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
	} else {
		return rW
	}
}
