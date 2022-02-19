package tested

import (
	"bytes"
	"reflect"
	"testing"
	"time"
)

func TestCountDown(t *testing.T) {
	t.Run("print 3 to Go!", func(t *testing.T) {
		buffer := &bytes.Buffer{}

		CountDown(buffer, &SpyCountdownOperations{})

		got := buffer.String()
		want := "3\n2\n1\nGo!"

		if got != want {
			t.Errorf("Got %q, but want %q", got, want)
		}

	})

	// test the sequence
	t.Run("sleep before every print", func(t *testing.T) {
		spySleepPrinter := &SpyCountdownOperations{}

		CountDown(spySleepPrinter, spySleepPrinter)

		want := []string{
			sleep, write, sleep, write, sleep, write, sleep, write,
		}

		if !reflect.DeepEqual(want, spySleepPrinter.Calls) {
			t.Errorf("wanted calls %v got %v", want, spySleepPrinter.Calls)
		}
	})
}

func TestConfigurableSleeper(t *testing.T) {
	sleepTime := 5 * time.Second

	spyTime := &SpyTime{}
	sleeper := ConfigurableSleeper{sleepTime, spyTime.Sleep}
	sleeper.Sleep()

	if spyTime.durationSlept != sleepTime {
		t.Errorf("should have slept %v but slept for %v", sleepTime, spyTime.durationSlept)
	}
}
