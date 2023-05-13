package main

import (
	"fmt"
	"runtime"

	"github.com/SimonBuckner/screen2d"
	"github.com/veandco/go-sdl2/sdl"
)

type game struct {
	es              *screen2d.EntityService
	screen          *screen2d.Screen
	keyb            *screen2d.KBState
	scale           float32
	backgroundColor sdl.Color
	sprites         *screen2d.SpriteMap
	font            *screen2d.SpriteAtlas
	fontKeys        map[rune]int32
	pm              *playMode
}

func main() {
	runtime.LockOSThread()

	g := &game{
		backgroundColor: sdl.Color{R: 0, G: 0, B: 0, A: 0},
	}

	screen, err := screen2d.NewScreen(winWidth, winHeight, "Space Invaders",
		screen2d.SetVSync(true),
		screen2d.SetScalingQuality(screen2d.ScreenScalingNearestPixel))
	if err != nil {
		panic(err)
	}
	g.screen = screen
	defer g.screen.Destroy()

	es := screen2d.NewEntityService()
	es.SetXYProjection(projectXY)
	es.SetScale(calcScale(winWidth, winHeight))
	g.es = es

	g.keyb = g.screen.GetKBState()
	g.loadSpriteMap()
	g.loadFontAtlas()
	g.pm = newPlayMode(g)
	g.activate()

	g.screen.Run()
}

func (g *game) activate() {
	g.screen.ClearFuncs()
	g.screen.SetKeyDownFunc(g.onKeyDown)
	g.screen.SetUpdateFunc(g.onUpdate)
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
	case sdl.SCANCODE_F1:
		g.pm.activate()
	}
}

func (g *game) onUpdate(ticks uint32, elapsed float32) {

	if g.keyb.IsKeyDown(uint8(sdl.K_q)) {
		g.screen.Close()
	} else if g.keyb.IsKeyDown(uint8(sdl.SCANCODE_F1)) {
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

func projectXY(x, y float32, scale float32) (int32, int32) {

	scaledW := float32(originalWidth) * scale
	scaledH := float32(originalHeight) * scale

	offsetX := (float32(winWidth) - scaledW) / 2
	offsetY := (float32(winHeight) - scaledH) / 2

	scaledX := x * scale
	scaledY := y * scale

	tX := int32(scaledX + offsetX)
	tY := int32(scaledY + offsetY)

	return tX, tY
}

func projectXYDebug(x, y float32, scale float32) (int32, int32) {
	tX, tY := projectXY(x, y, scale)

	fmt.Printf("X1: %f Y1: %f  -  X2: %04d Y2: %04d\n", x, y, tX, tY)

	return tX, tY
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
	g.loadSprite(keyAlienShotExploding, aShotExplo)

	g.loadSprite(keySpriteSaucer, spriteSaucer)
	g.loadSprite(keySpriteSaucerExp, spriteSaucerExp)

	g.loadSprite(keySquiglyShot0, squiglyShot0)
	g.loadSprite(keySquiglyShot1, squiglyShot1)
	g.loadSprite(keySquiglyShot2, squiglyShot2)
	g.loadSprite(keySquiglyShot3, squiglyShot3)

}

func (g *game) loadFontAtlas() {
	g.font = screen2d.NewSpriteAtlas(g.screen.Rend())
	g.fontKeys = make(map[rune]int32)

	err := g.font.LoadRGBAPixels(alphabet.Pixels, alphabet.Pitch, 7, 8)
	if err != nil {
		panic(err)
	}

	g.fontKeys['A'] = 0
	g.fontKeys['B'] = 1
	g.fontKeys['C'] = 2
	g.fontKeys['D'] = 3
	g.fontKeys['E'] = 4
	g.fontKeys['F'] = 5
	g.fontKeys['G'] = 6
	g.fontKeys['H'] = 7
	g.fontKeys['I'] = 8
	g.fontKeys['J'] = 9
	g.fontKeys['K'] = 10
	g.fontKeys['L'] = 11
	g.fontKeys['M'] = 12
	g.fontKeys['N'] = 13
	g.fontKeys['O'] = 14
	g.fontKeys['P'] = 15
	g.fontKeys['Q'] = 16
	g.fontKeys['R'] = 17
	g.fontKeys['S'] = 18
	g.fontKeys['T'] = 19
	g.fontKeys['U'] = 20
	g.fontKeys['V'] = 21
	g.fontKeys['W'] = 22
	g.fontKeys['X'] = 23
	g.fontKeys['Y'] = 24
	g.fontKeys['Z'] = 25
	g.fontKeys['0'] = 26
	g.fontKeys['1'] = 27
	g.fontKeys['2'] = 28
	g.fontKeys['3'] = 29
	g.fontKeys['4'] = 30
	g.fontKeys['5'] = 31
	g.fontKeys['6'] = 32
	g.fontKeys['7'] = 33
	g.fontKeys['8'] = 34
	g.fontKeys['9'] = 35
	g.fontKeys['<'] = 36
	g.fontKeys['>'] = 37
	g.fontKeys[' '] = 38
	g.fontKeys['='] = 39
	g.fontKeys['*'] = 40
	g.fontKeys['y'] = 41
	g.fontKeys['-'] = 42
}
