package user

import (
	userM "visitor_management/internal/module/user"
	"visitor_management/platform/logger"
)

type UserHandlerWrapper struct {
	*user
}

type user struct {
	logger     logger.Logger
	UserModule userM.UserModuleWrapper
}

func InitUser(logger logger.Logger, userModule userM.UserModuleWrapper) *UserHandlerWrapper {
	return &UserHandlerWrapper{
		&user{
			logger,
			userModule,
		},
	}
}
