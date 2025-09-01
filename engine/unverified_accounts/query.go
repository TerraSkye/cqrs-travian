package unverified_accounts

import (
	"context"

	"github.com/google/uuid"
	cqrs "github.com/terraskye/eventsourcing"
)

type QueryHandler struct {
	store cqrs.EventStore
}

func NewQueryHandler(store cqrs.EventStore) cqrs.GenericQueryHandler[Query, ReadModel] {
	return &QueryHandler{store: store}
}

func (q *QueryHandler) HandleQuery(ctx context.Context, _ Query) (ReadModel, error) {

	events, err := q.store.LoadFromAll(ctx, 0)

	if err != nil {
		return ReadModel{}, err
	}

	model := ReadModel{
		Version: "0",
		Items:   make(map[uuid.UUID]*WorldListing),
	}

	hydrate := cqrs.Hydrate(
		cqrs.NewHydrateHandler(model.OnWorldCreated),
	)

	for event := range events {
		hydrate(ctx, event.Event)
	}

	return model, nil

}
