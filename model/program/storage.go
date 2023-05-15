package program

import "context"

type Repository interface {
	Create(ctx context.Context, program Program) (string, error)
	FindAll(ctx context.Context) (p []Program, err error)
	FindOne(ctx context.Context, id string) (Program, error)
	Update(ctx context.Context, p Program) error
	Delete(ctx context.Context, id string) error
}
