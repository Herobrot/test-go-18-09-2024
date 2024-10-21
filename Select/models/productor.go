package models

import (
	"fmt"
	"time"
)

type Productor struct {
	flag bool
	ch   chan string
}

func NewProd(ch chan string) *Productor {
	return &Productor{
		flag: true,
		ch:   ch,
	}
}

func (p *Productor) Run() {
	i := 0
	for p.flag {
		p.ch <- fmt.Sprintf("P: %d", i)
		time.Sleep(1000)
		i++
	}
}

func (p *Productor) SetFlag(flag bool) {
	p.flag = flag
}
