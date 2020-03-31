package main

import (
	"flag"
	"fmt"
	"runtime"

	"github.com/SimonBuckner/screen2d"
	"github.com/veandco/go-sdl2/sdl"
)

const (
	originalWidth  = 224
	originalHeight = 256
	width          = 1024
	height         = 768
	alienRows      = 5
	alienCols      = 11
	alienRowHeight = 16
	alienColWidth  = 16
	alienStartX    = 10
	alienStartY    = originalHeight - 0x78
)

const (
	keyAlienSprCYA screen2d.SpriteMapKey = iota
	keyAlienSprCYB
	keyAlienSprA0
	keyAlienSprA1
	keyAlienSprB0
	keyAlienSprB1
	keyAlienSprC0
	keyAlienSprC1
	keyPlayerSprite
	keyPlrBlowupSprite0
	keyPlrBlowupSprite1
	keyPlayerShotSpr
	keyShotExploding
	keyAlienExplode
	keySquiglyShot0
	keySquiglyShot2
	keySquiglyShot3
	keyPlungerShot0
	keyPlungerShot1
	keyPlungerShot2
	keyPlungerShot3
	keyRollShot0
	keyRollShot1
	keyRollShot2
	keyRollShot3
	keyShieldImage
	keySpriteSaucer
	keySpriteSaucerExp
	keyAlienSprCA
	keyAlienSprCB
)

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")

type game struct {
	screen          *screen2d.Screen
	scale           float32
	backgroundColor sdl.Color
	sprites         *screen2d.SpriteMap
}

func main() {
	runtime.LockOSThread()

	g := &game{
		backgroundColor: sdl.Color{R: 0, G: 0, B: 0, A: 0},
	}

	g.scale = calcScale(width, height)

	{
		screen, err := screen2d.NewScreen(width, height, "Space Invaders")
		if err != nil {
			panic(err)
		}
		screen.SetKeyDownFunc(g.onKeyDown)
		g.screen = screen
	}
	defer g.screen.Destroy()

	g.loadSpriteMap()
	g.screen.Run()
}

func (g *game) loadSpriteMap() {
	g.sprites = screen2d.NewSpriteMap()

	g.loadSprite(keyAlienSprCYA, alienSprCYA)
	g.loadSprite(keyAlienSprCYB, alienSprCYB)
	g.loadSprite(keyAlienSprA0, alienSprA0)
	g.loadSprite(keyAlienSprA1, alienSprA1)
	g.loadSprite(keyAlienSprB0, alienSprB0)
	g.loadSprite(keyAlienSprB1, alienSprB1)
	g.loadSprite(keyAlienSprC0, alienSprC0)
	g.loadSprite(keyAlienSprC1, alienSprC1)
	g.loadSprite(keyPlayerSprite, playerSprite)
	g.loadSprite(keyPlrBlowupSprite0, plrBlowupSprite0)
	g.loadSprite(keyPlrBlowupSprite1, plrBlowupSprite1)
	g.loadSprite(keyPlayerShotSpr, playerShotSpr)
	g.loadSprite(keyShotExploding, shotExploding)
	g.loadSprite(keyAlienExplode, alienExplode)
	g.loadSprite(keySquiglyShot0, squiglyShot0)
	g.loadSprite(keySquiglyShot2, squiglyShot2)
	g.loadSprite(keySquiglyShot3, squiglyShot3)
	g.loadSprite(keyPlungerShot0, plungerShot0)
	g.loadSprite(keyPlungerShot1, plungerShot1)
	g.loadSprite(keyPlungerShot2, plungerShot2)
	g.loadSprite(keyPlungerShot3, plungerShot3)
	g.loadSprite(keyRollShot0, rollShot0)
	g.loadSprite(keyRollShot1, rollShot1)
	g.loadSprite(keyRollShot2, rollShot2)
	g.loadSprite(keyRollShot3, rollShot3)
	g.loadSprite(keyShieldImage, shieldImage)
	g.loadSprite(keySpriteSaucer, spriteSaucer)
	g.loadSprite(keySpriteSaucerExp, spriteSaucerExp)
	g.loadSprite(keyAlienSprCA, alienSprCA)
	g.loadSprite(keyAlienSprCB, alienSprCB)
}

func (g *game) loadSprite(key screen2d.SpriteMapKey, bm *Bitmap) {
	s := screen2d.NewSprite(g.screen.Rend())
	err := s.LoadRGBAPixels(bm.Pixels, bm.Pitch)
	if err != nil {
		panic(err)
	}
	g.sprites.AddSprite(key, s)
}

func (g *game) onKeyDown(e *sdl.KeyboardEvent) {
	switch e.Keysym.Scancode {
	case sdl.SCANCODE_Q:
		g.screen.Close()
		return
	case sdl.SCANCODE_D:
		// g.stage.DumpActors()
	case sdl.SCANCODE_F1:
		// if g.stage.Scene == g.testScreen.Scene {
		// 	g.stage.StopScene()
		// } else {
		// 	g.stage.StartScene(g.testScreen.Scene)
		// }
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

func (g *game) transformXYZ(e *screen2d.Entity) (int32, int32, int32) {

	ow := float32(originalWidth) * g.scale
	oh := float32(originalHeight) * g.scale

	offsetX := (float32(width) - ow) / 2
	offsetY := (float32(height) - oh) / 2

	newX := float32(e.X) * g.scale
	newY := float32(e.Y) * g.scale

	return int32(newX + offsetX), int32(newY + offsetY), 0
}

func (g *game) transformXYZDebug(e *screen2d.Entity) (int32, int32, int32) {
	x, y, _ := g.transformXYZ(e)

	fmt.Printf("X1: %04d Y041: %d  -  X2: %04d Y2: %04d\n", int32(e.X), int32(e.Y), int32(x), int32(y))

	return x, y, 0
}
