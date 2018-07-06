package main

import (
	"fmt"
	gotime "time"
)

type RealTime struct {
}

func (RealTime) Now() gotime.Time {
	return gotime.Now()
}

type Time interface {
	Now() gotime.Time
}

var time Time = RealTime{}

func main() {
	fmt.Println(Cry())
}

func Cry() string {
	when := time.Now().Format(gotime.Kitchen)
	return fmt.Sprintf("It's %v and all's well!\n", when)
}
