package dependencies

import (
	"database/sql"
	"melodia-events/internal/config"
	"melodia-events/internal/infrastructure/persistence"
	"melodia-events/internal/infrastructure/publishers"
	"melodia-events/internal/swagger"
	"melodia-events/internal/usecase/createevent"

	"github.com/Melodia-IS2/melodia-go-utils/pkg/router"
	"github.com/segmentio/kafka-go"

	_ "github.com/lib/pq"
)

type deferred struct {
	db          *sql.DB
	kafkaWriter *kafka.Writer
}

func (d *deferred) Close() {
	d.db.Close()
	d.kafkaWriter.Close()
}

type HandlerContainer struct {
	CreateEvent router.CanRegister
	Swagger     router.CanRegister
	Deferred    *deferred
}

func NewHandlerContainer(cfg *config.Config) *HandlerContainer {
	var db *sql.DB
	var kafkaWriter *kafka.Writer

	defer func() {
		r := recover()
		if r != nil {
			if db != nil {
				db.Close()
			}
			if kafkaWriter != nil {
				kafkaWriter.Close()
			}
		}
		panic(r)
	}()

	db, err := sql.Open("postgres", cfg.DatabaseURL)

	if err != nil {
		panic(err)
	}

	kafkaWriter = kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{cfg.KafkaURL},
	})

	deferred := &deferred{
		db:          db,
		kafkaWriter: kafkaWriter,
	}

	/* Repositories */

	eventRepo := &persistence.PostgresEventRepository{
		DB: db,
	}

	eventPublisher := &publishers.KafkaEventPublisher{
		Writer: kafkaWriter,
	}
	/* End of Repositories */

	/* Usecases */

	createEventUC := &createevent.CreateEventImpl{
		EventRepository: eventRepo,
		EventPublisher:  eventPublisher,
	}

	/* End of Usecases */

	/* Handlers */

	createEventHandl := &createevent.CreateEventHandler{
		CreateEventUC: createEventUC,
	}

	swaggerHandler := &swagger.SwaggerHandler{}

	/* End of Handlers */

	return &HandlerContainer{
		CreateEvent: createEventHandl,
		Swagger:     swaggerHandler,
		Deferred:    deferred,
	}
}
