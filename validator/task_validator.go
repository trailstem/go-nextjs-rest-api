package validater

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/trailstem/go-nextjs-rest-api/model"
)

// interface

type ITaskValidator interface {
	ITaskValidate(task model.Task) error
}

// 構造他

type taskValidator struct{}

// constructor

func NewTaskValidator() ITaskValidator {
	return &taskValidator{}
}

func (tv *taskValidator) ITaskValidate(task model.Task) error {
	return validation.ValidateStruct(&task,
		validation.Field(
			&task.Title,
			validation.Required.Error("title is required"),
			validation.RuneLength(1, 10).Error("limited max 10 char"),
		),
	)
}
