package payload

import (
	"database/sql"
	"github.com/asaskevich/govalidator"
	"github.com/pkg/errors"
	"github.com/wit-id/blueprint-backend-go/common/httpservice"
	sqlc "github.com/wit-id/blueprint-backend-go/src/repository/pgbo_sqlc"
	"time"
)

type UpdateUsageHistoryPayload struct {
	UsageHistoryID string `json:"usage_history_id" valid:"required"`
	InventarisID   string `json:"inventaris_id" valid:"required"`
	OldRoomID      string `json:"old_room_id" valid:"required"`
	NewRoomID      string `json:"new_room_id" valid:"required"`
	OldUserID      string `json:"old_user_id" valid:"required"`
	NewUserID      string `json:"new_user_id" valid:"required"`
	MovedAt        string `json:"moved_at"`
	UpdatedBy      string `json:"updated_by" valid:"required"`
}

func (p *UpdateUsageHistoryPayload) Validate() error {
	if _, err := govalidator.ValidateStruct(p); err != nil {
		return errors.Wrap(httpservice.ErrBadRequest, err.Error())
	}
	return nil
}

func (p *UpdateUsageHistoryPayload) ToEntity() (sqlc.UpdateUsageHistoryParams, error) {
	var movedAt sql.NullTime
	if p.MovedAt != "" {
		t, err := time.Parse("2006-01-02 15:04:05", p.MovedAt)
		if err != nil {
			return sqlc.UpdateUsageHistoryParams{}, errors.Wrap(httpservice.ErrBadRequest, "invalid format for moved_at")
		}
		movedAt = sql.NullTime{Time: t, Valid: true}
	}

	return sqlc.UpdateUsageHistoryParams{
		UsageHistoryID: p.UsageHistoryID,
		InventarisID:   p.InventarisID,
		OldRoomID:      p.OldRoomID,
		NewRoomID:      p.NewRoomID,
		OldUserID:      p.OldUserID,
		NewUserID:      p.NewUserID,
		MovedAt:        movedAt,
		UpdatedBy:      sql.NullString{String: p.UpdatedBy, Valid: true},
	}, nil
}
