package main

import (
	"testing"
	gotime "time"

	"github.com/echlebek/crock"
)

func init() {
	var testTime = gotime.Unix(0, 0)
	time = crock.NewTime(testTime)
}

func TestCry(t *testing.T) {
	got := Cry()
	want := "It's 4:00PM and all's well!\n"
	if got != want {
		t.Errorf("bad cry: got %q, want %q", got, want)
	}
}
