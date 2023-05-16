package program

type Program struct {
	Name   string `json:"name"`
	Code   string `json:"code"`
	Lang   string `json:"lang"`
	Id     string `json:"id"`
	TaskId string `json:"task_id"`
}
