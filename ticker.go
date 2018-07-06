package crock

import (
	"runtime"
	"sync/atomic"
	"time"
)

// NewTicker creates a new time.Ticker. It works like time.NewTicker, except
// that is only ticks when time is progressing. Because the ticker's channel
// is co-opted by crock, the Stop() method has no effect. However, ticks will
// stop being generated when the object becomes unreachable.
func (t *Time) NewTicker(d time.Duration) *time.Ticker {
	crockTicker := newTicker(t, d)
	// start guarantees that crockTicker will be reachable until stop is called
	crockTicker.start()
	ticker := &time.Ticker{C: crockTicker.ch}
	// Make sure we don't keep creating tick events after the ticker has gone
	// out of scope.
	runtime.SetFinalizer(ticker, func(interface{}) { crockTicker.stop() })
	return ticker
}

func newTicker(t *Time, d time.Duration) *crockTicker {
	return &crockTicker{
		time:     t,
		duration: d,
		ch:       make(chan time.Time, 1),
	}
}

func (t *crockTicker) start() {
	t.running = 1
	now := t.time.Now()
	f := new(func())
	// use a pointer to have this closure add itself to time events.
	// the magic of indirection!
	*f = func() {
		now := t.time.Now()
		if atomic.LoadInt64(&t.running) == 1 {
			t.time.event(now.Add(t.duration), *f)
		}
		t.ch <- now
	}
	t.time.event(now.Add(t.duration), *f)
}

func (t *crockTicker) stop() {
	atomic.StoreInt64(&t.running, 0)
}

// crockTicker is like time.Ticker, but will tick only when time is progressing
type crockTicker struct {
	running  int64
	time     *Time
	duration time.Duration
	ch       chan time.Time
}
