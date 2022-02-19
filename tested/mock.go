package tested

import (
	"fmt"
	"io"
	"time"
)

type Sleeper interface {
	Sleep()
}

type SpySleeper struct {
	Calls int
}

// For testing purpose
func (s *SpySleeper) Sleep() {
	s.Calls++
}

// For work
type DefaultSleeper struct{}

func (s *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

// Testing
type SpyCountdownOperations struct {
	Calls []string
}

func (s *SpyCountdownOperations) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *SpyCountdownOperations) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

const sleep = "sleep"
const write = "write"

func CountDown(buffer io.Writer, sleeper Sleeper) {
	for count := 3; count > 0; count-- {
		sleeper.Sleep()
		fmt.Fprintln(buffer, count)
	}
	sleeper.Sleep()
	fmt.Fprint(buffer, "Go!")
}

type SpyTime struct {
	durationSlept time.Duration
}

func (s *SpyTime) Sleep(duration time.Duration) {
	s.durationSlept = duration
}

type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

func (c *ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
}
