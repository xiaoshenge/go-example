package main

import (
	"sync"

	"github.com/gomodule/redigo/redis"
)

type pool struct {
	Dial   func() (redis.Conn, error)
	active int
	close  bool
	mu     sync.Mutex
}

func (p *pool) Get() (redis.Conn, error) {
	p.mu.Lock()
	p.active++
	p.mu.Unlock()
	c, err := p.Dial()
	if err != nil {
		p.mu.Lock()
		p.active--
		p.mu.Unlock()
		c = nil
	}
	return c, err
}
func (p *pool) Close() error {
	p.mu.Lock()
	if p.close {
		p.mu.Unlock()
		return nil
	}
	p.close = true
	return nil
}
