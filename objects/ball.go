package objects

import (
	"math"
	"pong/base"
)

type Ball struct {
	base.Color
	base.Move
	base.Cycle
}

// Collide implements base.DrawerUpdater.
func (b *Ball) Collide(drawer base.Drawer) {
	var paddle *Paddle
	switch drawer.(type) {
	case *Paddle:
		paddle = drawer.(*Paddle)

	case *AiPaddle:
		paddle = &drawer.(*AiPaddle).Paddle

	default:
		return
	}
	if b.intersects(paddle) {
		b.Xv = -b.Xv
		b.X += int(b.Xv) * 5
		b.Yv = -b.Yv
	}
}

func (b *Ball) Draw() {
	for y := -b.Radius; y < b.Radius; y++ {
		for x := -b.Radius; x < b.Radius; x++ {
			if x*x+y*y < b.Radius*b.Radius {
				base.World.Screen.SetPixel(int(b.X)+x, int(b.Y)+y, b.Color)
			}
		}
	}
}

func (b *Ball) Update() {
	b.Move.Update()
	if b.Y-(b.Radius) <= 0 || b.Y+(b.Radius) >= (base.World.Screen.Height) {
		b.Yv = -b.Yv
	}
	if b.X-(b.Radius) <= 0 || b.X+(b.Radius) >= (base.World.Screen.Width) {
		b.X = 300
		b.Y = 300
	}
}
func (b *Ball) intersects(p *Paddle) bool {
	circleDistanceX := math.Abs(float64(b.X - p.X))
	circleDistanceY := math.Abs(float64(b.Y - p.Y))

	if circleDistanceX > float64(p.W/2+b.Radius) {
		return false
	}
	if circleDistanceY > float64(p.H/2+b.Radius) {
		return false
	}

	if circleDistanceX <= float64(p.W/2) {
		return true
	}
	if circleDistanceY <= float64(p.H/2) {
		return true
	}

	cornerDistance_sq := math.Pow(circleDistanceX-float64(p.W/2), 2) +
		math.Pow((circleDistanceY-float64(p.H/2)), 2)

	return (cornerDistance_sq <= float64(b.Radius*b.Radius))
}