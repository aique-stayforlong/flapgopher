package main

import (
	"fmt"
	"log"
	"time"
	"github.com/veandco/go-sdl2/sdl"
	img "github.com/veandco/go-sdl2/img"
)

type scene struct {
	background *sdl.Texture
	bird *bird
	pipes *pipes
}

func newScene(r *sdl.Renderer) (*scene, error) {
	background, err := img.LoadTexture(r, "res/img/background.png")

	if err != nil {
		return nil, fmt.Errorf("could not load background image: %v", err)
	}

	bird, err := newBird(r)

	if err != nil {
		return nil, fmt.Errorf("could not create bird: %v", err)
	}

	pipes, err := newPipes(r)

	if err != nil {
		return nil, fmt.Errorf("could not create pipe: %v", err)
	}

	return &scene{ background, bird, pipes }, nil
}

func (s *scene) run(events <-chan sdl.Event, r *sdl.Renderer) error {
	refresh := time.NewTicker(50 * time.Millisecond)
	defer refresh.Stop()
	
	quit := false

	for !quit {
		select {
		case e := <-events:
			quit = s.handleEvent(e)
		case <-refresh.C: 
			s.update()

			if s.bird.isDead() {
				drawTitle(r, "Game Over")
				time.Sleep(time.Second)
				s.restart()
			}

			if err := s.paint(r) ; err != nil {
				return err
			}
		}
	}

	return nil
}

func (s *scene) handleEvent(e sdl.Event) bool {
	switch eType := e.(type) {
		case *sdl.QuitEvent:
			return true
		case *sdl.KeyboardEvent:
			s.bird.jump()
		case *sdl.MouseMotionEvent, *sdl.WindowEvent, *sdl.AudioDeviceEvent: // nothing to do
		default:
			log.Printf("unknown event %T", eType)
	}

	return false
}

func (s *scene) update() {
	s.bird.update()
	s.pipes.update()
	s.pipes.touch(s.bird)
}

func (s *scene) paint(r *sdl.Renderer) error {
	r.Clear()

	if err := r.Copy(s.background, nil, nil) ; err != nil {
		return fmt.Errorf("could not paint background: %v", err)
	}

	if err := s.bird.paint(r) ; err != nil {
		return fmt.Errorf("could not paint bird: %v", err)
	}

	if err := s.pipes.paint(r) ; err != nil {
		return fmt.Errorf("could not paint pipe: %v", err)
	}

	r.Present()

	return nil
}

func (s *scene) restart() {
	s.bird.restart()
	s.pipes.restart()
}

func (s *scene) destroy() {
	s.background.Destroy()
	s.bird.destroy()
	s.pipes.destroy()
}