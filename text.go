package main

import (
	"github.com/SimonBuckner/spaceinvader/gfx"
	"github.com/veandco/go-sdl2/sdl"
)

type text struct {
	*gfx.Actor
	value      string
	cacheChars string
	texMap     map[string]*sdl.Texture
	display    []*gfx.Prop
	maxLength  int
}

func newText(game *game, cacheChars string, value string, maxLength int) *text {
	// fmt.Println("newText")
	t := &text{
		Actor:      gfx.NewActor("text"),
		cacheChars: cacheChars,
		value:      value,
		texMap:     make(map[string]*sdl.Texture),
		maxLength:  maxLength,
		display:    make([]*gfx.Prop, maxLength),
	}
	t.StartEventHandler = t.onStart
	t.StopEventHandler = t.onStop
	t.UpdateEventHandler = t.onUpdate
	return t
}

func (t *text) onStart() {
	// fmt.Println("newText:onStart")
	stage := t.Scene.Stage
	t.Visible = true
	for _, r := range t.cacheChars {
		l := string(r)
		if _, ok := t.texMap[l]; ok {
			continue
		}
		tex, _ := alphabetAtlas.GetTexture(stage, l)
		t.texMap[l] = tex
	}

}

func (t *text) onStop() {
	// fmt.Println("newText:onStop")
	t.ClearProps()
	t.texMap = make(map[string]*sdl.Texture)
}

func (t *text) onUpdate(ticks uint32) {
	// fmt.Println("newText:onUpdate")
	if !t.Visible {
		return
	}

	x, y, _ := t.Pos.Int32()

	for _, r := range t.value {
		x1, y1 := convertXY(t.Scene, x, y)

		tex := t.texMap[string(r)]

		p := gfx.NewProp(string(r), tex)
		p.Pos.SetInt32(x1, y1, 0)
		p.Scale = t.Scale

		t.AddProp(p)

		_, _, w, _, _ := tex.Query()
		x += w
	}
}
