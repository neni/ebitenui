package input

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

var (
	LeftMouseButtonPressed   bool
	MiddleMouseButtonPressed bool
	RightMouseButtonPressed  bool
	CursorX                  int
	CursorY                  int
	WheelX                   float64
	WheelY                   float64

	LeftMouseButtonJustPressed   bool
	MiddleMouseButtonJustPressed bool
	RightMouseButtonJustPressed  bool

	LastLeftMouseButtonPressed   bool
	LastMiddleMouseButtonPressed bool
	LastRightMouseButtonPressed  bool

	InputChars    []rune
	KeyPressed    = map[ebiten.Key]bool{}
	AnyKeyPressed bool

	touchIDs        []ebiten.TouchID
	InputTouchs     = map[ebiten.TouchID]*Touch{}

)

// Touch is stats of touchs
type Touch struct {
	//Id          ebiten.TouchID
	CursorX     int
	CursorY     int
	JustPressed bool
	LastPressed bool
}


// Update updates the input system. This is called by the UI.
func Update() {
	LeftMouseButtonPressed = ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft)
	MiddleMouseButtonPressed = ebiten.IsMouseButtonPressed(ebiten.MouseButtonMiddle)
	RightMouseButtonPressed = ebiten.IsMouseButtonPressed(ebiten.MouseButtonRight)
	CursorX, CursorY = ebiten.CursorPosition()

	wx, wy := ebiten.Wheel()
	WheelX += wx
	WheelY += wy

	InputChars = append(InputChars, ebiten.InputChars()...)

	AnyKeyPressed = false
	for k := ebiten.Key(0); k <= ebiten.KeyMax; k++ {
		p := ebiten.IsKeyPressed(k)
		KeyPressed[k] = p

		if p {
			AnyKeyPressed = true
		}
	}

	// update touchs
	for id, t := range InputTouchs {
		if(t.LastPressed) {
			delete(InputTouchs, id)
		}else{
			if inpututil.IsTouchJustReleased(id) {
				t.LastPressed = true
			}
			t.JustPressed = false
			t.CursorX, t.CursorY = ebiten.TouchPosition(id)
		}
	}
	// add new touchs
	newTouchs := inpututil.AppendJustPressedTouchIDs(touchIDs[:0]);
	for _, id := range newTouchs {
	  newTouch := new(Touch)
		newTouch.JustPressed = true
		newTouch.LastPressed = false
		newTouch.CursorX, newTouch.CursorY = ebiten.TouchPosition(id)
		InputTouchs[id] = newTouch
	}

	// temp: emulate mouse
	for _, t := range InputTouchs {
		switch len(InputTouchs) {
			case 1:
				LeftMouseButtonJustPressed = /*LeftMouseButtonJustPressed ||*/ t.JustPressed
				LastLeftMouseButtonPressed = /*LastLeftMouseButtonPressed ||*/ t.LastPressed
				LeftMouseButtonPressed = true
				CursorX = t.CursorX
				CursorY = t.CursorY
		}
	}


}

// Draw updates the input system. This is called by the UI.
func Draw() {
	LeftMouseButtonJustPressed = LeftMouseButtonPressed && LeftMouseButtonPressed != LastLeftMouseButtonPressed
	MiddleMouseButtonJustPressed = MiddleMouseButtonPressed && MiddleMouseButtonPressed != LastMiddleMouseButtonPressed
	RightMouseButtonJustPressed = RightMouseButtonPressed && RightMouseButtonPressed != LastRightMouseButtonPressed

	LastLeftMouseButtonPressed = LeftMouseButtonPressed
	LastMiddleMouseButtonPressed = MiddleMouseButtonPressed
	LastRightMouseButtonPressed = RightMouseButtonPressed
}

// AfterDraw updates the input system after the Ebiten Draw function has been called. This is called by the UI.
func AfterDraw() {
	InputChars = InputChars[:0]
	WheelX, WheelY = 0, 0
	//touchIDs = touchIDs[:0]
}
