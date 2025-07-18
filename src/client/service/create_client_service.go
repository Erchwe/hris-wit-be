package service

import (
	"context"
	"database/sql"
	"github.com/pkg/errors"
	"github.com/wit-id/blueprint-backend-go/common/httpservice"
	"github.com/wit-id/blueprint-backend-go/src/repository/payload"
	sqlc "github.com/wit-id/blueprint-backend-go/src/repository/pgbo_sqlc"
	"github.com/wit-id/blueprint-backend-go/toolkit/config"
	"github.com/wit-id/blueprint-backend-go/toolkit/log"
)

func (s *ClientService) InsertClient(
	ctx context.Context,
	request payload.InsertClientPayload,
	user sqlc.GetUserBackofficeRow,
	cfg config.KVStore,
) (*payload.InsertClientPayload, error) { // Return *payload.Client instead of error only
	tx, err := s.mainDB.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		log.FromCtx(ctx).Error(err, "failed to begin tx")
		return nil, errors.WithStack(httpservice.ErrUnknownSource)
	}

	q := sqlc.New(s.mainDB).WithTx(tx)

	defer func() {
		if err != nil {
			if rollBackErr := tx.Rollback(); rollBackErr != nil {
				log.FromCtx(ctx).Error(err, "error rollback", rollBackErr)
			}
		}
	}()

	// Execute the insert query, and get the generated client_id
	clientID, err := q.CreateClient(ctx, request.ToEntity(cfg, user))
	if err != nil {
		log.FromCtx(ctx).Error(err, "failed to insert client")
		return nil, errors.WithStack(httpservice.ErrUnknownSource)
	}

	log.FromCtx(ctx).Info("Client successfully inserted with client_id:", clientID)

	// Commit the transaction after successful insertion
	if err = tx.Commit(); err != nil {
		log.FromCtx(ctx).Error(err, "error commit")
		return nil, errors.WithStack(httpservice.ErrUnknownSource)
	}

	// Construct the Client object to return
	insertedClient := &payload.InsertClientPayload{
		ClientID:        clientID, // The generated client_id
		ClientName:      request.ClientName,
		ShipmentAddress: request.ShipmentAddress,
		BillingAddress:  request.BillingAddress,
		CreatedBy:       user.CreatedBy,
	}

	log.FromCtx(ctx).Info("Client successfully inserted with client_id:", clientID)
	return insertedClient, nil
}
