package user

import (
	"net/http"
	"visitor_management/internal/constant"
	"visitor_management/internal/constant/errors"
	"visitor_management/internal/constant/model"
	"visitor_management/internal/handler/rest"
	"visitor_management/internal/module"
	"visitor_management/platform/logger"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type user struct {
	logger     logger.Logger
	UserModule module.UserModule
}

func InitUser(logger logger.Logger, userModule module.UserModule) rest.User {
	return &user{
		logger,
		userModule,
	}
}

func (o *user) Register(ctx *gin.Context) {
	user := &model.User{}
	err := ctx.ShouldBind(&user)
	if err != nil {
		o.logger.Info(ctx, zap.Error(err).String)
		_ = ctx.Error(errors.ErrInvalidInput.Wrap(err, "invalid input"))
		return
	}

	user, err = o.UserModule.RegisterUser(ctx, user)
	if err != nil {
		o.logger.Info(ctx, zap.Error(err).String)
		_ = ctx.Error(err)
		return
	}

	constant.SuccessResponse(ctx, http.StatusCreated, user, nil)
}

func (o *user) UpdateUser(ctx *gin.Context) {
	user := &model.User{}
	err := ctx.ShouldBind(&user)
	id := ctx.Param("id")
	if err != nil || id == "" {
		o.logger.Info(ctx, zap.Error(err).String)
		_ = ctx.Error(errors.ErrInvalidInput.Wrap(err, "invalid input"))
		return
	}
	user.UserID = id
	user, err = o.UserModule.UpdateUser(ctx, user)
	if err != nil {
		o.logger.Info(ctx, zap.Error(err).String)
		_ = ctx.Error(err)
		return
	}

	constant.SuccessResponse(ctx, http.StatusCreated, user, nil)
}

func (o *user) GetUser(ctx *gin.Context) {
	id := ctx.Param("id")
	user, err := o.UserModule.GetUser(ctx, id)
	if err != nil {
		o.logger.Info(ctx, zap.Error(err).String)
		_ = ctx.Error(err)
		return
	}

	constant.SuccessResponse(ctx, http.StatusOK, user, nil)
}

func (o *user) Login(ctx *gin.Context) {
	user := &model.User{}
	err := ctx.Bind(user)
	if err != nil {
		o.logger.Info(ctx, zap.Error(err).String)
		_ = ctx.Error(err)
		return
	}

	auth, err := o.UserModule.Login(ctx, user.Email, user.Password)
	if err != nil {
		o.logger.Info(ctx, zap.Error(err).String)
		_ = ctx.Error(err)
		return
	}

	constant.SuccessResponse(ctx, http.StatusOK, auth, nil)
}

func (o *user) GetUsers(ctx *gin.Context) {
	ftr := constant.ParseFilterPagination(ctx)

	states, err := o.UserModule.GetAll(ctx, ftr)
	if err != nil {
		o.logger.Info(ctx, zap.Error(err).String)
		_ = ctx.Error(err)
		return
	}

	constant.SuccessResponse(ctx, http.StatusOK, states, ftr)
}
