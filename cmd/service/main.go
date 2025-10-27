package main

import (
	"fmt"
	"melodia-events/internal/config"
	"melodia-events/internal/dependencies"

	"github.com/Melodia-IS2/melodia-go-utils/pkg/app"
	"github.com/Melodia-IS2/melodia-go-utils/pkg/http"
	"github.com/Melodia-IS2/melodia-go-utils/pkg/router"
)

func main() {
	cfg := config.Load()

	deps := dependencies.NewHandlerContainer(cfg)

	// Debugging purposes
	http.SetExposeErrorDetail(true)
	// End of debugging purposes

	builder, err := app.NewBuilder(&router.RouterConfig{
		ErrorHandler:            http.ErrorHandler,
		NotFoundHandler:         http.NotFoundHandler,
		MethodNotAllowedHandler: http.MethodNotAllowedHandler,
	}, cfg.Port)

	if err != nil {
		panic(err)
	}

	app := builder.
		RegisterHandler(deps.CreateEvent).
		RegisterHandler(deps.GetEvents).
		RegisterHandler(deps.Swagger).
		RegisterWorker(deps.Scheduler).
		Build()

	if err := app.Run(); err != nil {
		panic(fmt.Sprintf("failed to start application: %s", err.Error()))
	}
}
