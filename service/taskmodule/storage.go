package taskmodule

import (
	"context"
	"database/sql"
	"log"

	m "hugdev/ambiez-go/model"
)

type storage struct {
	AmbiezDB *sql.DB
}

func newStorage(db *sql.DB) *storage {
	return &storage{
		AmbiezDB: db,
	}
}

func (s *storage) AddTask(ctx context.Context, data m.TaskRequest) (result m.TaskResponse, err error) {
	var id int64
	if err := s.AmbiezDB.QueryRowContext(ctx, addTaskQuery,
		data.Title,
		data.Completed,
		data.Hour,
		data.Minute,
	).Scan(&id); err != nil {
		log.Println("[TaskModule][AddTask][Storage] problem querying to db, err: ", err.Error())
		return result, err
	}

	result.ID = id

	return
}

func (s *storage) GetTask(ctx context.Context, id int64) (result m.TaskResponse, err error) {
	if err := s.AmbiezDB.QueryRowContext(ctx, getTaskQuery, id).Scan(
		&result.Title,
		&result.Completed,
		&result.Hour,
		&result.Minute,
	); err != nil {
		log.Println("[TaskModule][GetTask] problem querying to db, err: ", err.Error())
		return result, err
	}
	result.ID = id

	return
}

func (s *storage) GetTaskAll(ctx context.Context) (result []m.TaskResponse, err error) {
	result = make([]m.TaskResponse, 0)

	rows, err := s.AmbiezDB.QueryContext(ctx, getTaskAllQuery)
	if err != nil {
		log.Println("[TaskModule][GetTaskAll] problem querying to db, err: ", err.Error())
		return
	}
	defer rows.Close()
	for rows.Next() {
		var rowData m.TaskResponse

		if err = rows.Scan(
			&rowData.ID,
			&rowData.Title,
			&rowData.Completed,
			&rowData.Hour,
			&rowData.Minute,
		); err != nil {
			log.Println("[TaskModule][GetTaskAll] problem with scanning db row, err: ", err.Error())
			return
		}
		result = append(result, rowData)
	}

	return
}

func (s *storage) UpdateTask(ctx context.Context, id int64, param m.TaskRequest) (result m.TaskResponse, err error) {
	res, err := s.AmbiezDB.ExecContext(ctx, updateTaskQuery,
		param.Title,
		param.Completed,
		param.Hour,
		param.Minute,
		id,
	)
	if err != nil {
		log.Println("[TaskModule][UpdateTask][Storage] problem querying to db, err: ", err.Error())
		return
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		log.Println("[TaskModule][UpdateTask] problem querying to db, err: ", err.Error())
		return
	}
	if rowsAffected == 0 {
		log.Println("[TaskModule][UpdateTask] no rows affected in db")
		return
	}

	result.ID = id

	return
}

func (s *storage) ToggleTask(ctx context.Context, id int64) (err error) {
	res, err := s.AmbiezDB.ExecContext(ctx, toggleTaskQuery,
		id,
	)
	if err != nil {
		log.Println("[TaskModule][UpdateTask][Storage] problem querying to db, err: ", err.Error())
		return
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		log.Println("[TaskModule][UpdateTask] problem querying to db, err: ", err.Error())
		return
	}
	if rowsAffected == 0 {
		log.Println("[TaskModule][UpdateTask] no rows affected in db")
		return
	}

	return
}

func (s *storage) RemoveTask(ctx context.Context, id int64) (result m.TaskResponse, err error) {

	res, err := s.AmbiezDB.ExecContext(ctx, removeTaskQuery, id)

	rowsAffected, err := res.RowsAffected()

	if err != nil {
		log.Println("[TaskModule][RemoveTask] problem querying to db, err: ", err.Error())
		return result, err
	}
	if rowsAffected == 0 {
		log.Println("[TaskModule][RemoveTask] no effect ")
		return
	}

	return
}
