package main

import (
	"fmt"

	"github.com/SimonBuckner/spaceinvader/gfx"
	"github.com/veandco/go-sdl2/sdl"
)

type state struct {
	running         bool
	backgroundColor sdl.Color
}

// IsRunning returns true if the game is running
func (s *state) IsRunning() bool {
	return s.running
}

func (s *state) keyb(e *sdl.KeyboardEvent) {

	if e.Type == sdl.KEYUP {
		switch e.Keysym.Scancode {
		case sdl.SCANCODE_Q:
			fmt.Println("Q")
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

func (s *state) update(vp *gfx.ViewPort) {
	vp.SetBackgroundColor(s.backgroundColor)
}

func main() {

	vp, err := gfx.NewViewPort("Space Invaders", 50, 100, 1024, 768)
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

	vp.Run(s)
}
