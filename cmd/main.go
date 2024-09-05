package main

import (
	testtask "github.com/KiseSaiyajin/testTask"
	"github.com/KiseSaiyajin/testTask/pkg/handler"
	"github.com/KiseSaiyajin/testTask/pkg/repository"
	"github.com/KiseSaiyajin/testTask/pkg/service"
	"log"
)

func main() {
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
	serv := new(testtask.Server)
	if err := serv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("error ocurred while running http server: %s", err.Error())
	}
}
