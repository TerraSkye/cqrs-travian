package handlers

import (
	"context"
	"cqrs-travian/engine"
	"cqrs-travian/engine/domain"
	"cqrs-travian/engine/domain/commands"
	cqrs "github.com/terraskye/eventsourcing"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
	"os"
)

func init() {
	engine.RegisterCommand(func(aggregate *domain.Account) func(ctx context.Context, command *commands.CreateUserAccountCommand) error {
		tracer := otel.Tracer("engine")

		return func(ctx context.Context, cmd *commands.CreateUserAccountCommand) error {
			ctx, span := tracer.Start(ctx, "World:CreateUserAccount",
				trace.WithAttributes(
					// Add meta-related attributes
					attribute.String("cqrs.aggregate_id", cmd.AggregateID()),
					attribute.String("cqrs.aggregate_version", cqrs.MustExtractAggregateVersion(ctx)),
					attribute.String("cqrs.application", os.Getenv("application")),
					attribute.String("cqrs.causation_id", cqrs.MustExtractCausationId(ctx)),
					attribute.String("cqrs.correlation_id", trace.SpanContextFromContext(ctx).TraceID().String()),
					attribute.String("cqrs.command", "CreateUserAccount"),
					attribute.String("cqrs.function", "CreateUserAccount"),
					// Messaging attributes
					attribute.String("messaging.conversation_id", trace.SpanContextFromContext(ctx).TraceID().String()),
					attribute.String("messaging.destination", "CreateUserAccount"),
					attribute.String("messaging.destination_kind", "aggregate"),
					attribute.String("messaging.message_id", cqrs.MustExtractCausationId(ctx)),
					attribute.String("messaging.operation", "receive"),
					attribute.String("messaging.system", "cqrs"),
				),
			)
			defer span.End()
			err := aggregate.CreateUserAccount(ctx, cmd)
			if err != nil {
				span.RecordError(err)
				span.SetStatus(codes.Error, err.Error())
			} else {
				span.SetStatus(codes.Ok, "")
			}

			return err
		}
	})
}
