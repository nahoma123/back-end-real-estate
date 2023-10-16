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
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type user struct {
	db *gorm.DB
}

func InitUserDB(db *gorm.DB) storage.UserStorage {
	return &user{
		db: db,
	}
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func (p *user) Create(ctx context.Context, user *model.User) (*model.User, error) {
	id, _ := uuid.NewV4()
	user.UserID = id.String()

	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	user.Status = constant.Active
	hash, err := HashPassword(user.Password)
	if err != nil {
		logger.Log().Error(ctx, err.Error())
		return nil, errors.ErrInvalidInput.New(errors.UnknownDbError)
	}

	user.Password = hash
	err = p.db.Create(user).Error
	if err != nil {
		logger.Log().Error(ctx, err.Error())
		if err == gorm.ErrInvalidData {
			return nil, errors.ErrDataExists.Wrap(err, errors.UserIsAlreadyRegistered)
		}
		return nil, errors.ErrInternalServerError.New("unknown error occurred")
	}

	return user, nil
}

func (p *user) Update(ctx context.Context, user *model.User) (*model.User, error) {
	err := p.db.Model(user).Updates(model.User{ /* fields to update */ }).Error
	if err != nil {
		logger.Log().Error(ctx, err.Error())
		return nil, errors.ErrInternalServerError.New("unknown error occurred")
	}

	return user, nil
}

func (p *user) Get(ctx context.Context, id string) (*model.User, error) {
	var user model.User
	err := p.db.First(&user, id).Error
	if err != nil {
		logger.Log().Error(ctx, err.Error())
		if err == gorm.ErrRecordNotFound {
			return nil, errors.ErrNoRecordFound.New("user not found")
		}
		return nil, errors.ErrInternalServerError.New("unknown error occurred")
	}

	return &user, nil
}

func (p *user) GetUserByEmail(ctx context.Context, email string) (*model.User, error) {
	var user model.User
	err := p.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		logger.Log().Error(ctx, err.Error())
		if err == gorm.ErrRecordNotFound {
			return nil, errors.ErrNoRecordFound.New("user not found")
		}
		return nil, errors.ErrInternalServerError.New("unknown error occurred")
	}

	return &user, nil
}

func (p *user) GetAll(ctx context.Context, filterPagination *constant.FilterPagination) ([]model.User, error) {
	var users []model.User
	err := p.db.Find(&users).Error
	if err != nil {
		logger.Log().Error(ctx, err.Error())
		return nil, errors.ErrInternalServerError.New("unknown error occurred")
	}

	return users, nil
}
