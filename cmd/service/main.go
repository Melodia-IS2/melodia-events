package main

import (
	"fmt"

	"github.com/Melodia-IS2/melodia-events/internal/config"
	"github.com/Melodia-IS2/melodia-events/internal/dependencies"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/rs/cors"

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
		RegisterMiddleware(middleware.Logger).
		RegisterMiddleware(middleware.Recoverer).
		RegisterMiddleware(cors.AllowAll().Handler).
		RegisterHandler(deps.CreateEvent).
		RegisterHandler(deps.GetEvents).
		RegisterHandler(deps.CreateLog).
		RegisterHandler(deps.GetLogs).
		RegisterHandler(deps.Swagger).
		RegisterWorker(deps.Scheduler).
		RegisterConsumer(deps.ConsumerUserDevices).
		Build()

	if err := app.Run(); err != nil {
		panic(fmt.Sprintf("failed to start application: %s", err.Error()))
	}
}
