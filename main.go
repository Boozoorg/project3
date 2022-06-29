package main

import (
	"time"
)

type THEWROLD interface {
	Start(stopduration *stopwatch) *int64
	Stop(starttime *stopwatch) *bool
}

type stopwatch struct {
	Second int64
}

func (starttime *stopwatch) Start(stoptime *stopwatch) *int64 {
	a := int64(2)
	var result stopwatch
	for i := a - 1; i < a; i++ {
		print(i/3600, ":", i/60%60, ":", i%60)
		println()
		b := result.Stop(stoptime)
		if *b == true {
			return &i
		}
		a++
	}
	return &a
}

func (Timer *stopwatch) Stop(Stoptime *stopwatch) *bool {
	var a bool = true
	if Timer.Second == Stoptime.Second-1 {
		return &a
	}
	<-time.After(time.Second * 1)
	Timer.Second++
	a = false
	return &a
}

func main() {
	var StarPlatinum THEWROLD
	StartTime := &stopwatch{1}
	StopTime := &stopwatch{5}
	StarPlatinum = StartTime 
	StarPlatinum.Start(StopTime)
}
