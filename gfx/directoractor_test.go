package gfx

import (
	"fmt"
	"testing"

	"github.com/veandco/go-sdl2/sdl"
)

func TestDirector(t *testing.T) {
	d := NewDirector()

	testStatus(t, d, true, false, false)

	{
		actor1, err := d.NewActor("first")
		if err != nil {
			fmt.Errorf("unexpected error; expecting nil, got %v", err)
		}
		if actor1.name != "first" {
			fmt.Errorf("unexpected state name; expecting 'first', got '%v'", actor1.name)
		}
		if len(d.actors) != 1 {
			fmt.Errorf("unexpected actors length; expecting 1, got %v", len(d.actors))
		}
		testStatus(t, d, true, false, false)
	}

	{
		_, err := d.NewActor("first")
		if err == nil {
			fmt.Errorf("unexpected success; expecting error, got nil")
		}
		if len(d.actors) != 1 {
			fmt.Errorf("unexpected actors length; expecting 1, got %v", len(d.actors))
		}
		testStatus(t, d, true, false, false)

	}

	{
		actor3, err := d.NewActor("third")
		if err != nil {
			fmt.Errorf("unexpected error; expecting nil, got %v", err)
		}
		if actor3.name != "first" {
			fmt.Errorf("unexpected state name; expecting 'third', got '%v'", actor3.name)
		}
		if len(d.actors) != 2 {
			fmt.Errorf("unexpected actors length; expecting 1, got %v", len(d.actors))
		}
		testStatus(t, d, true, false, false)

	}
}

func TestStartActor(t *testing.T) {
	d := NewDirector()
	actor1, _ := d.NewActor("first")
	actor2, _ := d.NewActor("second")
	{
		err := d.StartActor("bad")
		if err == nil {
			fmt.Errorf("unexpected success; expecting error, got nil")
		}
		testStatus(t, d, true, false, false)
	}

	{
		err := d.StartActor("first")
		if err != nil {
			fmt.Errorf("unexpected error; expecting nil, got %v", err)
		}
		testStatus(t, d, false, true, false)
		if d.current != actor1 {
			fmt.Errorf("unexpected state; expected 'first', got '%v'", d.current.name)
		}
	}

	{
		err := d.StartActor("second")
		if err != nil {
			fmt.Errorf("unexpected error; expecting nil, got %v", err)
		}
		testStatus(t, d, false, true, false)
		if d.current != actor2 {
			fmt.Errorf("unexpected state; expected 'second', got '%v'", d.current.name)
		}
	}

}

func testStatus(t *testing.T, d *Director, loading, running, closing bool) {
	if d.IsLoading() != loading {
		fmt.Errorf("expected loading to be %v, got %v", loading, d.IsLoading())
	}
	if d.IsRunning() != running {
		fmt.Errorf("expected running to be %v, got %v", loading, d.IsRunning())
	}
	if d.IsClosing() != closing {
		fmt.Errorf("expected flosing to be %v, got %v", loading, d.IsClosing())
	}
}

func TestKeyboard(t *testing.T) {
	d := NewDirector()
	a, _ := d.NewActor("a")
	b, _ := d.NewActor("b")

	d.SetKeyboardEvent(func(e *sdl.KeyboardEvent) {
		e.Timestamp = 1
	})

	a.SetKeyboardEvent(func(e *sdl.KeyboardEvent) {
		e.WindowID = 2
	})

	b.SetKeyboardEvent(func(e *sdl.KeyboardEvent) {
		e.WindowID = 3
	})

	e := &sdl.KeyboardEvent{
		Timestamp: 0,
		WindowID:  0,
	}

	d.KeyboardEvent(e)
	if e.Timestamp != 1 {
		t.Errorf("director keyboard did not fire")
	}
	if e.WindowID == 1 {
		t.Errorf("actor keyboard event fired unexpectedly")
	}

	d.StartActor("a")
	d.KeyboardEvent(e)
	if e.Timestamp != 1 {
		t.Errorf("actor keyboard event did not fire")
	}
	if e.WindowID != 2 {
		t.Errorf("actor 'a' keyboard event did not fire")
	}
	if e.WindowID == 3 {
		t.Errorf("actor 'b' keyboard event fired unexpectedly")
	}

	d.StartActor("b")
	e.Timestamp = 0
	e.WindowID = 0

	d.KeyboardEvent(e)
	if e.Timestamp != 1 {
		t.Errorf("actor keyboard event did not fire")
	}
	if e.WindowID == 2 {
		t.Errorf("actor 'a' keyboard event fired unexpectedly")
	}
	if e.WindowID != 3 {
		t.Errorf("actor 'b' keyboard event did not fire")
	}
}
