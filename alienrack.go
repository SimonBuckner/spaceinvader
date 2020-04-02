package main

type alienRack struct {
	aliens []*alien
	game   *game
	level  int
	x, y   float32
	scale  float32
}

func newAlienRack(game *game) *alienRack {
	ar := &alienRack{
		aliens: make([]*alien, alienRows*alienCols),
		game:   game,
	}
	for i := range ar.aliens {
		ar.aliens[i] = newAlien(game)
		ar.aliens[i].SetCalcScreenXYFunc(translatePos)
	}
	return ar
}

func (ar *alienRack) reset(level int) {

	ar.x = alienStartX
	ar.y = alienStartY

	i := 0
	for row := int32(0); row < alienRows; row++ {
		for col := int32(0); col < alienCols; col++ {
			switch row {
			case 0, 1:
				ar.aliens[i].setBreed(alienCrab)
			case 2, 3:
				ar.aliens[i].setBreed(alienOctopus)
			case 4:
				ar.aliens[i].setBreed(alienSquid)
			}
			ar.aliens[i].reset()
			i++
		}
	}
}

func (ar *alienRack) update(ticks uint32, elapsed float32, shipX float32) {
	x := ar.x
	y := ar.y

	i := 0
	for row := int32(0); row < alienRows; row++ {
		for col := int32(0); col < alienCols; col++ {
			ar.aliens[i].X = float32(x)
			ar.aliens[i].Y = float32(y)
			x = x + alienColWidth
			ar.aliens[i].update(ticks, elapsed)
			i++
		}
		x = alienStartX
		y = y - alienRowHeight
	}
}

func (ar *alienRack) draw() {
	for _, a := range ar.aliens {
		a.Draw()
	}
}
