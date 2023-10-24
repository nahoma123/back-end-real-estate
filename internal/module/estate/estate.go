package estate

import (
	"context"
	"visitor_management/internal/constant/errors"
	"visitor_management/internal/constant/model"
	"visitor_management/internal/module"
	"visitor_management/internal/storage/persistence"
	"visitor_management/platform/logger"

	"go.uber.org/zap"
)

type Estate struct {
	logger logger.Logger
	rlEst  persistence.EstateStorage
}

func InitEstate(logger logger.Logger, rlEst persistence.EstateStorage) module.EstateModule {
	return &Estate{
		logger,
		rlEst,
	}
}

func (es *Estate) CreateEstate(ctx context.Context, valuation *model.RealEstate) (*model.RealEstate, error) {
	if err := valuation.Validate(); err != nil {
		err = errors.ErrInvalidInput.Wrap(err, "invalid input")
		es.logger.Info(ctx, "invalid input", zap.Error(err))
		return nil, err
	}

	valuation, err := es.rlEst.AddEvaluation(ctx, valuation)
	if err != nil {
		es.logger.Warn(ctx, err.Error())
		return nil, err
	}
	return valuation, nil
}
