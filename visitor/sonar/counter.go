package sonar

import (
	"sync/atomic"

	"github.com/gsixo/gocognit/visitor"
)

func NewComplexityCounter() visitor.Counter {
	return &ComplexityCounter{
		counter: atomic.Int64{},
	}
}

type ComplexityCounter struct {
	counter atomic.Int64
}

func (c *ComplexityCounter) Inc(delta uint64) {
	c.counter.Add(int64(delta))
}

func (c *ComplexityCounter) Dec() {
	c.counter.Add(-1)
}

func (c *ComplexityCounter) Load() uint64 {
	return uint64(c.counter.Load())
}

func NewNestingCounter() visitor.Counter {
	return &NestingCounter{
		counter: atomic.Int64{},
	}
}

type NestingCounter struct {
	counter atomic.Int64
}

func (c *NestingCounter) Inc(delta uint64) {
	c.counter.Add(int64(delta))
}

func (c *NestingCounter) Dec() {
	c.counter.Add(-1)
}

func (c *NestingCounter) Load() uint64 {
	return uint64(c.counter.Load())
}
