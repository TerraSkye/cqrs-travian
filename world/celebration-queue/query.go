package celebration_queue

import "context"

type QueryHandler interface {
	HandleQuery(ctx context.Context, qry Query) ([]*ReadModel, string, error)
}

type queryHandler struct{}

func NewQueryHandler() QueryHandler {
	return &queryHandler{}
}

func (q *queryHandler) HandleQuery(ctx context.Context, qry Query) ([]*ReadModel, string, error) {

	completedTasks := tasks.Expired()

	return completedTasks, "", nil
}
