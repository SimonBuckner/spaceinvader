package main

import (
	"fmt"

	"github.com/SimonBuckner/spaceinvader/gfx"
)

func main() {

	vp, err := gfx.NewViewPort("Space Invaders", 50, 100, 1024, 768)
	if err != nil {
		fmt.Printf("error creating window: %v", err)
	}

	vp.Run()
}
