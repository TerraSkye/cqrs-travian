package building_queue

import (
	"context"
	"cqrs-travian/engine/support"

	"github.com/google/uuid"
	"github.com/io-da/query"
	cqrs "github.com/terraskye/eventsourcing"
)

// Query
type Query struct {
	VillageID uuid.UUID
}

func (q Query) ID() []byte {
	return []byte(q.VillageID.String())
}

type QueryHandlerList struct {
	queue support.Queue[PendingConstruction]
}

func NewQueryHandler() cqrs.GenericQueryHandler[Query, PendingConstruction] {
	return &QueryHandlerList{}
}

func (q *QueryHandlerList) HandleQuery(ctx context.Context, qry Query) (PendingConstruction, error) {
	q.queue.Expired(ctx)
}

func (q *QueryHandlerList) Handle(qry query.Query, res *query.IteratorResult) error {
	switch request := qry.(type) {
	case *Query:
		request.ProductID.Value()

		// fetch data

	}

	qry.ID()

	return nil
}
