package main

import "fmt"

type BisectInterface interface {
	run() error
}

type Bisect struct {
	cmd     string
	points  []string
	workers int
}

type BisectParams struct {
	Cmd     string
	Points  []string
	Workers int
}

func NewBisect(p *BisectParams) (*Bisect, error) {
	if p.Workers <= 0 || p.Workers > 512 {
		return nil, fmt.Errorf("invalid number of workers %d, worker count can't be less than 1 or more than 512", p.Workers)
	}

	return &Bisect{
		cmd:     p.Cmd,
		points:  p.Points,
		workers: p.Workers,
	}, nil
}

func (b *Bisect) run() error {
	fmt.Printf("Going to bisect using cmd: \"%s\", points: %v, workers: %d", b.cmd, b.points, b.workers)
	return nil
}
