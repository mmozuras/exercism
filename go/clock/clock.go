package clock

import "fmt"

type Clock int

const (
	testVersion = 4
	hour        = 60
	day         = 24 * hour
)

func New(hours int, minutes int) Clock {
	return Clock(0).Add(hours*hour + minutes)
}

func (clock Clock) String() string {
	return fmt.Sprintf("%02d:%02d", clock/hour, clock%hour)
}

func (clock Clock) Add(minutes int) Clock {
	clock = Clock((int(clock) + day + minutes%day) % day)
	return clock
}
