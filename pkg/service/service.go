package service

import "github.com/KiseSaiyajin/testTask/pkg/repository"

type Authorization interface {
}
type Service struct {
	Authorization
}

func NewService(repos *repository.Repository) *Service {
	return &Service{}
}
