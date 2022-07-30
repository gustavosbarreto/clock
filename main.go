package main

import (
	"context"
	"fmt"
	"time"

	"github.com/gustavosbarreto/clock/clock"
)

type utcClock struct {
}

func (c *utcClock) Now() time.Time {
	return time.Now().UTC()
}

type localClock struct {
}

func (c *localClock) Now() time.Time {
	return time.Now()
}

type timezoneClock struct {
	location string
}

func (c *timezoneClock) Now() time.Time {
	loc, _ := time.LoadLocation(c.location)
	return time.Now().In(loc)
}

var _ clock.Clock = (*utcClock)(nil)
var _ clock.Clock = (*localClock)(nil)
var _ clock.Clock = (*timezoneClock)(nil)

func printTimeNow(ctx context.Context) {
	fmt.Println(clock.Now(ctx))
}

func main() {
	// utc clock
	utc := clock.Context(context.Background(), &utcClock{})
	// local system clock
	local := clock.Context(context.Background(), &localClock{})
	// timezone clock
	timezone := clock.Context(context.Background(), &timezoneClock{location: "America/Los_Angeles"})

	printTimeNow(utc)
	printTimeNow(local)
	printTimeNow(timezone)
}
