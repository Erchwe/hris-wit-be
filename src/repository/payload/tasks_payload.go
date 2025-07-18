package payload

import (
	"database/sql"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/pkg/errors"
	"github.com/wit-id/blueprint-backend-go/common/httpservice"
	sqlc "github.com/wit-id/blueprint-backend-go/src/repository/pgbo_sqlc"
	"github.com/wit-id/blueprint-backend-go/toolkit/config"
)

type InsertTaskPayload struct {
	ListID       string `json:"list_id" valid:"required"`
	TaskName     string `json:"task_name"`
	TaskType     string `json:"task_type"`
	TaskPriority string `json:"task_priority"`
	TaskSize     string `json:"task_size"`
	TaskStatus   string `json:"task_status"`
	TaskColor    string `json:"task_color"`
	StartDate    string `json:"start_date"`
	DueDate      string `json:"due_date"`
	CreatedBy    string `json:"created_by" valid:"required"`
}

func (payload *InsertTaskPayload) Validate() (err error) {
	if _, err = govalidator.ValidateStruct(payload); err != nil {
		err = errors.Wrapf(httpservice.ErrBadRequest, "bad request: %s", err.Error())
		return
	}
	return
}

func (payload *InsertTaskPayload) ToEntity(cfg config.KVStore, userData sqlc.GetUserBackofficeRow) (data sqlc.CreateTaskParams) {
	data = sqlc.CreateTaskParams{
		ListID:       payload.ListID,
		TaskName:     sql.NullString{}, // Nullable field
		TaskType:     sql.NullString{}, // Nullable ENUM field
		TaskPriority: sql.NullString{}, // Nullable ENUM field
		TaskSize:     sql.NullString{}, // Nullable ENUM field
		TaskStatus:   sql.NullString{}, // Nullable ENUM field
		TaskColor:    sql.NullString{}, // Nullable field
		StartDate:    sql.NullTime{},   // Nullable field
		DueDate:      sql.NullTime{},   // Nullable field
		CreatedBy:    userData.CreatedBy,
	}

	if payload.TaskName != "" {
		data.TaskName = sql.NullString{
			String: payload.TaskName,
			Valid:  true,
		}
	}
	if payload.TaskType != "" {
		data.TaskType = sql.NullString{
			String: payload.TaskType,
			Valid:  true,
		}
	}
	if payload.TaskPriority != "" {
		data.TaskPriority = sql.NullString{
			String: payload.TaskPriority,
			Valid:  true,
		}
	}
	if payload.TaskSize != "" {
		data.TaskSize = sql.NullString{
			String: payload.TaskSize,
			Valid:  true,
		}
	}
	if payload.TaskStatus != "" {
		data.TaskStatus = sql.NullString{
			String: payload.TaskStatus,
			Valid:  true,
		}
	}
	if payload.TaskColor != "" {
		data.TaskColor = sql.NullString{
			String: payload.TaskColor,
			Valid:  true,
		}
	}
	if payload.StartDate != "" {
		parsedStartDate, err := time.Parse("2006-01-02", payload.StartDate)
		if err == nil {
			data.StartDate = sql.NullTime{Time: parsedStartDate, Valid: true}
		}
	}
	if payload.DueDate != "" {
		parsedDueDate, err := time.Parse("2006-01-02", payload.DueDate)
		if err == nil {
			data.DueDate = sql.NullTime{Time: parsedDueDate, Valid: true}
		}
	}

	return
}

type UpdateTaskPayload struct {
	TaskID       string  `json:"task_id" valid:"required"`
	ListID       string  `json:"list_id"`
	TaskOrder    float64 `json:"task_order"`
	TaskName     string  `json:"task_name"`
	TaskType     string  `json:"task_type"`
	TaskPriority string  `json:"task_priority"`
	TaskSize     string  `json:"task_size"`
	TaskStatus   string  `json:"task_status"`
	TaskColor    string  `json:"task_color"`
	StartDate    string  `json:"start_date"`
	DueDate      string  `json:"due_date"`
	UpdatedBy    string  `json:"updated_by" valid:"required"`
}

func (payload *UpdateTaskPayload) Validate() (err error) {
	if _, err = govalidator.ValidateStruct(payload); err != nil {
		err = errors.Wrapf(httpservice.ErrBadRequest, "bad request: %s", err.Error())
		return
	}
	return
}

func (payload *UpdateTaskPayload) ToEntity(userData sqlc.GetUserBackofficeRow) (data sqlc.UpdateTaskParams) {
	data = sqlc.UpdateTaskParams{
		TaskID: payload.TaskID,
		ListID: payload.ListID,
		TaskOrder: sql.NullFloat64{
			Float64: payload.TaskOrder,
			Valid:   payload.TaskOrder != 0,
		},
		TaskName:     sql.NullString{}, // Nullable field
		TaskType:     sql.NullString{}, // Nullable field
		TaskPriority: sql.NullString{}, // Nullable field
		TaskSize:     sql.NullString{}, // Nullable field
		TaskStatus:   sql.NullString{}, // Nullable field
		TaskColor:    sql.NullString{}, // Nullable field
		StartDate:    sql.NullTime{},   // Nullable field
		DueDate:      sql.NullTime{},   // Nullable field
		UpdatedBy: sql.NullString{
			String: userData.UpdatedBy.String,
			Valid:  true,
		},
	}

	if payload.TaskName != "" {
		data.TaskName = sql.NullString{
			String: payload.TaskName,
			Valid:  true,
		}
	}

	if payload.TaskType != "" {
		data.TaskType = sql.NullString{
			String: payload.TaskType,
			Valid:  true,
		}
	}

	if payload.TaskPriority != "" {
		data.TaskPriority = sql.NullString{
			String: payload.TaskPriority,
			Valid:  true,
		}
	}

	if payload.TaskSize != "" {
		data.TaskSize = sql.NullString{
			String: payload.TaskSize,
			Valid:  true,
		}
	}

	if payload.TaskStatus != "" {
		data.TaskStatus = sql.NullString{
			String: payload.TaskStatus,
			Valid:  true,
		}
	}

	if payload.TaskColor != "" {
		data.TaskColor = sql.NullString{
			String: payload.TaskColor,
			Valid:  true,
		}
	}

	if payload.StartDate != "" {
		parsedStartDate, err := time.Parse("2006-01-02", payload.StartDate)
		if err == nil {
			data.StartDate = sql.NullTime{Time: parsedStartDate, Valid: true}
		}
	}

	if payload.DueDate != "" {
		parsedDueDate, err := time.Parse("2006-01-02", payload.DueDate)
		if err == nil {
			data.DueDate = sql.NullTime{Time: parsedDueDate, Valid: true}
		}
	}

	return
}
