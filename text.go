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

	err := t.atlas.LoadRGBAPixels(alphabet.Pixels, alphabet.Pitch, 7, 8)
	if err != nil {
		panic(err)
	}

	t.keys['A'] = 0
	t.keys['B'] = 1
	t.keys['C'] = 2
	t.keys['D'] = 3
	t.keys['E'] = 4
	t.keys['F'] = 5
	t.keys['G'] = 6
	t.keys['H'] = 7
	t.keys['I'] = 8
	t.keys['J'] = 9
	t.keys['K'] = 10
	t.keys['L'] = 11
	t.keys['M'] = 12
	t.keys['N'] = 13
	t.keys['O'] = 14
	t.keys['P'] = 15
	t.keys['Q'] = 16
	t.keys['R'] = 17
	t.keys['S'] = 18
	t.keys['T'] = 19
	t.keys['U'] = 20
	t.keys['V'] = 21
	t.keys['W'] = 22
	t.keys['X'] = 23
	t.keys['Y'] = 24
	t.keys['Z'] = 25
	t.keys['0'] = 26
	t.keys['1'] = 27
	t.keys['2'] = 28
	t.keys['3'] = 29
	t.keys['4'] = 30
	t.keys['5'] = 31
	t.keys['6'] = 32
	t.keys['7'] = 33
	t.keys['8'] = 34
	t.keys['9'] = 35
	t.keys['<'] = 36
	t.keys['>'] = 37
	t.keys[' '] = 38
	t.keys['='] = 39
	t.keys['*'] = 40
	t.keys['y'] = 41
	t.keys['-'] = 42

	return t
}

func (t *text) drawText() {
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
