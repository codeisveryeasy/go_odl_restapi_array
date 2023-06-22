package models

type Todo struct {
	ID        int64  `json:"id"`
	Task      string `json:"task"`
	Status    string `json:"status"`
	Completed string `json:"completed"`
}
