package crock

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestTimer(t *testing.T) {
	tm := NewTime(testTime)

	timer := tm.NewTimer(time.Minute)
	tm.Set(tm.Now().Add(time.Hour))
	currentTime := <-timer.C

	if got, want := currentTime, tm.Now(); !got.Equal(want) {
		t.Errorf("bad timer time: got %v, want %v", got, want)
	}

	canceled := timer.Stop()

	if got, want := canceled, false; got != want {
		t.Fatalf("bad timer cancel: got %v, want %v", got, want)
	}

}

func TestTimerStopBeforeFiring(t *testing.T) {
	tm := NewTime(testTime)

	timer := tm.NewTimer(time.Minute)

	canceled := timer.Stop()
	if got, want := canceled, true; got != want {
		t.Fatalf("bad timer cancel: got %v, want %v", got, want)
	}

	select {
	case <-timer.C:
		t.Fatal("received a value on timer.C")
	default:
	}
}

func TestResetTimer(t *testing.T) {
	tm := NewTime(testTime)

	timer := tm.NewTimer(time.Minute)
	tm.Set(tm.Now().Add(time.Hour))
	currentTime := <-timer.C
	if got, want := currentTime, tm.Now(); !got.Equal(want) {
		t.Errorf("bad currenTime: got %v, want %v", got, want)
	}

	timer.Reset(time.Hour)
	tm.Set(tm.Now().Add(2 * time.Hour))

	<-timer.C
}

func TestTimeAfterFunc(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(1)
	var called bool

	tm := NewTime(testTime)

	tm.AfterFunc(time.Minute, func() {
		fmt.Println("AfterFunc")
		called = true
		wg.Done()
	})
	tm.Set(tm.Now().Add(time.Hour))

	wg.Wait()

	if got, want := called, true; got != want {
		t.Fatal("AfterFunc never called")
	}
}
