package main

import (
	"fmt"
	"time"
	"context"
	"github.com/veandco/go-sdl2/sdl"
	img "github.com/veandco/go-sdl2/img"
)

type scene struct {
	time int
	background *sdl.Texture
	birds []*sdl.Texture
}

func newScene(r *sdl.Renderer) (*scene, error) {
	background, err := img.LoadTexture(r, "res/img/background.png")

	if err != nil {
		return nil, fmt.Errorf("could not load background image: %v", err)
	}

	var birds []*sdl.Texture

	for i := 1 ; i <= 4 ; i++ {
		path := fmt.Sprintf("res/img/bird_frame_%d.png", i)
		bird, err := img.LoadTexture(r, path)

		if err != nil {
			return nil, fmt.Errorf("could not load bird image: %v", err)
		}

		birds = append(birds, bird)
	}

	return &scene{ 0, background, birds }, nil
}

func (s *scene) run(context context.Context, r *sdl.Renderer) error {
	refresh := time.NewTicker(50 * time.Millisecond)
	defer refresh.Stop()

	end := time.NewTicker(5 * time.Second)
	defer end.Stop()
	

	for {
		select {
		case <-context.Done():
			return nil
		case <-refresh.C: 
			if err := s.paint(r) ; err != nil {
				return err
			}
		case <- end.C:
			return nil
		}
	}

	return nil
}

func (s *scene) paint(r *sdl.Renderer) error {
	s.time++

	r.Clear()

	if err := r.Copy(s.background, nil, nil) ; err != nil {
		return fmt.Errorf("could not copy texture: %v", err)
	}

	rect := &sdl.Rect{ X: 10, Y: 300 - 43/2, W: 50, H: 43 }
	bird := s.birds[s.time % len(s.birds)]

	if err := r.Copy(bird, nil, rect) ; err != nil {
		return fmt.Errorf("could not copy texture: %v", err)
	}

	r.Present()

	return nil
}

func (s *scene) destroy() {
	s.background.Destroy()

	for _, bird := range s.birds {
		bird.Destroy()
	}
}