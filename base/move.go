package base

type Move struct {
	Position
	Speed          float32
	Xv             float32
	Yv             float32
	AlwaysOnScreen bool
}

func (b *Move) Update() {

	b.X += int(b.Xv * b.Speed)
	b.Y += int(b.Yv * b.Speed)
	if b.AlwaysOnScreen {
		if b.Y > World.Screen.Height {
			b.Y = World.Screen.Height
		}
		if b.Y < 0 {
			b.Y = 0

		}
		if b.X > World.Screen.Width {
			b.X = World.Screen.Width
		}
		if b.X < 0 {
			b.X = 0

		}
	}

}
