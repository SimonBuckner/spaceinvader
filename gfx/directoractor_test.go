package gfx

// func TestDirector(t *testing.T) {
// 	d := NewDirector()

// 	testStatus(t, d, true, false, false)

// 	{
// 		scene1, err := d.NewScene("first")
// 		if err != nil {
// 			fmt.Errorf("unexpected error; expecting nil, got %v", err)
// 		}
// 		if scene1.name != "first" {
// 			fmt.Errorf("unexpected state name; expecting 'first', got '%v'", scene1.name)
// 		}
// 		if len(d.scenes) != 1 {
// 			fmt.Errorf("unexpected scenes length; expecting 1, got %v", len(d.scenes))
// 		}
// 		testStatus(t, d, true, false, false)
// 	}

// 	{
// 		_, err := d.NewScene("first")
// 		if err == nil {
// 			fmt.Errorf("unexpected success; expecting error, got nil")
// 		}
// 		if len(d.scenes) != 1 {
// 			fmt.Errorf("unexpected scenes length; expecting 1, got %v", len(d.scenes))
// 		}
// 		testStatus(t, d, true, false, false)

// 	}

// 	{
// 		scene3, err := d.NewScene("third")
// 		if err != nil {
// 			fmt.Errorf("unexpected error; expecting nil, got %v", err)
// 		}
// 		if scene3.name != "first" {
// 			fmt.Errorf("unexpected state name; expecting 'third', got '%v'", scene3.name)
// 		}
// 		if len(d.scenes) != 2 {
// 			fmt.Errorf("unexpected scenes length; expecting 1, got %v", len(d.scenes))
// 		}
// 		testStatus(t, d, true, false, false)

// 	}
// }

// func TestStartScene(t *testing.T) {
// 	d := NewDirector()
// 	scene1, _ := d.NewScene("first")
// 	scene2, _ := d.NewScene("second")
// 	{
// 		err := d.StartScene("bad")
// 		if err == nil {
// 			fmt.Errorf("unexpected success; expecting error, got nil")
// 		}
// 		testStatus(t, d, true, false, false)
// 	}

// 	{
// 		err := d.StartScene("first")
// 		if err != nil {
// 			fmt.Errorf("unexpected error; expecting nil, got %v", err)
// 		}
// 		testStatus(t, d, false, true, false)
// 		if d.current != scene1 {
// 			fmt.Errorf("unexpected state; expected 'first', got '%v'", d.current.name)
// 		}
// 	}

// 	{
// 		err := d.StartScene("second")
// 		if err != nil {
// 			fmt.Errorf("unexpected error; expecting nil, got %v", err)
// 		}
// 		testStatus(t, d, false, true, false)
// 		if d.current != scene2 {
// 			fmt.Errorf("unexpected state; expected 'second', got '%v'", d.current.name)
// 		}
// 	}

// }

// func testStatus(t *testing.T, d *Director, loading, running, closing bool) {
// 	if d.IsLoading() != loading {
// 		fmt.Errorf("expected loading to be %v, got %v", loading, d.IsLoading())
// 	}
// 	if d.IsRunning() != running {
// 		fmt.Errorf("expected running to be %v, got %v", loading, d.IsRunning())
// 	}
// 	if d.IsClosing() != closing {
// 		fmt.Errorf("expected flosing to be %v, got %v", loading, d.IsClosing())
// 	}
// }

// func TestKeyboard(t *testing.T) {
// 	d := NewDirector()
// 	a, _ := d.NewScene("a")
// 	b, _ := d.NewScene("b")

// 	d.SetKeyboardEvent(func(e *sdl.KeyboardEvent) {
// 		e.Timestamp = 1
// 	})

// 	a.SetKeyboardEvent(func(e *sdl.KeyboardEvent) {
// 		e.WindowID = 2
// 	})

// 	b.SetKeyboardEvent(func(e *sdl.KeyboardEvent) {
// 		e.WindowID = 3
// 	})

// 	e := &sdl.KeyboardEvent{
// 		Timestamp: 0,
// 		WindowID:  0,
// 	}

// 	d.KeyboardEvent(e)
// 	if e.Timestamp != 1 {
// 		t.Errorf("director keyboard did not fire")
// 	}
// 	if e.WindowID == 1 {
// 		t.Errorf("scene keyboard event fired unexpectedly")
// 	}

// 	d.StartScene("a")
// 	d.KeyboardEvent(e)
// 	if e.Timestamp != 1 {
// 		t.Errorf("scene keyboard event did not fire")
// 	}
// 	if e.WindowID != 2 {
// 		t.Errorf("scene 'a' keyboard event did not fire")
// 	}
// 	if e.WindowID == 3 {
// 		t.Errorf("scene 'b' keyboard event fired unexpectedly")
// 	}

// 	d.StartScene("b")
// 	e.Timestamp = 0
// 	e.WindowID = 0

// 	d.KeyboardEvent(e)
// 	if e.Timestamp != 1 {
// 		t.Errorf("scene keyboard event did not fire")
// 	}
// 	if e.WindowID == 2 {
// 		t.Errorf("scene 'a' keyboard event fired unexpectedly")
// 	}
// 	if e.WindowID != 3 {
// 		t.Errorf("scene 'b' keyboard event did not fire")
// 	}
// }
