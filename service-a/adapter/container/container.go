package container

import (
	"github.com/rabbitmq/amqp091-go"
	user2 "github/herochi/orbi/service-a/adapter/grpc/user"
	"github/herochi/orbi/service-a/application/notifier"
	"github/herochi/orbi/service-a/application/user"
	"github/herochi/orbi/service-a/infrastructure/datastore"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
	"log"
)

const serviceANotifies = "service-a-notifies"

type Container struct {
	DB      *mongo.Database
	Service Service
}

type Service struct {
	UserService user.UserService
}

func Inject() *Container {

	db, err := datastore.NewMongoDB()
	if err != nil {
		log.Fatalf("Error al conectar mongo: %v", err)
	}

	conn, err := grpc.Dial("service-b:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Error al conectar: %v", err)
	}

	connRabbit, err := amqp091.Dial("amqp://guest:guest@rabbitmq:5672/")
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}

	ch, err := connRabbit.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %v", err)
	}

	_, err = ch.QueueDeclare(
		serviceANotifies,
		false,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		log.Fatalf("Failed to declare queue: %v", err)
	}

	grpcS := user2.NewUserServiceClient(conn)

	services := injectServices(db, grpcS, ch)
	return &Container{
		DB:      db,
		Service: services,
	}
}

func injectNotifierService(rabbitMQChannel *amqp091.Channel) notifier.Notifier {
	q := serviceANotifies
	return notifier.NewNotifier(rabbitMQChannel, &q)
}

func injectUserService(db *mongo.Database, userGRPC user2.UserServiceClient, notifier notifier.Notifier) user.UserService {
	userRepository := user.NewUserRepository(db)
	userPresenter := user.NewUserPresenter()
	userService := user.NewUserService(userRepository, userPresenter, userGRPC, notifier)
	return userService
}

func injectServices(db *mongo.Database, userGRPC user2.UserServiceClient, rabbitMQChannel *amqp091.Channel) Service {

	notifierService := injectNotifierService(rabbitMQChannel)

	services := Service{
		UserService: injectUserService(db, userGRPC, notifierService),
	}
	return services
}
