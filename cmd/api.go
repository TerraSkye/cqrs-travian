package cmd

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"time"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	_ "github.com/terraskye/awesomeProject2/express/events"
	"github.com/terraskye/awesomeProject2/express/slices/change_stock"
	parts_to_model "github.com/terraskye/awesomeProject2/express/slices/parts-to-model"
	"github.com/terraskye/awesomeProject2/express/slices/tick"
	cqrs "github.com/terraskye/eventsourcing"
	memorybus "github.com/terraskye/eventsourcing/eventbus/memory"
	"github.com/terraskye/eventsourcing/eventstore/memory"
	tsotel "github.com/terraskye/eventsourcing/otel"
)

type Service interface {
	Run(cmd *cobra.Command, args []string)
}

type apiService struct {
	logger *logrus.Entry
}

func NewApiService(log *logrus.Entry) Service {
	return apiService{
		logger: log,
	}
}

func (s apiService) Run(cmd *cobra.Command, args []string) {
	s.logger.Info("service started")
	defer func() {
		s.logger.Info("service stopped")
	}()

	httptest.DefaultRemoteAddr
	var commandBus = cqrs.NewCommandBus(100, 10)

	var eventbus *memorybus.EventBus

	{
		eventbus = memorybus.NewEventBus(1000)

		go func() {
			for err := range eventbus.Errors() {
				s.logger.WithError(err).Error("eventbus does not implement Dispatch()")
			}
		}()

		if err := eventbus.Subscribe(cmd.Context(), "logger",
			cqrs.NewEventHandlerFunc(func(ctx context.Context, event cqrs.Event) error {
				fmt.Println(cqrs.StreamIDFromContext(ctx), cqrs.VersionFromContext(ctx), cqrs.TypeName(event))
				return nil
			})); err != nil {
			s.logger.WithError(err).Error("failed to subscribe to logger")
		}

	}

	var eventstore cqrs.EventStore

	{
		//fs, err := filestore.NewFileStore("./.store")

		//if err != nil {
		//	s.logger.Fatal(err)
		//}

		fs := memory.NewMemoryStore(500)

		go func() {
			for envelope := range fs.Events() {
				eventbus.Dispatch(envelope)
			}
		}()

		eventstore = fs
	}

	var router = mux.NewRouter()

	{
		eventbus = memorybus.NewEventBus(1000)

	}

	{
		parts_to_model.MakeHttpHandler(router)

		//projecor := parts_to_model.NewProjector()

		eventbus.Subscribe(cmd.Context(), "logger", cqrs.NewEventHandlerFunc(func(ctx context.Context, event cqrs.Event) error {
			fmt.Println(cqrs.StreamIDFromContext(ctx), cqrs.VersionFromContext(ctx))
			return nil
		}))

	}

	//queue := infra.NewUnboundedQueue[string]()
	//
	//cqrs.OnEvent(func(ctx context.Context, ev events.ProductAdded) error {
	//	queue.Enqueue(ev.Sku)
	//	return nil
	//})

	{

		cqrs.Register(commandBus,
			tsotel.WithCommandTelemetry(change_stock.NewCommandHandler(eventstore)),
		)
	}

	{
		cqrs.Register(commandBus, tsotel.WithCommandTelemetry(tick.NewCommandHandler(eventstore)))
	}

	go func() {
		ticker := time.NewTicker(1 * time.Millisecond)

		for t := range ticker.C {

			go func() {
				res, err := commandBus.Dispatch(context.Background(), tick.Command{Name: t.String()})
				fmt.Println(res, err)
			}()

		}
	}()
	//if err := import_inventory.NewService(commandBus).Start(cmd.Context()); err != nil {
	//	s.logger.WithError(err).Error("failed to start import inventory")
	//}

	http.Handle("/", router)

	errs := make(chan error, 2)

	go func() {
		httpAddr := ":9090"
		fmt.Println("serving on 0.0.0.0:9090")
		errs <- http.ListenAndServe(httpAddr, nil)
	}()

	<-cmd.Context().Done()

}
