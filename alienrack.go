package main

import (
	"math"

	"github.com/SimonBuckner/screen2d"
)

type arState int

const (
	arReady arState = iota
	arMoving
	arExploding
)

type alienRack struct {
	aliens       []*alien
	game         *game
	state        arState
	level        int
	x, y         float32
	scale        float32
	currentFrame int
	cursor       int
	stepL        float32
	stepR        float32
	direction    int
	timer        int
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

	i := 0
	x := alienStartX
	y := alienStartY
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
			ar.aliens[i].X = float32(x)
			ar.aliens[i].Y = float32(y)
			x = x + alienColWidth

			i++
		}
		x = alienStartX
		y = y - alienRowHeight
	}

	ar.currentFrame = 0
	ar.cursor = -1
	ar.stepL = 2.0
	ar.stepR = 2.0
	ar.state = arReady
	ar.timer = 30
	ar.direction = 1
}

func (ar *alienRack) update(ticks uint32, elapsed float32, p *player, shot *playerShot) {
	switch ar.state {
	case arReady:
		ar.timer--
		if ar.timer == 0 {
			ar.state = arMoving
		}
	case arMoving:
		ar.move()
		ar.aliens[ar.cursor].frame = ar.currentFrame

		if alien := ar.checkForHit(shot); alien != nil {
			shot.setHit()
			ar.state = arExploding
			ar.timer = 10
			p.score += alien.score
		}
	case arExploding:
		ar.timer--
		if ar.timer == 0 {
			ar.state = arMoving
			for i, a := range ar.aliens {
				if a.state == alienExploding {
					ar.aliens[i].state = alienDead
				}
			}
		}
	}
	for _, a := range ar.aliens {
		a.update(ticks, elapsed)
	}
}

func (ar *alienRack) checkForHit(shot *playerShot) *alien {
	for i := range ar.aliens {
		if ar.aliens[i].state == alienAlive && screen2d.CheckBoxHit(ar.aliens[i], shot) {
			if screen2d.CheckPixelHit(shot, ar.aliens[i]) {
				ar.aliens[i].setHit()
				shot.setHit()
				return ar.aliens[i]
			}
		}
	}
	return nil
}

func (ar *alienRack) advanceCursor() {
	if ar.state != arMoving {
		return
	}
	for {
		ar.cursor++
		if ar.cursor >= len(ar.aliens) {
			ar.cursor = 0
			if ar.currentFrame == 0 {
				ar.currentFrame = 1
			} else {
				ar.currentFrame = 0
			}
			ar.checkBounds()
		}
		if ar.aliens[ar.cursor].state == alienAlive {
			return
		}
	}
}

func (ar *alienRack) checkBounds() {
	minX := float32(math.MaxFloat32)
	maxX := float32(-1)

	for i := range ar.aliens {
		if ar.aliens[i].state == alienAlive {
			if ar.aliens[i].X < minX {
				minX = ar.aliens[i].X
			}
			if ar.aliens[i].X > maxX {
				maxX = ar.aliens[i].X
			}
		}
	}
	if ar.direction < 0 && minX <= 0 {
		ar.direction = +1
	}

	if ar.direction > 0 && maxX >= (originalWidth-alienColWidth) {
		ar.direction = -1
	}
}

func (ar *alienRack) remainCount() int {
	rc := 0
	for _, a := range ar.aliens {
		if a.state == alienAlive {
			rc++
		}
	}
	return rc
}

func (ar *alienRack) move() {

	if ar.direction > 0 {
		ar.aliens[ar.cursor].X += ar.stepR
	} else {
		ar.aliens[ar.cursor].X -= ar.stepL
	}

}

func (ar *alienRack) drawRack() {
	for _, a := range ar.aliens {
		a.Draw()
	}
}
