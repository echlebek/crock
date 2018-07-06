package crock

import (
	"testing"
	"time"
)

func TestCrockTicker(t *testing.T) {
	tm := NewTime(testTime)
	tm.Resolution = time.Millisecond
	ticker := tm.NewTicker(time.Millisecond)
	defer ticker.Stop()
	ch := ticker.C
	time.Sleep(time.Millisecond * 2)
	select {
	case <-ch:
		t.Fatal("should not have received an event")
	default:
	}
	tm.Start()
	defer tm.Stop()
	for i := 0; i < 10; i++ {
		tickTime := <-ch
		if got, want := tickTime, testTime.Add(time.Millisecond*time.Duration(i+1)); !got.Equal(want) {
			t.Errorf("bad tick time: got %v, want %v", got, want)
		}
	}
}

func TestTick(t *testing.T) {
	tm := NewTime(testTime)
	tm.Resolution = time.Millisecond
	tm.Start()
	defer tm.Stop()
	ch := tm.Tick(time.Millisecond)
	tickTime := <-ch
	if got, want := tickTime, testTime.Add(time.Millisecond); !got.Equal(want) {
		t.Errorf("bad tick time: got %v, want %v", got, want)
	}
}
