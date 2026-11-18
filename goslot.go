package goslot

import (
	"math/rand"
	"time"
)

type Config struct {
	MinDelay   time.Duration
	MaxDelay   time.Duration
	StopTime   time.Duration
	NumReels   uint8
	NumSymbols uint
}

type GoSlot struct {
	minDelay   time.Duration
	maxDelay   time.Duration
	stopTime   time.Duration
	reels      []*reel
	isRunning  bool
	isStopping bool
	stopping   time.Time
}

func New(cfg Config) *GoSlot {
	reels := make([]*reel, cfg.NumReels)
	for i := range reels {
		reels[i] = newReel(cfg.NumSymbols)
	}

	return &GoSlot{
		minDelay:   cfg.MinDelay,
		maxDelay:   cfg.MaxDelay,
		stopTime:   cfg.StopTime,
		reels:      reels,
		isRunning:  false,
		isStopping: false,
	}
}

func (s *GoSlot) Start() {
	delayRange := int64(s.maxDelay - s.minDelay)
	stopTime := time.Duration(int(s.stopTime) / len(s.reels))
	s.isRunning = true
	for i, reel := range s.reels {
		go func() {
			for {
				delay := time.Duration(int64(s.minDelay) + rand.Int63n(delayRange))
				time.Sleep(delay)

				if !s.isStopping || time.Since(s.stopping) < stopTime*time.Duration(i) {
					reel.increment()
				}

				if !s.isRunning {
					break
				}
			}
		}()
	}
}

func (s *GoSlot) Stop() {
	go func() {
		s.isStopping = true
		s.stopping = time.Now()
		time.Sleep(s.stopTime)
		s.isRunning = false
		s.isStopping = false
	}()
}

func (s *GoSlot) Values() []uint {
	values := make([]uint, len(s.reels))
	for i := range s.reels {
		values[i] = s.reels[i].getValue()
	}
	return values
}

func (s *GoSlot) IsRunning() bool {
	return s.isRunning
}
