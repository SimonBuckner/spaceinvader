package main

import (
	"github.com/SimonBuckner/spaceinvader/gfx"
	"github.com/veandco/go-sdl2/sdl"
)

type text struct {
	*gfx.Actor
	value    string
	texMap   map[string]*sdl.Texture
	display  []*gfx.Prop
	reqChars string
}

func newText(game *game, reqChars string) *text {
	t := &text{
		Actor:    gfx.NewActor("text"),
		texMap:   make(map[string]*sdl.Texture),
		reqChars: reqChars,
	}
	t.StartEventHandler = t.onStart
	t.StopEventHandler = t.onStop
	return t
}

func (t *text) onStart() {
	stage := t.Scene.Stage
	t.Visible = true
	for _, r := range t.reqChars {
		l := string(r)
		tex, _ := alphabetAtlas.GetTexture(stage, l)
		t.texMap[l] = tex
	}
}

func (t *text) onStop() {
	t.ClearProps()
	t.texMap = make(map[string]*sdl.Texture)
}

func (t *text) update(ticks uint32) {
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
