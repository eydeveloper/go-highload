package main

import (
	"fmt"
	"os"

	social "github.com/eydeveloper/highload-social"
	"github.com/eydeveloper/highload-social/internal/handler"
	"github.com/eydeveloper/highload-social/internal/repository"
	"github.com/eydeveloper/highload-social/internal/service"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := initConfig(); err != nil {
		logrus.Fatalf("error initializing config: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error loading env variables: %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})

	if err != nil {
		logrus.Fatalf("failed to initialize db: %s", err.Error())
	}

	dbSlave, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db-slave.host"),
		Port:     viper.GetString("db-slave.port"),
		Username: viper.GetString("db-slave.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   viper.GetString("db-slave.dbname"),
		SSLMode:  viper.GetString("db-slave.sslmode"),
	})

	if err != nil {
		logrus.Fatalf("failed to initialize db: %s", err.Error())
	}

	redisClient := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", viper.GetString("redis.host"), viper.GetString("redis.port")),
		Password: "",
		DB:       0,
	})

	defer redisClient.Close()

	amqpConnection, err := amqp.Dial("amqp://guest:guest@localhost:5672/")

	if err != nil {
		logrus.Fatalf("failed to connect to the amqp connection: %s", err.Error())
	}

	defer amqpConnection.Close()

	amqpChannel, err := amqpConnection.Channel()

	if err != nil {
		logrus.Fatalf("failed to connect to the amqp channel: %s", err.Error())
	}

	defer amqpChannel.Close()

	err = amqpChannel.ExchangeDeclare(
		"feed",
		"direct",
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		logrus.Fatalf("failed to declare the posts queue: %s", err.Error())
	}

	repositories := repository.NewRepository(db, dbSlave)
	services := service.NewService(repositories, redisClient, amqpChannel)
	handlers := handler.NewHandler(services)

	err = new(social.Server).Run(viper.GetString("port"), handlers.InitRoutes())

	if err != nil {
		logrus.Fatalf("error occured while running http server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("config")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
