package gigasecond

import "time"

const (
	testVersion = 4
	giga        = 1e9
)

func AddGigasecond(in time.Time) time.Time {
	return in.Add(time.Second * giga)
}
