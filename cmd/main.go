package main

import (
	testtask "github.com/KiseSaiyajin/testTask"
	"github.com/KiseSaiyajin/testTask/pkg/handler"
	"github.com/KiseSaiyajin/testTask/pkg/repository"
	"github.com/KiseSaiyajin/testTask/pkg/service"
	"github.com/spf13/viper"
	"log"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error in initializing configs: %s", err.Error())
	}
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
	serv := new(testtask.Server)
	if err := serv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatalf("error ocurred while running http server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
