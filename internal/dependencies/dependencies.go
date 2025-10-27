package dependencies

import (
	"fmt"
	"melodia-events/internal/config"
	"melodia-events/internal/infrastructure/persistence"
	"melodia-events/internal/infrastructure/publishers"
	"melodia-events/internal/infrastructure/scheduler"
	"melodia-events/internal/swagger"
	"melodia-events/internal/usecase/createevent"
	"melodia-events/internal/usecase/getevents"

	"github.com/Melodia-IS2/melodia-go-utils/pkg/app"
	"github.com/Melodia-IS2/melodia-go-utils/pkg/router"
	"github.com/segmentio/kafka-go"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"

	_ "github.com/lib/pq"
)

// TODO ADD DEFERS
type HandlerContainer struct {
	CreateEvent router.CanRegister
	GetEvents   router.CanRegister
	Swagger     router.CanRegister
	Scheduler   app.Worker
}

func NewHandlerContainer(cfg *config.Config) *HandlerContainer {
	var client *mongo.Client
	var kafkaWriter *kafka.Writer

	clientOpts := options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%s@%s:%s/%s?authSource=admin",
		cfg.MongoConfig.User,
		cfg.MongoConfig.Password,
		cfg.MongoConfig.Host,
		cfg.MongoConfig.Port,
		cfg.MongoConfig.Database,
	))

	client, err := mongo.Connect(clientOpts)

	if err != nil {
		panic(err)
	}

	eventsDatabase := client.Database(cfg.MongoConfig.Database)
	eventsCollection := eventsDatabase.Collection("events")

	kafkaWriter = kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{cfg.KafkaURL},
	})

	print(cfg.KafkaURL)
	/* Repositories */

	eventRepo := &persistence.MongoEventRepository{
		Collection: eventsCollection,
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

	getEventsUC := &getevents.GetEventsImpl{
		EventRepository: eventRepo,
	}

	/* End of Usecases */

	/* Handlers */

	createEventHandl := &createevent.CreateEventHandler{
		CreateEventUC: createEventUC,
	}

	getEventsHandler := &getevents.GetEventsHandler{
		GetEventsUC: getEventsUC,
	}

	swaggerHandler := &swagger.SwaggerHandler{}

	/* End of Handlers */

	/* Workers */

	schedulerWorker := &scheduler.Worker{
		EventRepository: eventRepo,
		EventPublisher:  eventPublisher,
	}

	/* End of Scheduler */

	return &HandlerContainer{
		CreateEvent: createEventHandl,
		GetEvents:   getEventsHandler,
		Swagger:     swaggerHandler,
		Scheduler:   schedulerWorker,
	}
}
