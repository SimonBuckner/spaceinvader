package main

import (
	"fmt"
	"runtime"

	"github.com/SimonBuckner/screen2d"
	"github.com/veandco/go-sdl2/sdl"
)

type game struct {
	screen          *screen2d.Screen
	keyb            *screen2d.KBState
	scale           float32
	backgroundColor sdl.Color
	sprites         *screen2d.SpriteMap
	font            *screen2d.SpriteAtlas
	pm              *playMode
}

func main() {
	runtime.LockOSThread()

	g := &game{
		backgroundColor: sdl.Color{R: 0, G: 0, B: 0, A: 0},
	}

	g.scale = calcScale(winWidth, winHeight)

	{
		screen, err := screen2d.NewScreen(winWidth, winHeight, "Space Invaders",
			screen2d.SetVSync(true),
			screen2d.SetScalingQuality(screen2d.ScreenScalingNearestPixel))
		if err != nil {
			panic(err)
		}
		g.screen = screen
	}
	defer g.screen.Destroy()

	g.keyb = g.screen.GetKBState()
	g.loadSpriteMap()
	g.pm = newPlayMode(g)
	g.activate()

	g.screen.Run()
}

func (g *game) activate() {
	g.screen.ClearFuncs()
	g.screen.SetKeyDownFunc(g.onKeyDown)
	g.screen.SetUpdateFunc(g.onUpdate)
}

func (g *game) loadSpriteMap() {
	g.sprites = screen2d.NewSpriteMap()

	g.loadSprite(keyAlienExplode, alienExplode)

	g.loadSprite(keyAlienSprA0, alienSprA0)
	g.loadSprite(keyAlienSprA1, alienSprA1)

	g.loadSprite(keyAlienSprB0, alienSprB0)
	g.loadSprite(keyAlienSprB1, alienSprB1)

	g.loadSprite(keyAlienSprC0, alienSprC0)
	g.loadSprite(keyAlienSprC1, alienSprC1)

	g.loadSprite(keyAlienSprCA, alienSprCA)
	g.loadSprite(keyAlienSprCB, alienSprCB)

	g.loadSprite(keyAlienSprCYA, alienSprCYA)
	g.loadSprite(keyAlienSprCYB, alienSprCYB)

	g.loadSprite(keyPlayerShotSpr, playerShotSpr)
	g.loadSprite(keyPlayerSprite, playerSprite)

	g.loadSprite(keyPlrBlowupSprite0, plrBlowupSprite0)
	g.loadSprite(keyPlrBlowupSprite1, plrBlowupSprite1)

	g.loadSprite(keyPlungerShot0, plungerShot0)
	g.loadSprite(keyPlungerShot1, plungerShot1)
	g.loadSprite(keyPlungerShot2, plungerShot2)
	g.loadSprite(keyPlungerShot3, plungerShot3)

	g.loadSprite(keyRollShot0, rollShot0)
	g.loadSprite(keyRollShot1, rollShot1)
	g.loadSprite(keyRollShot2, rollShot2)
	g.loadSprite(keyRollShot3, rollShot3)

	g.loadSprite(keyShieldImage, shieldImage)

	g.loadSprite(keyShotExploding, shotExploding)

	g.loadSprite(keySpriteSaucer, spriteSaucer)
	g.loadSprite(keySpriteSaucerExp, spriteSaucerExp)

	g.loadSprite(keySquiglyShot0, squiglyShot0)
	g.loadSprite(keySquiglyShot2, squiglyShot2)
	g.loadSprite(keySquiglyShot3, squiglyShot3)

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
	case sdl.SCANCODE_Q, sdl.SCANCODE_ESCAPE:
		g.screen.Close()
		return
	case sdl.SCANCODE_D:
		// g.stage.DumpActors()
	case sdl.SCANCODE_F1:
		g.pm.activate()
	}
}

func (g *game) onUpdate(ticks uint32, elapsed float32) {

	if g.keyb.IsKeyDown(sdl.SCANCODE_Q) {
		fmt.Println("Game key down")
		g.screen.Close()
	} else if g.keyb.IsKeyDown(sdl.SCANCODE_F1) {
		fmt.Println("Game key down")
		g.pm.activate()
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

func translatePos(x, y float32, scale float32) (tX, tY int32) {

	scaledW := float32(originalWidth) * scale
	scaledH := float32(originalHeight) * scale

	offsetX := (float32(winWidth) - scaledW) / 2
	offsetY := (float32(winHeight) - scaledH) / 2

	scaledX := x * scale
	scaledY := y * scale

	tX = int32(scaledX + offsetX)
	tY = int32(scaledY + offsetY)

	return
}

func translatePosDebug(x, y float32, scale float32) (int32, int32) {
	tX, tY := translatePos(x, y, scale)

	fmt.Printf("X1: %f Y1: %f  -  X2: %04d Y2: %04d\n", x, y, tX, tY)

	return tX, tY
}
