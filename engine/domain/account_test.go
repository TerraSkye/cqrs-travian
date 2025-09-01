package domain

import (
	"context"
	"cqrs-travian/world"
	"cqrs-travian/world/domain/commands"
	"cqrs-travian/world/events"
	"github.com/google/uuid"
	cqrs "github.com/terraskye/eventsourcing"
	"testing"
)

func TestAccount_CreateUserAccount(t *testing.T) {

	world1 := uuid.New()

	tests := []struct {
		name           string
		history        []cqrs.Event
		command        cqrs.Command
		wantErr        bool
		expectedEvents []cqrs.Event
	}{
		{
			name:    "Happy path",
			history: []cqrs.Event{},
			command: &commands.CreateUserAccountCommand{
				WorldID:  world1,
				Username: "test",
				Email:    "test@example.org",
				Password: "test123",
			},
			wantErr:        false,
			expectedEvents: nil,
		},
		{
			name: "duplicate registration",
			history: []cqrs.Event{
				events.UserAccountCreatedEvent{
					WorldID:  world1,
					Username: "test",
					Email:    "test@example.org",
					Password: "test123",
				},
			},
			command: &commands.CreateUserAccountCommand{
				WorldID:  world1,
				Username: "test",
				Email:    "test@example.org",
				Password: "test123",
			},
			wantErr:        true,
			expectedEvents: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			aggregate, err := world.AggregateForCommand(tt.command)

			if err != nil {
				t.Errorf("expected an aggregate but got: %v", err)
			}

			for _, event := range tt.history {
				if err := world.DispatchEvent(aggregate, event); err != nil {
					t.Errorf("failed to appy event `%v` onto aggregate %v, got err %v", aggregate, event, err)
				}
			}

			if err := world.DispatchCommand(context.Background(), aggregate, tt.command); (err != nil) != tt.wantErr {
				t.Errorf("CreateUserAccount() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
