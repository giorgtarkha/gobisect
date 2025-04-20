package main

import (
	"fmt"
)

type BisectInterface interface {
	Run() error

	reCalc()
}

type Bisect struct {
	cmd             string
	points          []string
	workers         int
	moreWeightRight bool

	leftBound  int
	rightBound int
	checkOrder []int
}

type BisectParams struct {
	Cmd             string
	Points          []string
	Workers         int
	MoreWeightRight bool
}

func NewBisect(p *BisectParams) (*Bisect, error) {
	if p.Workers <= 0 || p.Workers > 512 {
		return nil, fmt.Errorf("invalid number of workers %d, worker count can't be less than 1 or more than 512", p.Workers)
	}

	return &Bisect{
		cmd:             p.Cmd,
		points:          p.Points,
		workers:         p.Workers,
		moreWeightRight: p.MoreWeightRight,
		leftBound:       0,
		rightBound:      len(p.Points) - 1,
	}, nil
}

func (b *Bisect) Run() error {
	fmt.Printf("Going to bisect using cmd: \"%s\", points: %v, workers: %d\n", b.cmd, b.points, b.workers)

	b.reCalc()

	fmt.Printf("%v\n", b.checkOrder)

	return nil
}

func (b *Bisect) reCalc() {
	idx := 0
	b.checkOrder = make([]int, b.rightBound-b.leftBound+1)

	o := 0
	if b.moreWeightRight {
		o = 1
	}

	var f func(int, int)
	f = func(l, r int) {
		if l > r {
			return
		}
		m := l + (r-l+o)/2
		b.checkOrder[idx] = m
		idx++
		if b.moreWeightRight {
			f(m+1, r)
			f(l, m-1)
		} else {
			f(l, m-1)
			f(m+1, r)
		}
	}

	f(b.leftBound, b.rightBound)
}
