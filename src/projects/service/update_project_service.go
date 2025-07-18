package service

import (
	"context"
	"database/sql"
	"encoding/base64"
	"github.com/pkg/errors"
	"github.com/wit-id/blueprint-backend-go/common/httpservice"
	"github.com/wit-id/blueprint-backend-go/src/repository/payload"
	sqlc "github.com/wit-id/blueprint-backend-go/src/repository/pgbo_sqlc"
	"github.com/wit-id/blueprint-backend-go/toolkit/log"
)

// Function to decode the Base64 project status
func DecodeBase64(ctx context.Context, encodedStatus string) (string, error) {
	// Log the incoming status to check if it's already a valid string or Base64 encoded
	log.FromCtx(ctx).Info("Received project status: ", encodedStatus)

	// Decode the Base64 encoded string
	decodedBytes, err := base64.StdEncoding.DecodeString(encodedStatus)
	if err != nil {
		return "", err // Return error if decoding fails
	}

	decodedStatus := string(decodedBytes)

	// Validate if decoded status is one of the allowed enum values
	if decodedStatus != "open" && decodedStatus != "on progress" && decodedStatus != "done" {
		return "", errors.New("invalid project status")
	}

	return decodedStatus, nil
}

func (s *ProjectService) UpdateProject(
	ctx context.Context,
	request payload.UpdateProjectPayload,
	user sqlc.GetUserBackofficeRow,
) error {
	tx, err := s.mainDB.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		log.FromCtx(ctx).Error(err, "failed to begin tx")
		return errors.WithStack(httpservice.ErrUnknownSource)
	}

	q := sqlc.New(s.mainDB).WithTx(tx)

	defer func() {
		if err != nil {
			if rollBackErr := tx.Rollback(); rollBackErr != nil {
				log.FromCtx(ctx).Error(err, "error rollback", rollBackErr)
			}
		}
	}()

	// Decode the Base64 project status to its plain text value
	decodedStatus, err := DecodeBase64(ctx, request.ProjectStatus) // Pass ctx here
	if err != nil {
		log.FromCtx(ctx).Error(err, "invalid project status")
		return errors.WithStack(httpservice.ErrBadRequest)
	}

	// Set the decoded project status in the request
	request.ProjectStatus = decodedStatus

	// Call the UpdateProject query with the decoded status
	_, err = q.UpdateProject(ctx, request.ToEntity(user))
	if err != nil {
		log.FromCtx(ctx).Error(err, "failed to update project")
		return errors.WithStack(httpservice.ErrUnknownSource)
	}

	// Commit the transaction
	if err = tx.Commit(); err != nil {
		log.FromCtx(ctx).Error(err, "error commit")
		return errors.WithStack(httpservice.ErrUnknownSource)
	}

	return nil
}
