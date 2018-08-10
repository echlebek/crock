package crock

import (
	"testing"
	"time"
)

var testTime time.Time

func init() {
	var err error
	testTime, err = time.Parse("Mon Jan 2 15:04:05 -0700 MST 2006", "Mon Jan 2 15:04:05 -0700 MST 2006")
	if err != nil {
		panic(err)
	}
}

func TestNewTime(t *testing.T) {
	tm := NewTime(testTime)
	if got, want := tm.Now(), testTime; !got.Equal(want) {
		t.Errorf("bad init time: got %v, want %v", got, want)
	}
	if got, want := tm.Resolution, DefaultResolution; got != want {
		t.Errorf("bad init resolution: got %v, want %v", got, want)
	}
	if got, want := tm.Multiplier, DefaultMultiplier; got != want {
		t.Errorf("bad init multiplier: got %v, want %v", got, want)
	}
	if tm.done != nil {
		t.Errorf("bad init done: got %v, want nil", tm.done)
	}
	if tm.events == nil {
		t.Error("bad init events: got nil")
	}
}

func TestTimeNow(t *testing.T) {
	tm := NewTime(testTime)
	if got, want := tm.Now(), testTime; !got.Equal(want) {
		t.Errorf("bad Now(); got %v, want %v", got, want)
	}
	then := tm.Now().Add(time.Duration(time.Second))
	tm.Set(then)
	if got, want := tm.Now(), then; got != want {
		t.Errorf("bad Now(): got %v, want %v", got, want)
	}
}

func TestTimeStartStopNow(t *testing.T) {
	tm := NewTime(testTime)
	tm.Resolution = time.Nanosecond
	tm.Multiplier = 100000.0
	tm.Start()
	defer tm.Stop()
	time.Sleep(time.Millisecond)
	if got, want := tm.Now(), testTime; !want.Before(got) {
		t.Errorf("bad time: want (%v) is not before got (%v)", want, got)
	}
}

func TestTimeSleep(t *testing.T) {
	tm := NewTime(testTime)
	tm.Multiplier = 100000.0
	// Times out if it's broken :P
	tm.Sleep(time.Hour)
}

func TestTimeAfter(t *testing.T) {
	tm := NewTime(testTime)
	ch := tm.After(time.Nanosecond)
	select {
	case <-ch:
		t.Fatal("shouldn't have seen a channel send")
	default:
	}
	tm.Set(tm.Now().Add(time.Nanosecond))
	// Times out if it's broken :P
	<-ch
}

func TestTimeAfterTimeNotFlowing(t *testing.T) {
	tm := NewTime(testTime)
	tm.Set(tm.Now().Add(time.Hour))
	ch := tm.After(-time.Minute)
	// Times out if it's broken :P
	<-ch
}
