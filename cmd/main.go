package main

import (
	"cashbox"
	"cashbox/pkg/handler"
	"cashbox/pkg/repository"
	"cashbox/pkg/service"
	"context"
	"fmt"
	"github.com/gofiber/adaptor/v2"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	logger := zap.NewExample() //Change to zap.NewProduction()

	if err := InitConfig(); err != nil {
		logger.Fatal(fmt.Sprintf("error initializing configs: %s", err.Error()))
	}

	if err := godotenv.Load(); err != nil {
		logger.Fatal(fmt.Sprintf("error loading env variables: %s", err.Error()))
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
		logger.Fatal(fmt.Sprintf("failed to initialize db: %s", err.Error()))
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(cashbox.Server)
	go func() {
		if err := srv.Run(viper.GetString("port"), adaptor.FiberApp(handlers.InitRoutes())); err != nil {
			logger.Fatal(fmt.Sprintf("error occured while running http server: %s", err.Error()))
		}
	}()
	logger.Info("TodoApp Started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	<-quit

	logger.Info("TodoApp Shutting down")

	if err := srv.Shutdown(context.Background()); err != nil{
		logger.Error(fmt.Sprintf("error occured on server shutting down: %s", err.Error()))
	}
	if err := db.Close(); err != nil{
		logger.Error(fmt.Sprintf("error occured on db connection close: %s", err.Error()))
	}
}

func InitConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
