package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/veandco/go-sdl2/sdl"
	ttf "github.com/veandco/go-sdl2/ttf"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(2)
	}
}

func run() error {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		return fmt.Errorf("could not initialize SDL: %v", err)
	}

	defer sdl.Quit()

	if err := ttf.Init(); err != nil {
		return fmt.Errorf("could not initialize TTF: %v", err)
	}

	defer ttf.Quit()

	w, r, err := sdl.CreateWindowAndRenderer(800, 600, sdl.WINDOW_SHOWN)

	if err != nil {
		return fmt.Errorf("could not create a window: %v", err)
	}

	defer w.Destroy()

	if err := drawTitle(r, "Flappy Gopher"); err != nil {
		return fmt.Errorf("could not draw title: %v", err)
	}

	log.Printf("Flappy Gopher is running")

	time.Sleep(2 * time.Second)

	s, err := newScene(r)

	if err != nil {
		return fmt.Errorf("could not create scene: %v", err)
	}

	defer s.destroy()

	events := make(chan sdl.Event)

	go func() {
		for {
			events <- sdl.WaitEvent()
		}
	}()

	return s.run(events, r)
}

func drawTitle(r *sdl.Renderer, text string) error {
	r.Clear()

	f, err := ttf.OpenFont("res/fonts/Flappy.ttf", 20)

	if err != nil {
		return fmt.Errorf("could not load font: %v", err)
	}

	defer f.Close()

	color := sdl.Color{R: 255, G: 100, B: 0, A: 255}

	s, err := f.RenderUTF8Solid(text, color)

	if err != nil {
		return fmt.Errorf("could not render title: %v", err)
	}

	defer s.Free()

	t, err := r.CreateTextureFromSurface(s)

	if err != nil {
		return fmt.Errorf("could not create texture: %v", err)
	}

	defer t.Destroy()

	if err := r.Copy(t, nil, nil); err != nil {
		return fmt.Errorf("could not copy texture: %v", err)
	}

	r.Present()

	return nil
}
