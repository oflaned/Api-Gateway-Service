package programDB

import (
	"Mehmat/model/program"
	"Mehmat/pkg/client/postgressql"
	"context"
	"fmt"
	"github.com/jackc/pgconn"
	"time"
)

type Repository struct {
	client postgressql.Client
}

func (r Repository) Create(ctx context.Context, program program.Program) (string, error) {
	q := `
		INSERT INTO program (code, lang, name, date) 
		VALUES ($1, $2, $3, $4)
		RETURNING id
		`
	currentTime := time.Now() // Получаем текущее время

	err := r.client.QueryRow(ctx, q, program.Code, program.Lang, program.Name, currentTime).Scan(&program.Id)
	if err != nil {
		if pgError, ok := err.(*pgconn.PgError); ok {
			fmt.Sprintf("SQL Error: %s", pgError.Message)
		}
		return "", err
	}
	return program.Id, nil
}

func (r Repository) FindAll(ctx context.Context) ([]program.Program, error) {
	q := `SELECT id, code, lang, name, date from program`
	rows, err := r.client.Query(ctx, q)

	if err != nil {
		return nil, err
	}

	programs := make([]program.Program, 0)

	for rows.Next() {
		var p program.Program
		err = rows.Scan(&p.Id, &p.Code, &p.Lang, &p.Name, &p.Date)
		if err != nil {
			return nil, err
		}
		programs = append(programs, p)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return programs, nil

}

func (r Repository) FindOne(ctx context.Context, id string) (program.Program, error) {
	q := `SELECT id, code, lang, date from public.program WHERE id=$1`

	var p program.Program
	err := r.client.QueryRow(ctx, q, id).Scan(&p.Id, &p.Code, &p.Lang, &p.Date)
	if err != nil {
		return p, err
	}
	return p, nil
}

func (r Repository) Update(ctx context.Context, p program.Program) error {
	query := `
		UPDATE programs
		SET name = $1, code = $2, lang = $3, date = &4
		WHERE id = $5
	`

	_, err := r.client.Exec(ctx, query, p.Name, p.Code, p.Lang, p.Date, p.Id)
	if err != nil {
		return fmt.Errorf("failed to update program: %w", err)
	}

	return nil
}

func (r Repository) Delete(ctx context.Context, id string) error {
	query := `
		DELETE FROM programs
		WHERE id = $1
	`

	_, err := r.client.Exec(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed to delete program: %w", err)
	}

	return nil
}

func NewRepository(client postgressql.Client) program.Repository {
	return &Repository{
		client: client,
	}
}
