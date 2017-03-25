package clock

import "fmt"

type Clock struct {
	minutes int
}

const (
	testVersion = 4
	hour        = 60
	day         = 24 * hour
)

func New(hours int, minutes int) Clock {
	return Clock{}.Add(hours*hour + minutes)
}

func (clock Clock) String() string {
	return fmt.Sprintf("%02d:%02d", clock.minutes/hour, clock.minutes%hour)
}

func (clock Clock) Add(minutes int) Clock {
	clock.minutes = (clock.minutes + day + minutes%day) % day
	return clock
}
