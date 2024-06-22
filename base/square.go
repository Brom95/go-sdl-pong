package base

type Square struct {
	Entity
	Color
	W int
	H int
}

func (p *Square) Draw() {
	startX := int(p.X) - p.W/2
	startY := int(p.Y) - p.H/2

	for y := 0; y < p.H; y++ {
		for x := 0; x < p.W; x++ {
			World.Screen.SetPixel(startX+x, startY+y, p.Color)
		}
	}
}
