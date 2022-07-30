package clock

import (
	"context"
	"errors"
	"time"
)

var ErrInstanceNotFound = errors.New("instance not found in context")

type Clock interface {
	Now() time.Time
}

type contextKey struct{}

func Now(ctx context.Context) time.Time {
	clock := Instance(ctx)
	return clock.Now()
}

func Context(ctx context.Context, v Clock) context.Context {
	return context.WithValue(ctx, contextKey{}, v)
}

func Instance(ctx context.Context) Clock {
	clock, ok := ctx.Value(contextKey{}).(Clock)
	if !ok {
		panic(ErrInstanceNotFound)
	}

	return clock
}
