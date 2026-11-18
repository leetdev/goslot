package goslot

import "sync"

type reel struct {
	mu        sync.Mutex
	numValues uint
	value     uint
}

func newReel(numValues uint) *reel {
	return &reel{
		numValues: numValues,
		value:     0,
	}
}

func (r *reel) increment() {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.value++
	if r.value >= r.numValues {
		r.value = 0
	}
}

func (r *reel) getValue() uint {
	r.mu.Lock()
	defer r.mu.Unlock()
	return r.value
}
