package lib

import rl "github.com/gen2brain/raylib-go/raylib"

type Timer struct {
	startTime float64
	lifeTime  float64
}

func StartTimer(lifeTime float64) *Timer {
	timer := Timer{}

	timer.startTime = rl.GetTime()
	timer.lifeTime = lifeTime

	return &timer
}

func (timer Timer) IsTimerDone() bool {
	return rl.GetTime()-timer.startTime >= timer.lifeTime
}

func (timer Timer) GetElapsed() float64 {
	return rl.GetTime() - timer.startTime
}

func (timer *Timer) ResetTimer(lifeTime float64) {
	timer.startTime = rl.GetTime()
	timer.lifeTime = lifeTime
}
