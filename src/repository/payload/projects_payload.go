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

type InsertProjectPayload struct {
	ClientID        string `json:"client_id" valid:"required"`
	ProjectName     string `json:"project_name"`
	ProjectDesc     string `json:"project_desc"`
	ProjectStatus   string `json:"project_status"`
	ProjectPriority string `json:"project_priority"`
	ProjectColor    string `json:"project_color"`
	StartDate       string `json:"start_date"`
	DueDate         string `json:"due_date"`
	CreatedBy       string `json:"created_by" valid:"required"`
}

func (payload *InsertProjectPayload) Validate() (err error) {
	if _, err = govalidator.ValidateStruct(payload); err != nil {
		err = errors.Wrapf(httpservice.ErrBadRequest, "bad request: %s", err.Error())
		return
	}
	return
}

func (payload *InsertProjectPayload) ToEntity(cfg config.KVStore, userData sqlc.GetUserBackofficeRow) (data sqlc.CreateProjectParams) {
	data = sqlc.CreateProjectParams{
		ClientID:        payload.ClientID,
		ProjectName:     sql.NullString{}, // Nullable field
		ProjectDesc:     sql.NullString{}, // Nullable field
		ProjectStatus:   sql.NullString{}, // Nullable field
		ProjectPriority: sql.NullString{}, // Nullable field
		ProjectColor:    sql.NullString{}, // Nullable field
		StartDate:       sql.NullTime{},   // Nullable field
		DueDate:         sql.NullTime{},   // Nullable field
		CreatedBy:       userData.CreatedBy,
	}

	if payload.ProjectName != "" {
		data.ProjectName = sql.NullString{
			String: payload.ProjectName,
			Valid:  true,
		}
	}
	if payload.ProjectDesc != "" {
		data.ProjectDesc = sql.NullString{
			String: payload.ProjectDesc,
			Valid:  true,
		}
	}
	if payload.ProjectStatus != "" {
		data.ProjectStatus = sql.NullString{
			String: payload.ProjectStatus,
			Valid:  true,
		}
	}
	if payload.ProjectPriority != "" {
		data.ProjectPriority = sql.NullString{
			String: payload.ProjectPriority,
			Valid:  true,
		}
	}
	if payload.ProjectColor != "" {
		data.ProjectColor = sql.NullString{
			String: payload.ProjectColor,
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

type UpdateProjectPayload struct {
	ProjectID       string `json:"project_id" valid:"required"`
	ProjectName     string `json:"project_name"`
	ProjectDesc     string `json:"project_desc"`
	ProjectStatus   string `json:"project_status"`
	ProjectPriority string `json:"project_priority"`
	ProjectColor    string `json:"project_color"`
	StartDate       string `json:"start_date"`
	DueDate         string `json:"due_date"`
	ClientID        string `json:"client_id"`
	UpdatedBy       string `json:"updated_by" valid:"required"`
}

func (payload *UpdateProjectPayload) Validate() (err error) {
	if _, err = govalidator.ValidateStruct(payload); err != nil {
		err = errors.Wrapf(httpservice.ErrBadRequest, "bad request: %s", err.Error())
		return
	}
	return
}

func (payload *UpdateProjectPayload) ToEntity(userData sqlc.GetUserBackofficeRow) (data sqlc.UpdateProjectParams) {
	data = sqlc.UpdateProjectParams{
		ProjectID:       payload.ProjectID,
		ProjectName:     sql.NullString{}, // Nullable field
		ProjectDesc:     sql.NullString{}, // Nullable field
		ProjectStatus:   sql.NullString{}, // Nullable field
		ProjectPriority: sql.NullString{}, // Nullable field
		ProjectColor:    sql.NullString{}, // Nullable field
		StartDate:       sql.NullTime{},   // Nullable field
		DueDate:         sql.NullTime{},   // Nullable field
		ClientID:        payload.ClientID,
		UpdatedBy:       sql.NullString{},
	}

	if payload.ProjectName != "" {
		data.ProjectName = sql.NullString{
			String: payload.ProjectName,
			Valid:  true,
		}
	}
	if payload.ProjectDesc != "" {
		data.ProjectDesc = sql.NullString{
			String: payload.ProjectDesc,
			Valid:  true,
		}
	}
	if payload.ProjectStatus != "" {
		data.ProjectStatus = sql.NullString{
			String: payload.ProjectStatus,
			Valid:  true,
		}
	}
	if payload.ProjectPriority != "" {
		data.ProjectPriority = sql.NullString{
			String: payload.ProjectPriority,
			Valid:  true,
		}
	}
	if payload.ProjectColor != "" {
		data.ProjectColor = sql.NullString{
			String: payload.ProjectColor,
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
