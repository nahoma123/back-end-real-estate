package persistence

import (
	"context"
	"time"
	"visitor_management/internal/constant"
	"visitor_management/internal/constant/errors"
	"visitor_management/internal/constant/model"
	"visitor_management/internal/storage"
	"visitor_management/platform/logger"

	"github.com/gofrs/uuid"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

type EstateStorage struct {
	db  *gorm.DB
	gnr storage.GenericStorage
}

func InitRlEstateDB(db *gorm.DB, gnr storage.GenericStorage) *EstateStorage {
	return &EstateStorage{
		db:  db,
		gnr: gnr,
	}
}

func (re *EstateStorage) AddEvaluation(ctx context.Context, vl *model.RealEstate) (*model.RealEstate, error) {
	vl.CreatedAt = time.Now()
	vl.UpdatedAt = time.Now()

	id, _ := uuid.NewV4()

	vl.RealEstateId = id.String()
	vl.Status = constant.Active

	err := re.gnr.CreateOne(ctx, constant.DbValuations, vl)
	if err != nil {
		logger.Log().Error(ctx, err.Error())
		if mongo.IsDuplicateKeyError(err) {
			return nil, errors.ErrDataExists.Wrap(err, errors.EstateIsAlreadyRegistered)
		}
		return nil, errors.ErrInternalServerError.New("unknown error occurred")
	}

	return vl, err
}
