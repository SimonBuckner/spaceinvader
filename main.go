package main

import (
	"fmt"

	"github.com/SimonBuckner/spaceinvader/gfx"
	"github.com/veandco/go-sdl2/sdl"
)

func main() {

	vp, err := gfx.NewViewPort("Space Invaders", 50, 100, 1024, 768)
	if err != nil {
		fmt.Printf("error creating window: %v", err)
	}
	defer vp.Destroy()

	vp.KeyboardHandler = keyb
	vp.Run()
}

func keyb(vp *gfx.ViewPort, e *sdl.KeyboardEvent) {

	if e.Type == sdl.KEYUP {
		switch e.Keysym.Scancode {
		case sdl.SCANCODE_Q:
			fmt.Println("Q")
			vp.Exit()
		}
	}
	// if e.Type == sdl.KEYDOWN {
	// 	if e.Keysym.Scancode == sdl.SCANCODE_R {
	// 		if (e.Keysym.Mod & sdl.KMOD_SHIFT) == 0 {
	// 			if r < 254 {
	// 				r++
	// 			}
	// 		} else {
	// 			if r > 0 {
	// 				r--
	// 			}
	// 		}
	// 	}
	// 	if e.Keysym.Scancode == sdl.SCANCODE_G {
	// 		if (e.Keysym.Mod & sdl.KMOD_SHIFT) == 0 {
	// 			if g < 254 {
	// 				g++
	// 			}
	// 		} else {
	// 			if g > 0 {
	// 				g--
	// 			}
	// 		}
	// 	}
	// 	if e.Keysym.Scancode == sdl.SCANCODE_B {
	// 		if (e.Keysym.Mod & sdl.KMOD_SHIFT) == 0 {
	// 			if b < 254 {
	// 				b++
	// 			}
	// 		} else {
	// 			if b > 0 {
	// 				b--
	// 			}
	// 		}
	// 	}
	// }
}
