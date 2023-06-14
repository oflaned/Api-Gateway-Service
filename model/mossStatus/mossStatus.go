package mossStatus

type mossStatus struct {
	Id       string `json:"id"`
	Status   string `json:"status"`
	Date     string `json:"date"`
	Programs []int  `json:"Programs"`
}
