package main

import (
	"fmt"
	"time"
	"math/rand"
	"github.com/veandco/go-sdl2/sdl"
	img "github.com/veandco/go-sdl2/img"
)

type pipes struct {
	img *sdl.Texture
	speed int32

	pipes []*pipe
}

const (
	nextPipeInterval = 3
	initialSpeed = 5
)

func newPipes(r *sdl.Renderer) (*pipes, error) {
	img, err := img.LoadTexture(r, "res/img/pipe.png")

	if err != nil {
		return nil, fmt.Errorf("could not load pipe image: %v", err)
	}

	ps := &pipes{ img: img, speed: initialSpeed }

	go func() {
		for {
			ps.pipes = append(ps.pipes, newPipe())
			time.Sleep(nextPipeInterval * time.Second)
		}
	}()

	return ps, nil
}

func (p *pipes) paint(r *sdl.Renderer) error {
	for _, pipe := range p.pipes {
		if err := pipe.paint(r, p.img) ; err != nil {
			return err
		}
	}

	return nil
}

func (p *pipes) update() {
	var remaining []*pipe

	for _, pipe := range p.pipes {
		pipe.x -= p.speed

		if pipe.x + pipe.width > 0 {
			remaining = append(remaining, pipe)
		}
	}

	p.pipes = remaining
}

func (p *pipes) restart() {
	p.pipes = nil
}

func (p *pipes) destroy() {
	p.img.Destroy()
}

func (p *pipes) touch(b *bird) {
	for _, pipe := range p.pipes {
		pipe.touch(b)
	}
}

type pipe struct {
	x int32
	width, height int32
	inverted bool
}

func newPipe() *pipe {
	return &pipe{
		x: 800,
		width: 50,
		height: 100 + int32(rand.Intn(300)),
		inverted: rand.Float32() > 0.5,
	}
}

func (p *pipe) paint(r *sdl.Renderer, img *sdl.Texture) error {
	rect := &sdl.Rect{ X: p.x, Y: 600 - p.height, W: p.width, H: p.height }
	flip := sdl.FLIP_NONE

	if p.inverted {
		rect.Y = 0
		flip = sdl.FLIP_VERTICAL
	}
	
	if err := r.CopyEx(img, nil, rect, 0, nil, flip) ; err != nil {
		return fmt.Errorf("could not copy pipe texture background: %v", err)
	}

	return nil
}

func (p *pipe) touch(b *bird) {
	if p.x > b.x + b.width {
		return
	}

	if p.x < b.x - b.width {
		return
	}

	if p.x > b.x + b.width {
		return
	}

	if !p.inverted && p.height < b.y - (b.height / 2) {
		return
	}

	if p.inverted && p.height > b.y - (b.height / 2) {
		return
	}

	b.dead = true
}
