package program

import "github.com/jackc/pgx/pgtype"

type Program struct {
	Name string      `json:"name"`
	Code string      `json:"code"`
	Lang string      `json:"lang"`
	Id   string      `json:"id"`
	Date pgtype.Date `json:"date"`
}
