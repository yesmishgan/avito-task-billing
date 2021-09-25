package main

import (
	"cashbox"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"go.uber.org/zap"

	//"github.com/spf13/viper"
	//"github.com/yesmishgan/autumn-2021-intern-assignment"
)

func main() {
	logger := zap.NewExample() //Change to zap.NewProduction()

	if err := InitConfig(); err != nil {
		logger.Fatal(fmt.Sprintf("error initializing configs: %s", err.Error()))
	}

	srv := new(cashbox.Server)
	if err := srv.Run(viper.GetString("port")); err != nil{
		logrus.Fatalf("")
	}
	fmt.Println("test")
}

func InitConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
