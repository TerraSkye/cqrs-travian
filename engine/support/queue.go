package support

import (
	"context"

	"github.com/google/uuid"
)

type Queue[T any] interface {
	Add(ctx context.Context, ticks int64, taskID uuid.UUID, task T)
	Poll(ctx context.Context)
	PollN(ctx context.Context, n int64)
	ExpireAll(ctx context.Context, expireFunction func(T) bool)
	RemoveByID(ctx context.Context, id uuid.UUID) bool
	Expired(ctx context.Context) []T
	Pending(ctx context.Context, filter func(T) bool) []T
}
