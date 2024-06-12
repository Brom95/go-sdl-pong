package main

import (
	"pong/base"
	"pong/objects"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	WinWidth  int32 = 800
	WinHeight int32 = 600
)

func main() {

	window, err := sdl.CreateWindow("ИГРА!", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, WinWidth, WinHeight, sdl.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}
	defer window.Destroy()

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		panic(err)
	}
	defer renderer.Destroy()

	tex, err := renderer.CreateTexture(sdl.PIXELFORMAT_RGBA8888, sdl.TEXTUREACCESS_STREAMING, WinWidth, WinHeight)
	if err != nil {
		panic(err)
	}
	defer tex.Destroy()
	screen := base.NewScreen(int(WinWidth), int(WinHeight))
	keyState := sdl.GetKeyboardState()
	base.NewWorld(screen, keyState)
	player1 := objects.Paddle{

		Move: base.Move{
			Position:       base.Position{X: 50, Y: 100},
			Speed:          5,
			Xv:             0,
			Yv:             0,
			AlwaysOnScreen: true,
		},
		Color:  base.Color{R: 255, G: 255, B: 255},
		Square: base.Square{W: 20, H: 100},
	}
	ball := objects.Ball{
		Color: base.Color{R: 255, G: 255, B: 255},
		Cycle: base.Cycle{Radius: 10},
		Move:  base.Move{Speed: 3, Xv: -1, Yv: 1, Position: base.Position{X: 500, Y: 300}},
	}
	player2 := objects.AiPaddle{
		Paddle: objects.Paddle{
			Move: base.Move{
				Position:       base.Position{X: screen.Width - 50, Y: 100},
				Speed:          5,
				Xv:             0,
				Yv:             0,
				AlwaysOnScreen: true,
			},
			Color:  base.Color{R: 255, G: 255, B: 255},
			Square: base.Square{W: 20, H: 100},
		},
	}
	base.World.AddDrawerUpdater("ball", &ball)
	base.World.AddDrawerUpdater("player1", &player1)
	base.World.AddDrawerUpdater("player2", &player2)
	for {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				return
			}
		}
		base.World.Update()
		base.World.Draw()
		tex.Update(nil, screen.UnsafePointer(), int(WinWidth)*4)

		renderer.Copy(tex, nil, nil)
		renderer.Present()

		screen.Clear()
		sdl.Delay(16)

	}
}
