package counter

import "sync/atomic"

type Counter struct {
	Value atomic.Int64
}

func New() *Counter {
	return &Counter{}
}

func (c *Counter) Get() int64 {
	return c.Value.Load()
}

func (c *Counter) Inc() int64 {
	return c.Value.Add(1)
}

func (c *Counter) Set(newval int64) {
	c.Value.Store(newval)
}
