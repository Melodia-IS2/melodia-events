package dependencies

import (
	"fmt"
	"time"

	"github.com/Melodia-IS2/melodia-events/internal/config"
	kafkahelper "github.com/Melodia-IS2/melodia-events/internal/infrastructure/kafka"
	"github.com/Melodia-IS2/melodia-events/internal/infrastructure/persistence"
	"github.com/Melodia-IS2/melodia-events/internal/infrastructure/publishers"
	"github.com/Melodia-IS2/melodia-events/internal/infrastructure/scheduler"
	"github.com/Melodia-IS2/melodia-events/internal/swagger"
	"github.com/Melodia-IS2/melodia-events/internal/usecase/consumerdevices"
	"github.com/Melodia-IS2/melodia-events/internal/usecase/createevent"
	"github.com/Melodia-IS2/melodia-events/internal/usecase/createlog"
	"github.com/Melodia-IS2/melodia-events/internal/usecase/getevents"
	"github.com/Melodia-IS2/melodia-events/internal/usecase/getlogs"
	kafkaconsumer "github.com/Melodia-IS2/melodia-events/pkg/suscriber/kafka"
	"github.com/Melodia-IS2/melodia-go-utils/pkg/app"
	"github.com/Melodia-IS2/melodia-go-utils/pkg/router"
	"github.com/redis/go-redis/v9"
	"github.com/segmentio/kafka-go"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"

	_ "github.com/lib/pq"
)

// TODO ADD DEFERS
type HandlerContainer struct {
	CreateEvent         router.CanRegister
	GetEvents           router.CanRegister
	CreateLog           router.CanRegister
	Swagger             router.CanRegister
	GetLogs             router.CanRegister
	ConsumerUserDevices kafkaconsumer.Consumer
	Scheduler           app.Worker
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

	mongoDatabase := client.Database(cfg.MongoConfig.Database)
	logsCollection := mongoDatabase.Collection("logs")

	if err := kafkahelper.WaitForKafka(cfg.KafkaURL, 30*time.Second); err != nil {
		panic(fmt.Errorf("Kafka not available: %w", err))
	}

	if err := kafkahelper.EnsureTopics(cfg.KafkaURL, cfg.KafkaTopics); err != nil {
		fmt.Printf("Error creating topics: %v\n", err)
	}

	_ = kafkahelper.ListTopics(cfg.KafkaURL)

	kafkaWriter = kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{cfg.KafkaURL},
	})

	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", cfg.RedisConfig.Host, cfg.RedisConfig.Port),
		Password: cfg.RedisConfig.Password,
		DB:       0,
	})

	/* Repositories */

	eventRepo := &persistence.RedisEventRepository{
		Rdb: rdb,
		Key: "events",
	}

	deviceRepo := &persistence.RedisDevicesRepository{
		Rdb: rdb,
		Key: "devices",
	}

	logRepo := &persistence.MongoLogRepository{
		Collection: logsCollection,
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

	createLogUC := &createlog.CreateLogImpl{
		LogRepository: logRepo,
	}

	getLogsUC := &getlogs.GetLogsImpl{
		LogRepository: logRepo,
	}
	/* End of Usecases */

	/* Handlers */

	createEventHandl := &createevent.CreateEventHandler{
		CreateEventUC: createEventUC,
	}

	getEventsHandler := &getevents.GetEventsHandler{
		GetEventsUC: getEventsUC,
	}

	createLogHandler := &createlog.CreateLogHandler{
		CreateLogUC: createLogUC,
	}

	swaggerHandler := &swagger.SwaggerHandler{}

	/* End of Handlers */

	/* Workers */

	schedulerWorker := &scheduler.Worker{
		EventRepository: eventRepo,
		EventPublisher:  eventPublisher,
	}

	getLogsHandler := &getlogs.GetLogsHandler{
		GetLogsUC: getLogsUC,
	}

	/* End of Scheduler */

	consumerUserDevices := consumerdevices.ConsumerUserDevices{
		DeviceRepository: deviceRepo,
	}

	ConsumerUserDevices := kafkaconsumer.NewBatchConsumer(kafkaconsumer.BatchConfig{
		Brokers:      []string{cfg.KafkaURL},
		GroupID:      cfg.AppName,
		Topic:        "user_update",
		BatchSize:    1000,
		BatchTimeout: 10 * time.Second,
	}, consumerUserDevices.ConsumeBatch)

	return &HandlerContainer{
		CreateEvent:         createEventHandl,
		GetEvents:           getEventsHandler,
		CreateLog:           createLogHandler,
		Swagger:             swaggerHandler,
		GetLogs:             getLogsHandler,
		Scheduler:           schedulerWorker,
		ConsumerUserDevices: ConsumerUserDevices,
	}
}
