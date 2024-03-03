package main

import (
	"time"
)

type Timer struct {
	GameObj *GameObj
	Duration time.Duration
	LastFrameTime time.Time
	Elapsed time.Duration
	Paused bool
	Callback func()
}

func (*Timer) Id() string {
	return "timer"
}

type TimerOption func(*Timer)

func (obj *GameObj) NewTimer(
	duration time.Duration,
	callback func(),
	opts ...TimerOption) *GameObj {

	timer := &Timer{
		GameObj: obj,
		Duration: duration,
		Callback: callback,
		Paused: true, // not ready to start at invoke
	}

	for _, opt := range opts {
		opt(timer)
	}

	obj.AddComponents(timer)

	return obj
}

func (t *Timer) Start() {
	
	t.Paused = false
	t.LastFrameTime = time.Now()
}

func (t *Timer) Reset() {
	t.Elapsed = 0
}


func (t *Timer) Update() {

	if t.Paused { return }

	now := time.Now()
	elapsed := now.Sub(t.LastFrameTime)
	t.LastFrameTime = now
	t.Elapsed += elapsed

	if t.Elapsed >= t.Duration {
		t.Callback()
	}

}

func (t *Timer) Draw() {
	// no op
}