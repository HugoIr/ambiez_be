package taskmodule

import (
	"errors"

	m "hugdev/ambiez-go/model"
)

func i(i uint) *uint { return &i }

func SanitizeInsert(param m.TaskRequest) (m.TaskRequest, error) {

	if param.Title == nil {
		return param, errors.New("Title url cannot be empty")
	}
	if param.Hour == nil {
		param.Hour = i(0)
	}
	if param.Minute == nil {
		param.Minute = i(0)
	}

	return param, nil
}

func SanitizeUpdate(param m.TaskRequest) (m.TaskRequest, error) {

	if param.Title == nil {
		return param, errors.New("Title url cannot be empty")
	}

	return param, nil
}
