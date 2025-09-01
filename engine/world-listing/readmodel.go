package world_listing

import (
	"context"
	"cqrs-travian/world/events"
	"github.com/google/uuid"
	cqrs "github.com/terraskye/eventsourcing"
)

// once it is installed, and loaded with the integration
var _ cqrs.ReadModel = (*ReadModel)(nil)

type Query struct {
}

func (q Query) ID() []byte { return nil }

type ReadModel struct {
	Version string
	Items   map[uuid.UUID]*WorldListing
}

func (p *ReadModel) OnWorldCreated(ctx context.Context, ev *events.WorldCreatedEvent) {
	p.Version = cqrs.MustExtractAggregateVersion(ctx)
	p.Items[ev.WorldID] = &WorldListing{
		ID:                ev.WorldID,
		Name:              ev.Name,
		Speed:             1,
		TotalPlayerCount:  0,
		OnlinePlayerCount: 0,
	}
}

type WorldListing struct {
	ID                uuid.UUID
	Name              string
	Speed             int8
	TotalPlayerCount  int64
	OnlinePlayerCount int64
}
