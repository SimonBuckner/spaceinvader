package main

import (
	"github.com/SimonBuckner/spaceinvader/gfx"
	"github.com/veandco/go-sdl2/sdl"
)

type banner struct {
	*gfx.Actor
	text      string
	cacheText string
	texMap    map[string]*sdl.Texture
	props     []*gfx.Prop
}

func newBanner(game *game, cacheChars string, text string, maxLength int) *banner {
	t := &banner{
		Actor:     gfx.NewActor("text"),
		cacheText: cacheChars,
		text:      text,
		texMap:    make(map[string]*sdl.Texture),
		props:     make([]*gfx.Prop, maxLength),
	}
	return t
}

func (b *banner) Start(scene *gfx.Scene) {
	b.Scene = scene
	b.Scale = scene.Scale()
	stage := scene.Stage
	b.Visible = true
	for _, r := range b.cacheText {
		l := string(r)
		if _, ok := b.texMap[l]; ok {
			continue
		}
		tex, _ := alphabetAtlas.GetTexture(stage, l)
		b.texMap[l] = tex
	}
	for i := 0; i < len(b.props); i++ {
		b.props[i] = gfx.NewProp("banner", nil)
		b.props[i].Scale = b.Scale
	}
}

func (b *banner) Update(ticks uint32) {
	// fmt.Println("newText:onUpdate")
	if !b.Visible {
		return
	}

	x, y, _ := b.Pos.Int32()

	for i, r := range b.text {
		x1, y1 := convertXY(b.Scene, x, y)

		tex := b.texMap[string(r)]
		p := b.props[i]
		p.Pos.SetInt32(x1, y1, 0)
		p.Texture = b.texMap[string(r)]

		_, _, w, _, _ := tex.Query()
		x += w
	}
}

func (b *banner) Draw() {
	for _, p := range b.props {
		p.Draw(b.Scene.Renderer())
	}
}
