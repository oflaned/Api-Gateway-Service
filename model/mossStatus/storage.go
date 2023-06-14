package mossStatus

import "context"

type Repository interface {
	Create(ctx context.Context, status mossStatus) (string, error)
	FindAll(ctx context.Context) (s []mossStatus, err error)
	FindOne(ctx context.Context, id string) (mossStatus, error)
	Update(ctx context.Context, s mossStatus) error
	Delete(ctx context.Context, id string) error
}
