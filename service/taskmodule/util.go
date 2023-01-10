package taskmodule

import (
	"errors"

	m "hugdev/ambiez-go/model"
)

func int(i uint) *uint     { return &i }
func boolean(b bool) *bool { return &b }

func SanitizeInsert(param m.TaskRequest) (m.TaskRequest, error) {

	if param.Title == nil {
		return param, errors.New("Title url cannot be empty")
	}
	if param.Completed == nil {
		param.Completed = boolean(false)
	}
	if param.Hour == nil {
		param.Hour = int(0)
	}
	if param.Minute == nil {
		param.Minute = int(0)
	}

	return param, nil
}

func SanitizeUpdate(param m.TaskRequest) (m.TaskRequest, error) {

	if param.Title == nil {
		return param, errors.New("Title url cannot be empty")
	}

	return param, nil
}
