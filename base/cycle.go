package base

type Cycle struct {
	Radius int
	Entity
	Color
}

func (c *Cycle) Draw() {
	for y := -c.Radius; y < c.Radius; y++ {
		for x := -c.Radius; x < c.Radius; x++ {
			if x*x+y*y < c.Radius*c.Radius {
				World.Screen.SetPixel(int(c.X)+x, int(c.Y)+y, c.Color)
			}
		}
	}
}
