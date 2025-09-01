package domain

import (
	"cqrs-travian/engine/support"
	cqrs "github.com/terraskye/eventsourcing"
)

type Player struct {
	*cqrs.AggregateBase
	// current cold
	gold  int64
	tribe support.Tribe
}
