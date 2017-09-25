package utils

import (
	"github.com/veandco/go-sdl2/sdl"
	"main/types"
)

const (
	WINDOW_HEIGHT = 640
	WINDOW_WIDTH = 800
	WINDOW_DELAY = 4000
)

func DrawCardioid(values []types.CardioidPart) {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}
	defer sdl.Quit()

	window, renderer, err := sdl.CreateWindowAndRenderer(WINDOW_WIDTH, WINDOW_HEIGHT, sdl.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}
	defer window.Destroy()

	renderer.SetDrawColor( 0, 0, 0, sdl.ALPHA_OPAQUE)
	renderer.Clear()

	for _, value := range values {
		renderer.SetDrawColor(value.Color.R, value.Color.G, value.Color.B, sdl.ALPHA_OPAQUE)
		renderer.DrawPoints(value.Points)
	}

	renderer.Present()

	sdl.Delay(WINDOW_DELAY)
}


/* TODO: For the nearest future ...
	var event sdl.Event

	for true {
		event = sdl.PollEvent()
		switch event {
			case sdl.MOUSEBUTTONUP:
				break
			case sdl.MOUSEBUTTONDOWN:
				break
		}
	}*/