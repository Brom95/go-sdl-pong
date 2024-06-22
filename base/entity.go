package base

type Entity struct {
	X, Y           int
	Speed          float32
	Xv             float32
	Yv             float32
	AlwaysOnScreen bool
}

func (e *Entity) Update() {
	e.X += int(e.Xv * e.Speed)
	e.Y += int(e.Yv * e.Speed)
	if e.AlwaysOnScreen {
		if e.Y > World.Screen.Height {
			e.Y = World.Screen.Height
		}
		if e.Y < 0 {
			e.Y = 0

		}
		if e.X > World.Screen.Width {
			e.X = World.Screen.Width
		}
		if e.X < 0 {
			e.X = 0

		}
	}
}
