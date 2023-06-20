package mossStatus

import "github.com/jackc/pgx/pgtype"

type MossStatus struct {
	Id       string      `json:"id"`
	Status   string      `json:"status"`
	Date     pgtype.Date `json:"date"`
	Programs string      `json:"Programs"`
	Link     string      `json:"link"`
}
