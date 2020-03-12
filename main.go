package main

import (
	"flag"
	"log"
	"os"
	"runtime"
	"runtime/pprof"

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

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")

type game struct {
	stage           *gfx.Stage
	backgroundColor sdl.Color
	testScreen      *testScreen
}

func main() {
	runtime.LockOSThread()

	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	game := &game{
		backgroundColor: sdl.Color{R: 0, G: 0, B: 0, A: 0},
	}
	game.testScreen = newTestScene(game)

	scale := calcScale(width, height)
	stage, err := gfx.NewStage("Space Invaders", 26, 400, width, height, scale)
	if err != nil {
		panic(err)
	}
	defer stage.Destroy()
	game.stage = stage

	stage.KeyboardEventHandler = game.onKeyboard

	stage.Start()
}

func (g *game) onKeyboard(e *sdl.KeyboardEvent) {
	if e.Type == sdl.KEYUP {
		switch e.Keysym.Scancode {
		case sdl.SCANCODE_Q:
			g.stage.Stop()
			return
		case sdl.SCANCODE_D:
			g.stage.DumpActors()
		case sdl.SCANCODE_F1:
			if g.stage.Scene == g.testScreen.Scene {
				g.stage.StopScene()
			} else {
				g.stage.StartScene(g.testScreen.Scene)
			}
		}
	}
}

func calcScale(w, h int32) float32 {

	rW := w / originalWidth
	rH := h / originalHeight
	if rW > rH {
		return float32(rH)
	}
	return float32(rW)
}

func convertXY(scene *gfx.Scene, x, y int32) (int32, int32) {
	w, h := scene.Stage.WindowSize()

	ow := float32(originalWidth) * scene.Stage.Scale
	oh := float32(originalHeight) * scene.Stage.Scale

	offsetX := (float32(w) - ow) / 2
	offsetY := (float32(h) - oh) / 2

	newX := float32(x) * scene.Stage.Scale
	newY := float32(y) * scene.Stage.Scale

	return int32(newX + offsetX), int32(newY + offsetY)
}
