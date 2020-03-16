package main

import (
	"github.com/SimonBuckner/spaceinvader/gfx"
	"github.com/veandco/go-sdl2/sdl"
)

type alienRack struct {
	*gfx.Actor
	game         *game
	props        []*gfx.Prop
	alienA       []*sdl.Texture
	alienB       []*sdl.Texture
	alienC       []*sdl.Texture
	startTick    uint32
	refreshCount int
	texIndex     int
	remaining    int
}

func newAlienRack(game *game) *alienRack {
	ar := &alienRack{
		Actor:     gfx.NewActor("alien rack"),
		game:      game,
		props:     make([]*gfx.Prop, alienRows*alienCols),
		alienA:    make([]*sdl.Texture, 2),
		alienB:    make([]*sdl.Texture, 2),
		alienC:    make([]*sdl.Texture, 2),
		remaining: alienRows * alienCols,
	}
	return ar
}

func (ar *alienRack) Start(scene *gfx.Scene) {
	ar.Scene = scene
	ar.Scale = scene.Scale()

	for i := 0; i < len(ar.props); i++ {
		ar.props[i] = gfx.NewProp("alien", nil, ar.game.transformXYZ)
		ar.props[i].Scale = scene.Scale()
	}
	ar.alienA[0], _ = alienSprA0.ToTexture(scene.Stage)
	ar.alienA[1], _ = alienSprA1.ToTexture(scene.Stage)
	ar.alienB[0], _ = alienSprB0.ToTexture(scene.Stage)
	ar.alienB[1], _ = alienSprB1.ToTexture(scene.Stage)
	ar.alienC[0], _ = alienSprC0.ToTexture(scene.Stage)
	ar.alienC[1], _ = alienSprC1.ToTexture(scene.Stage)

	ar.reset()
}

func (ar *alienRack) Update(ticks uint32) {
	if ticks-ar.startTick > 16 {
		ar.startTick = ticks
		ar.refreshCount++
		if ar.refreshCount >= ar.remaining {
			ar.refreshCount = 0
			ar.texIndex = 0 + (1 - ar.texIndex)
		}
	}

	i := 0
	x, y, _ := ar.Pos.Int32()

	for r := int32(0); r < alienRows; r++ {
		for c := int32(0); c < alienCols; c++ {
			switch r {
			case 0, 1:
				ar.props[i].Texture = ar.alienA[ar.texIndex]
			case 2, 3:
				ar.props[i].Texture = ar.alienB[ar.texIndex]
			case 4:
				ar.props[i].Texture = ar.alienC[ar.texIndex]
			}
			ar.props[i].Pos.SetInt32(x, y, 0)
			x = x + alienColWidth
			i++
		}
		x = alienStartX
		y = y - alienRowHeight
	}
}

func (ar *alienRack) Draw() {
	for _, p := range ar.props {
		if p != nil {
			p.Draw(ar.Scene.Renderer())
		}
	}
}

func (ar *alienRack) reset() {
	ar.startTick = 0
	ar.refreshCount = 0
	ar.texIndex = 0
	ar.remaining = alienRows * alienCols

	ar.Pos.SetInt32(alienStartX, alienStartY, 0)
	x, y, _ := ar.Pos.Int32()

	i := 0
	for r := int32(0); r < alienRows; r++ {
		for c := int32(0); c < alienCols; c++ {
			switch r {
			case 0, 1:
				ar.props[i].Texture = ar.alienA[0]
			case 2, 3:
				ar.props[i].Texture = ar.alienB[0]
			case 4:
				ar.props[i].Texture = ar.alienC[0]
			}
			ar.props[i].Pos.SetInt32(x, y, 0)
			x = x + alienColWidth
			i++
		}
		x = alienStartX
		y = y - alienRowHeight
	}
}

func (ar *alienRack) checkForHit(player *player) bool {
	sx, sy, _ := player.shot.Pos.Int32()
	_, _, sw, sh, _ := player.shot.Texture.Query()
	sRect := sdl.Rect{X: sx, Y: sy, W: sw, H: sh}

	for _, a := range ar.props {
		ax, ay, _ := a.Pos.Int32()
		_, _, aw, ah, _ := a.Texture.Query()
		aRect := sdl.Rect{X: ax, Y: ay, W: aw, H: ah}
		if sRect.HasIntersection(&aRect) {
			return true
		}
	}
	return false
}
