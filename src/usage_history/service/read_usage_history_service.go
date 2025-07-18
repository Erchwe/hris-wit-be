package service

import (
	"context"

	"github.com/pkg/errors"
	"github.com/wit-id/blueprint-backend-go/common/httpservice"
	sqlc "github.com/wit-id/blueprint-backend-go/src/repository/pgbo_sqlc"
	"github.com/wit-id/blueprint-backend-go/toolkit/log"
)

func (s *UsageHistoryService) GetUsageHistoryByID(ctx context.Context, usageHistoryID string) (sqlc.UsageHistory, error) {
	q := sqlc.New(s.mainDB)

	result, err := q.GetUsageHistoryByID(ctx, usageHistoryID)
	if err != nil {
		log.FromCtx(ctx).Error(err, "failed to get usage history by ID")
		return sqlc.UsageHistory{}, errors.WithStack(httpservice.ErrNoResultData)
	}

	return result, nil
}

func (s *UsageHistoryService) GetListUsageHistory(ctx context.Context) ([]sqlc.UsageHistory, error) {
	q := sqlc.New(s.mainDB)

	result, err := q.ListUsageHistory(ctx)
	if err != nil {
		log.FromCtx(ctx).Error(err, "failed to list usage history")
		return nil, errors.WithStack(httpservice.ErrUnknownSource)
	}

	return result, nil
}
