package models

import (
	"fmt"
	"time"
)

type Consumidor struct {
	flag bool
	ch   chan string
}

func NewCon(ch chan string) *Consumidor {
	return &Consumidor{
		flag: true,
		ch:   ch,
	}
}

func (c *Consumidor) Run() {
	i := 0
	for c.flag {
		c.ch <- `Consumidor - \{i}`
		time.Sleep(1000)
		fmt.Sprintf("C %d", i)
		i++
	}
}

func (c *Consumidor) SetFlag(flag bool) {
	c.flag = flag
}
