package service

import "Mehmat/Api/repository"

type Authorization interface {
}

type Task interface {
}

type Service struct {
	Authorization
	Task
}

func NewService(repos *repository.Repository) *Service {
	return &Service{}
}
