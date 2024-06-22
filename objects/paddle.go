package objects

import (
	"pong/base"

	"github.com/veandco/go-sdl2/sdl"
)

type Paddle struct {
	base.Square
}

// Collide implements base.DrawerUpdater.
func (p *Paddle) Collide(base.Drawer) {

}

func (p *Paddle) Draw() {
	startX := int(p.X) - p.W/2
	startY := int(p.Y) - p.H/2

	for y := 0; y < p.H; y++ {
		for x := 0; x < p.W; x++ {
			base.World.Screen.SetPixel(startX+x, startY+y, p.Color)
		}
	}
}

func (p *Paddle) Update() {
	if base.World.KeyState[sdl.SCANCODE_UP] != 0 {
		p.Yv = -1
	} else if base.World.KeyState[sdl.SCANCODE_DOWN] != 0 {
		p.Yv = 1
	} else {
		p.Yv = 0
	}

	p.Square.Update()
}

type AiPaddle struct {
	Paddle
	Ball *Ball
}

func (p *AiPaddle) Update() {
	if p.Ball == nil {
		p.Ball = base.World.GetObjectToUpdate("ball").(*Ball)
	}

	if p.Y+p.H/2 <= p.Ball.Y {
		p.Yv = 1
	} else if p.Y-p.H/2 >= p.Ball.Y {
		p.Yv = -1
	} else {
		p.Yv = 0
	}

	p.Paddle.Square.Update()

}
