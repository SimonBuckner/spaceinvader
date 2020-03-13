package main

import (
	"github.com/SimonBuckner/spaceinvader/gfx"
	"github.com/veandco/go-sdl2/sdl"
)

type alienRack struct {
	*gfx.Actor
	game   *game
	props  []*gfx.Prop
	alienA []*sdl.Texture
	alienB []*sdl.Texture
	alienC []*sdl.Texture
}

func newAlienRack(game *game) *alienRack {
	ar := &alienRack{
		Actor:  gfx.NewActor("alien rack"),
		game:   game,
		props:  make([]*gfx.Prop, alienRows*alienCols),
		alienA: make([]*sdl.Texture, 2),
		alienB: make([]*sdl.Texture, 2),
		alienC: make([]*sdl.Texture, 2),
	}

	return ar
}

func (ar *alienRack) Start(scene *gfx.Scene) {
	ar.Scene = scene
	ar.Scale = scene.Scale()

	for _, p := range ar.props {
		p = gfx.NewProp("alient", nil, ar.game.transformXYZ)
		p.Scale = scene.Scale()
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

}

func (ar *alienRack) Draw() {
	for _, p := range ar.props {
		if p.Texture != nil {
			p.Draw(ar.Scene.Renderer())
		}
	}
}

func (ar *alienRack) reset() {
	// i := 0

	// var x, y int32
	// for r := int32(0); r < alienRows; r++ {
	// 	x = alienStartX
	// 	y = alienStartY - (r - alienRowHeight)
	// 	for c := int32(0); c <= alienCols; c++ {
	// 		x = x + (x * alienColWidth)
	// 		switch r {
	// 		case 0, 1:

	// 		case 2, 3:
	// 		case 4:
	// 		}
	// 	}
	// }
}
