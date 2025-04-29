package main

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
	img "github.com/veandco/go-sdl2/img"
)

type bird struct {
	time int
	frames []*sdl.Texture

	y, speed float64
}

const (
	gravity = 0.75
	jumpSpeed = 10
)

func newBird(r *sdl.Renderer) (*bird, error) {
	var frames []*sdl.Texture

	for i := 1 ; i <= 4 ; i++ {
		path := fmt.Sprintf("res/img/bird_frame_%d.png", i)
		bird, err := img.LoadTexture(r, path)

		if err != nil {
			return nil, fmt.Errorf("could not load bird image: %v", err)
		}

		frames = append(frames, bird)
	}

	return &bird{ frames: frames, y: 300 }, nil
}

func (b *bird) paint(r *sdl.Renderer) error {
	b.time++
	b.y -= b.speed
	b.speed += gravity

	if b.y < 0 {
		b.speed = -b.speed
		b.y = 0
	}

	rect := &sdl.Rect{ X: 10, Y: (600 - int32(b.y)) - 43/2, W: 50, H: 43 }
	frame := b.frames[b.time % len(b.frames)]

	if err := r.Copy(frame, nil, rect) ; err != nil {
		return fmt.Errorf("could not copy texture: %v", err)
	}

	return nil
}

func (b *bird) jump() {
	b.speed = -jumpSpeed
}

func (b *bird) destroy() {
	for _, frame := range b.frames {
		frame.Destroy()
	}
}