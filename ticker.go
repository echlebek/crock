package crock

import (
	"time"
)

// NewTicker creates a new ticker. It works like time.NewTicker, except
// it returns this package's Ticker. Since time.Ticker exposes a struct
// field, it is not possible to implement the time.Ticker interface.
func (t *Time) NewTicker(d time.Duration) *Ticker {
	ticker := newTicker(t, d)
	ticker.start()
	return ticker
}

func newTicker(t *Time, d time.Duration) *Ticker {
	return &Ticker{
		time:     t,
		duration: d,
		ch:       make(chan time.Time, 1),
	}
}

func (t *Ticker) start() {
	now := t.time.Now()
	f := new(func())
	*f = func() {
		now := t.time.Now()
		t.time.event(now.Add(t.duration), *f)
		t.ch <- now
	}
	t.time.event(now.Add(t.duration), *f)
}

func (t *Ticker) Stop() {
	// Intentionally blank
}

// Ticker is like time.Ticker, but uses a method for exposing its channel
// instead of a struct field. Ticker will tick only when time is progressing.
type Ticker struct {
	time     *Time
	duration time.Duration
	ch       chan time.Time
}

// C returns the ticker's channel.
func (t *Ticker) C() <-chan time.Time {
	return t.ch
}
