package taskmodule

import (
	"context"
	"log"

	m "hugdev/ambiez-go/model"
)

func (p *Module) AddTask(ctx context.Context, data m.TaskRequest) (result m.TaskResponse, err error) {
	data, err = SanitizeInsert(data)
	if err != nil {
		log.Println("[TaskModule][AddTask] bad request, err: ", err.Error())
		return
	}

	result, err = p.Storage.AddTask(ctx, data)
	if err != nil {
		log.Println("[TaskModule][AddTask] problem in getting from storage, err: ", err.Error())
		return
	}
	result.Title = data.Title
	result.Completed = data.Completed
	result.Hour = data.Hour
	result.Minute = data.Minute
	return
}

func (p *Module) GetTask(ctx context.Context, id int64) (result m.TaskResponse, err error) {
	result, err = p.Storage.GetTask(ctx, id)
	if err != nil {
		log.Println("[TaskModule][GetTask] problem getting storage data, err: ", err.Error())
		return
	}

	return
}

func (p *Module) GetTaskAll(ctx context.Context) (result []m.TaskResponse, err error) {
	result, err = p.Storage.GetTaskAll(ctx, getTaskAllQuery)
	if err != nil {
		log.Println("[TaskModule][GetTaskAll] problem getting storage data, err: ", err.Error())
		return
	}

	return
}

func (p *Module) GetTodoTaskAll(ctx context.Context) (result []m.TaskResponse, err error) {
	result, err = p.Storage.GetTaskAll(ctx, getTodoTaskAllQuery)
	if err != nil {
		log.Println("[TaskModule][GetTodoTaskAll] problem getting storage data, err: ", err.Error())
		return
	}

	return
}

func (p *Module) GetCompletedTaskAll(ctx context.Context) (result []m.TaskResponse, err error) {
	result, err = p.Storage.GetTaskAll(ctx, getCompletedTaskAllQuery)
	if err != nil {
		log.Println("[TaskModule][GetCompletedTaskAll] problem getting storage data, err: ", err.Error())
		return
	}

	return
}

func (p *Module) UpdateTask(ctx context.Context, id int64, data m.TaskRequest) (result m.TaskResponse, err error) {
	result, err = p.Storage.UpdateTask(ctx, id, data)
	if err != nil {
		log.Println("[TaskModule][UpdateTask] problem getting storage data, err: ", err.Error())
		return
	}

	return
}

func (p *Module) ToggleTask(ctx context.Context, id int64) (err error) {
	err = p.Storage.ToggleTask(ctx, id)
	if err != nil {
		log.Println("[TaskModule][UpdateTask] problem getting storage data, err: ", err.Error())
		return
	}

	return
}

func (p *Module) RemoveTask(ctx context.Context, id int64) (result m.TaskResponse, err error) {
	result, err = p.Storage.RemoveTask(ctx, id)
	if err != nil {
		log.Println("[TaskModule][RemoveTask] problem getting storage data, err: ", err.Error())
		return
	}

	return
}
