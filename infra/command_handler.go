package infra

import (
	"context"
	"time"

	"github.com/sirupsen/logrus"
	cqrs "github.com/terraskye/eventsourcing"
)

type LoggingCommandHandler struct {
	inner  cqrs.CommandHandler
	logger *logrus.Entry
}

func NewLoggingCommandHandler(inner cqrs.CommandHandler, logger *logrus.Entry) *LoggingCommandHandler {
	return &LoggingCommandHandler{
		inner:  inner,
		logger: logger,
	}
}

func (l *LoggingCommandHandler) Handle(ctx context.Context, cmd cqrs.Command) error {
	causationID := cqrs.MustExtractCausationId(ctx)
	start := time.Now()

	l.logger.WithFields(logrus.Fields{
		"command":      cqrs.TypeName(cmd),
		"causation_id": causationID,
	}).Info("Handling command started")

	err := l.inner.Handle(ctx, cmd)

	duration := time.Since(start)

	entry := l.logger.WithFields(logrus.Fields{
		"command":      cqrs.TypeName(cmd),
		"causation_id": causationID,
		"duration":     duration.String(),
	})

	if err != nil {
		entry.WithField("error", err).Error("Handling command failed")
	} else {
		entry.Info("Handling command succeeded")
	}

	return err
}
