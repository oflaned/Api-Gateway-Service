package repository

type Authorization interface {
}

type Task interface {
}

type Repository struct {
	Authorization
	Task
}

func NewRepository() *Repository {
	return &Repository{}
}
