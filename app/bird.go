package main

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
	img "github.com/veandco/go-sdl2/img"
)

type bird struct {
	time int
	dead bool
	frames []*sdl.Texture

	x, y int32

	width, height int32
	
	speed float64
}

const (
	initialWidth = 50
	initialHeight = 43
	initialX = 10
	initialY = 300
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

	return &bird{
		frames: frames,
		x: initialX,
		y: initialY,
		width: initialWidth,
		height: initialHeight,
	}, nil
}

func (b *bird) isDead() bool {
	return b.dead
}

func (b *bird) update() {
	b.time++
	b.y -= int32(b.speed)
	b.speed += gravity

	if b.y < 0 {
		b.dead = true
	}
}

func (b *bird) paint(r *sdl.Renderer) error {
	rect := &sdl.Rect{ X: b.x, Y: 600 - b.y - b.height / 2, W: b.width, H: b.height }
	frame := b.frames[b.time % len(b.frames)]

	if err := r.Copy(frame, nil, rect) ; err != nil {
		return fmt.Errorf("could not copy bird texture: %v", err)
	}

	return nil
}

func (b *bird) jump() {
	b.speed = -jumpSpeed
}

func (b *bird) restart() {
	b.dead = false
	b.y = initialY
	b.speed = 0
}

func (b *bird) destroy() {
	for _, frame := range b.frames {
		frame.Destroy()
	}
}