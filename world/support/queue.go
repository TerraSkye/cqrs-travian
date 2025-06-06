package support

import (
	"context"
	"github.com/google/uuid"
)

type Queue[T any] interface {
	Add(ctx context.Context, ticks int64, askID uuid.UUID, task T)
	Poll(ctx context.Context) []T
	PollN(ctx context.Context, n int64) []T
	ExpireAll(ctx context.Context, expireFunction func(T) bool) []T
	Expired(ctx context.Context) []T
	RemoveByID(ctx context.Context, id uuid.UUID) bool
}
