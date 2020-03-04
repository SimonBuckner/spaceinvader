package main

import (
	"github.com/SimonBuckner/spaceinvader/gfx"
	"github.com/veandco/go-sdl2/sdl"
)

const testSceneName = "Test Scene"

type testScene struct {
	*gfx.Scene
	// p1       *player
	// p1Score  *score
}

func newTestScene(stage *gfx.Stage) (*testScene, error) {
	ts := &testScene{
		Scene: gfx.NewScene(testSceneName),
	}

	if err := stage.AttachScene(ts.Scene); err != nil {
		return nil, err
	}

	ts.SetKeyboardEvent(ts.keyb)
	ts.SetUpdateEvent(ts.update)

	// scene.SetStartEvent(scene.start)
	// p1, err := newPlayer(ts.scene..)
	// if err != nil {
	// 	return nil, fmt.Errorf("eorror getting new testScene; %v", err)
	// }
	// s.p1 = p1
	// p1Score, err := newScore(director, scoreP1Pos)
	// if err != nil {
	// 	return nil, fmt.Errorf("eorror getting score for testScene; %v", err)
	// }
	// s.p1Score = p1Score

	return ts, nil
}

// func (s *testScene) Load()
// func (s *testScene) start() {
// 	s.p1.Reset()
// 	s.director.stage.AddProp(s.p1.Prop)

// 	s.p1Score.Reset()
// 	for _, score := range s.p1Score.props {
// 		s.director.stage.AddProp(score)
// 	}
// }

func (s *testScene) keyb(e *sdl.KeyboardEvent) {
	if e.Type == sdl.KEYUP {
		switch e.Keysym.Scancode {
		case sdl.SCANCODE_V:
			// if s.p1.Visible() {
			// 	s.p1.SetVisible(false)
			// } else {
			// 	s.p1.Reset()
			// 	s.p1.SetVisible(true)
			// }
		}
	}
	// 	case sdl.SCANCODE_H:
	// 		if s.p1.IsVisible() && s.p1.exploding == false {
	// 			s.p1.Hit()
	// 		}
	// 	case sdl.SCANCODE_LEFT:
	// 		s.p1.x = float32(int32(s.p1.x))
	// 	case sdl.SCANCODE_RIGHT:
	// 		s.p1.x = float32(int32(s.p1.x))
	// 	}
	// }
	// if e.Type == sdl.KEYDOWN {
	// 	switch e.Keysym.Scancode {
	// 	case sdl.SCANCODE_LEFT:
	// 		s.p1.MoveLeft()
	// 	case sdl.SCANCODE_RIGHT:
	// 		s.p1.MoveRight()
	// 	}
	// }
}

func (s *testScene) update(ticks uint32) {
	// kb := s.director.stage.Keyboard()
	// if kb.IsKeyDown(sdl.SCANCODE_LEFT) {
	// 	s.p1.MoveLeft()
	// } else if kb.IsKeyDown(sdl.SCANCODE_RIGHT) {
	// 	s.p1.MoveRight()
	// }

	// s.p1.update(ticks)
}
