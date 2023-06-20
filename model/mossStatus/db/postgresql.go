package db

import (
	"Mehmat/model/mossStatus"
	"Mehmat/pkg/client/postgressql"
	"context"
	"fmt"
	"github.com/jackc/pgconn"
)

type Repository struct {
	client postgressql.Client
}

func (r *Repository) Create(ctx context.Context, status mossStatus.MossStatus) (string, error) {
	q := `
		INSERT INTO mossstatus (status, date, programs, links) 
		VALUES ($1, $2, $3)
		RETURNING id
	`

	err := r.client.QueryRow(ctx, q, status.Status, status.Date, status.Programs, status.Link).Scan(&status.Id)
	if err != nil {
		if pgError, ok := err.(*pgconn.PgError); ok {
			fmt.Sprintf("SQL Error: %s", pgError.Message)
		}
		return "", err
	}

	return status.Id, nil
}

func (r *Repository) FindAll(ctx context.Context) ([]mossStatus.MossStatus, error) {
	q := `SELECT id, status, date, programs, links from mossstatus`

	rows, err := r.client.Query(ctx, q)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	statuses := make([]mossStatus.MossStatus, 0)

	for rows.Next() {
		var p mossStatus.MossStatus
		err = rows.Scan(&p.Id, &p.Status, &p.Date, &p.Programs, &p.Link)
		if err != nil {
			return nil, err
		}
		statuses = append(statuses, p)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return statuses, nil

}

func (r *Repository) FindOne(ctx context.Context, id string) (mossStatus.MossStatus, error) {
	q := `SELECT id, status, date, programs, links from mossstatus WHERE id=$1`

	var p mossStatus.MossStatus
	err := r.client.QueryRow(ctx, q, id).Scan(&p.Id, &p.Status, &p.Date, &p.Programs, &p.Link)
	if err != nil {
		return p, err
	}

	return p, nil
}

func (r *Repository) Update(ctx context.Context, p mossStatus.MossStatus) error {
	query := `
		UPDATE mossstatus
		SET status = $1, date = $2, programs = $3, links = $4
		WHERE id = $5
	`

	_, err := r.client.Exec(ctx, query, p.Status, p.Date, p.Programs, p.Link, p.Id)
	if err != nil {
		return fmt.Errorf("failed to update program: %w", err)
	}

	return nil
}

func (r *Repository) Delete(ctx context.Context, id string) error {
	query := `
		DELETE FROM mossstatus
		WHERE id = $1
	`

	_, err := r.client.Exec(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed to delete program: %w", err)
	}

	return nil
}

func NewRepository(client postgressql.Client) mossStatus.Repository {
	return &Repository{
		client: client,
	}
}
