package program

import "Mehmat/model/Task"

type Program struct {
	Code   string    `json:"code"`
	Lang   string    `json:"lang"`
	Id     int       `json:"id"`
	TaskId Task.Task `json:"task_id"`
}
