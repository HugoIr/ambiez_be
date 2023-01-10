package model

type TaskResponse struct {
	ID        int64   `json:"id" db:"id"`
	Title     *string `json:"title" db:"title"`
	Completed *bool   `json:"completed" db:"completed"`
	Hour      *uint   `json:"hour" db:"hour"`
	Minute    *uint   `json:"minute" db:"minute"`
}

type TaskRequest struct {
	Title     *string `json:"title" db:"title"`
	Completed *bool   `json:"completed" db:"completed"`
	Hour      *uint   `json:"hour" db:"hour"`
	Minute    *uint   `json:"minute" db:"minute"`
}
