package gigasecond

import "time"

const (
	testVersion = 4
	giga        = 1000000000
)

func AddGigasecond(in time.Time) time.Time {
	return in.Add(time.Second * giga)
}
