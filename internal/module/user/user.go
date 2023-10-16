package user

import (
	"context"
	"visitor_management/internal/constant/errors"
	"visitor_management/internal/constant/model"

	"go.uber.org/zap"
)

func (o *user) RegisterUser(ctx context.Context, user *model.User) (*model.User, error) {
	//
	if err := user.Validate(); err != nil {
		err = errors.ErrInvalidInput.Wrap(err, "invalid input")
		o.logger.Info(ctx, "invalid input", zap.Error(err))
		return nil, err
	}

	user, err := o.userStorage.Create(ctx, user)
	if err != nil {
		o.logger.Warn(ctx, err.Error())
		return nil, err
	}
	return user, nil
}

func (o *user) UpdateUser(ctx context.Context, user *model.User) (*model.User, error) {
	// if err := user.ValidateUpdateUser(); err != nil {
	// 	err = errors.ErrInvalidInput.Wrap(err, "invalid input")
	// 	o.logger.Info(ctx, "invalid input", zap.Error(err))
	// 	return nil, err
	// }

	// user, err := o.userStorage.Update(ctx, user)
	// if err != nil {
	// 	o.logger.Warn(ctx, err.Error())
	// 	return nil, err
	// }
	return nil, nil
}

func (o *user) GetUser(ctx context.Context, id string) (*model.User, error) {
	user, err := o.userStorage.Get(ctx, id)
	if err != nil {
		o.logger.Warn(ctx, err.Error())
		return nil, err
	}
	return user, nil
}

func (o *user) GetUserByEmail(ctx context.Context, email string) (*model.User, error) {
	user, err := o.userStorage.GetUserByEmail(ctx, email)
	if err != nil {
		o.logger.Warn(ctx, err.Error())
		return nil, err
	}
	return user, nil
}
