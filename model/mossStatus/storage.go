package mossStatus

import "context"

type Repository interface {
	Create(ctx context.Context, status MossStatus) (string, error)
	FindAll(ctx context.Context) (s []MossStatus, err error)
	FindOne(ctx context.Context, id string) (MossStatus, error)
	Update(ctx context.Context, s MossStatus) error
	Delete(ctx context.Context, id string) error
}
