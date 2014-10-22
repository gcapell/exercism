package gigasecond

import "time"

const Gigasecond = time.Second * 1e9

func AddGigasecond(t time.Time) time.Time {
	return t.Add(Gigasecond)
}
