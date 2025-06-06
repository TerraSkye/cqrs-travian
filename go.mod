module cqrs-travian

go 1.24.1

replace github.com/terraskye/eventsourcing => ../github.com/terraskye/eventsourcing

require (
	github.com/google/uuid v1.6.0
	github.com/gorilla/mux v1.8.1
	github.com/terraskye/eventsourcing v0.0.0-20250325155942-bfb0ab12c18a
)

require (
	github.com/go-logr/logr v1.4.2 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/io-da/query v1.3.5 // indirect
	go.opentelemetry.io/auto/sdk v1.1.0 // indirect
	go.opentelemetry.io/otel v1.35.0 // indirect
	go.opentelemetry.io/otel/metric v1.35.0 // indirect
	go.opentelemetry.io/otel/trace v1.35.0 // indirect
)
