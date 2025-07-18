package payload

import (
	"database/sql"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/pkg/errors"
	"github.com/wit-id/blueprint-backend-go/common/httpservice"
	"github.com/wit-id/blueprint-backend-go/common/utility"
	sqlc "github.com/wit-id/blueprint-backend-go/src/repository/pgbo_sqlc"
	"github.com/wit-id/blueprint-backend-go/toolkit/config"
)

type CreateUsageHistoryPayload struct {
	UsageHistoryID string `json:"usage_history_id"`
	InventarisID   string `json:"inventaris_id" valid:"required"`
	OldRoomID      string `json:"old_room_id" valid:"required"`
	NewRoomID      string `json:"new_room_id" valid:"required"`
	OldUserID      string `json:"old_user_id" valid:"required"`
	NewUserID      string `json:"new_user_id" valid:"required"`
	MovedAt        string `json:"moved_at"` // format: YYYY-MM-DD HH:MM:SS (optional)
}

func (p *CreateUsageHistoryPayload) Validate() error {
	if _, err := govalidator.ValidateStruct(p); err != nil {
		return errors.Wrap(httpservice.ErrBadRequest, err.Error())
	}
	return nil
}

func (p *CreateUsageHistoryPayload) ToEntity(cfg config.KVStore, user sqlc.GetUserBackofficeRow) (sqlc.CreateUsageHistoryParams, error) {
	id := p.UsageHistoryID
	if id == "" {
		id = utility.GenerateGoogleUUID()
	}

	var movedAt sql.NullTime
	if p.MovedAt != "" {
		t, err := time.Parse("2006-01-02 15:04:05", p.MovedAt)
		if err != nil {
			return sqlc.CreateUsageHistoryParams{}, errors.Wrap(httpservice.ErrBadRequest, "invalid format for moved_at")
		}
		movedAt = sql.NullTime{Time: t, Valid: true}
	}

	return sqlc.CreateUsageHistoryParams{
		UsageHistoryID: id,
		InventarisID:   p.InventarisID,
		OldRoomID:      p.OldRoomID,
		NewRoomID:      p.NewRoomID,
		OldUserID:      p.OldUserID,
		NewUserID:      p.NewUserID,
		MovedAt:        movedAt,
		CreatedBy:      user.CreatedBy,
	}, nil
}
