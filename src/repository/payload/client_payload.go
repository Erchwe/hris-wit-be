package payload

import (
	"database/sql"
	"github.com/asaskevich/govalidator"
	"github.com/pkg/errors"
	"github.com/wit-id/blueprint-backend-go/common/httpservice"
	sqlc "github.com/wit-id/blueprint-backend-go/src/repository/pgbo_sqlc"
	"github.com/wit-id/blueprint-backend-go/toolkit/config"
)

type InsertClientPayload struct {
	ClientID        string `json:"client_id"`
	ClientName      string `json:"client_name"`
	ShipmentAddress string `json:"shipment_address"`
	BillingAddress  string `json:"billing_address"`
	CreatedBy       string `json:"created_by" valid:"required"`
}

func (payload *InsertClientPayload) Validate() (err error) {
	if _, err = govalidator.ValidateStruct(payload); err != nil {
		err = errors.Wrapf(httpservice.ErrBadRequest, "bad request: %s", err.Error())
		return
	}
	return
}

func (payload *InsertClientPayload) ToEntity(cfg config.KVStore, userData sqlc.GetUserBackofficeRow) (data sqlc.CreateClientParams) {
	data = sqlc.CreateClientParams{
		ClientName:      sql.NullString{}, // Nullable field
		ShipmentAddress: sql.NullString{}, // Nullable field
		BillingAddress:  sql.NullString{}, // Nullable field
		CreatedBy:       userData.CreatedBy,
	}

	if payload.ClientName != "" {
		data.ClientName = sql.NullString{
			String: payload.ClientName,
			Valid:  true,
		}
	}
	if payload.ShipmentAddress != "" {
		data.ShipmentAddress = sql.NullString{
			String: payload.ShipmentAddress,
			Valid:  true,
		}
	}
	if payload.BillingAddress != "" {
		data.BillingAddress = sql.NullString{
			String: payload.BillingAddress,
			Valid:  true,
		}
	}

	return
}

type UpdateClientPayload struct {
	ClientID        string `json:"client_id" valid:"required"`
	ClientName      string `json:"client_name"`
	ShipmentAddress string `json:"shipment_address"`
	BillingAddress  string `json:"billing_address"`
}

func (payload *UpdateClientPayload) Validate() (err error) {
	if _, err = govalidator.ValidateStruct(payload); err != nil {
		err = errors.Wrapf(httpservice.ErrBadRequest, "bad request: %s", err.Error())
		return
	}
	return
}

func (payload *UpdateClientPayload) ToEntity(userData sqlc.GetUserBackofficeRow) (data sqlc.UpdateClientParams) {
	data = sqlc.UpdateClientParams{
		ClientID:        payload.ClientID,
		ClientName:      sql.NullString{}, // Nullable field
		ShipmentAddress: sql.NullString{}, // Nullable field
		BillingAddress:  sql.NullString{}, // Nullable field
	}

	if payload.ClientName != "" {
		data.ClientName = sql.NullString{
			String: payload.ClientName,
			Valid:  true,
		}
	}
	if payload.ShipmentAddress != "" {
		data.ShipmentAddress = sql.NullString{
			String: payload.ShipmentAddress,
			Valid:  true,
		}
	}
	if payload.BillingAddress != "" {
		data.BillingAddress = sql.NullString{
			String: payload.BillingAddress,
			Valid:  true,
		}
	}

	return
}
