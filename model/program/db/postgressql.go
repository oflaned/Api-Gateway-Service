package programDB

import (
	"Mehmat/model/program"
	"Mehmat/model/program/db/utils"
	"Mehmat/pkg/client/postgressql"
	"context"
	"fmt"
	"github.com/jackc/pgconn"
)

type repository struct {
	client postgressql.Client
}

func (r repository) Create(ctx context.Context, program program.Program) (string, error) {
	q := `
		INSERT INTO program (code, lang, task_id) 
		VALUES ($1, $2, $3)
		RETURNING id
		`
	utils.CheckTaskId(&program)
	err := r.client.QueryRow(ctx, q, program.Code, program.Lang, program.TaskId).Scan(&program.Id)
	if err != nil {
		if pgError, ok := err.(*pgconn.PgError); ok {
			fmt.Sprintf("SQL Error: %s", pgError.Message)
		}
		return "", err
	}
	return program.Id, nil
}

func (r repository) FindAll(ctx context.Context) ([]program.Program, error) {
	q := `SELECT id, code, lang, task_id from program`
	fmt.Println(1)
	rows, err := r.client.Query(ctx, q)

	if err != nil {
		return nil, err
	}

	programs := make([]program.Program, 0)

	for rows.Next() {
		var p program.Program
		err = rows.Scan(&p.Id, &p.Code, &p.Lang, &p.TaskId)
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

func (r repository) FindOne(ctx context.Context, id string) (program.Program, error) {
	q := `SELECT id, code, lang, task_id from public.program WHERE id=$1`

	var p program.Program
	err := r.client.QueryRow(ctx, q, id).Scan(&p.Id, &p.Code, &p.Lang, &p.TaskId)
	if err != nil {
		return p, err
	}
	return p, nil
}

func (r repository) Update(ctx context.Context, p program.Program) error {
	//TODO implement me
	panic("implement me")
}

func (r repository) Delete(ctx context.Context, id string) error {
	//TODO implement me
	panic("implement me")
}

func NewRepository(client postgressql.Client) program.Repository {
	return &repository{
		client: client,
	}
}
