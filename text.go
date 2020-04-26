package main

import (
	"github.com/SimonBuckner/screen2d"
)

type text struct {
	X, Y, Z  float32
	scale    float32
	game     *game
	value    string
	atlas    *screen2d.SpriteAtlas
	keys     map[rune]int32
	colWidth int32
	visible  bool
}

func newText(game *game) *text {
	t := &text{
		game:     game,
		value:    "",
		atlas:    screen2d.NewSpriteAtlas(game.screen.Rend()),
		keys:     make(map[rune]int32),
		scale:    1.0,
		colWidth: 7,
	}
	t.scale = game.scale

	return t
}

func (t *text) load(atlas *screen2d.SpriteAtlas, keys map[rune]int32) {
	t.atlas = atlas
	t.keys = keys
}

func (t *text) drawText() {
	if t.visible == false {
		return
	}
	x, y := t.X, t.Y
	for _, r := range t.value {
		if tileY, ok := t.keys[r]; ok {
			newX, newY := translatePos(x, y, t.scale)
			t.atlas.DrawTileAt(0, tileY, newX, newY, t.scale)
		}
		x += float32(t.colWidth)
	}
}

func (t *text) setText(value string) {
	t.value = value
}
