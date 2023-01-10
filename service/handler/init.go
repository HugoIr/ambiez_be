package handler

import "hugdev/ambiez-go/taskmodule"

type InsertTaskResponse struct {
	ID int64 `json:"id"`
}

type Handler struct {
	ambiez *taskmodule.Module
}

func NewAmbiezHandler(p *taskmodule.Module) *Handler {
	return &Handler{
		ambiez: p,
	}
}
