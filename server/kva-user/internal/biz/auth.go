package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"go.uber.org/zap"
	kvaAuth "kva-user/api/kva-user"
	"kva-user/internal/model"
)

type KvaAuthRepo interface {
	Login(ctx context.Context, username, password string)(error, *model.SysUser)
	tokenNext(ctx context.Context, user model.SysUser)(error, *kvaAuth.UserLoginRes, string)
}

type KvaAuthUsecase struct {
	repo KvaAuthRepo
	log  *log.Helper
}

func NewKvaAuthUsecase(repo KvaAuthRepo, logger log.Logger) *KvaAuthUsecase {
	return &KvaAuthUsecase{
		repo: repo,
		log:  log.NewHelper(logger),
	}
}

func (au *KvaAuthUsecase) Login(ctx context.Context, username, password string) (error, *model.SysUser) {
	au.log.WithContext(ctx).Infof("UserLogin: %v", username)
	u := &model.SysUser{Username: username, Password: password}
	if err, user := au.repo.Login(u); err != nil {
		au.log.Error("登陆失败! 用户名不存在或者密码错误!", zap.Error(err))
		response.FailWithMessage("用户名不存在或者密码错误", c)
	} else {
		au.repo.tokenNext(ctx, *user)
	}
	return au.repo.Login(ctx, username, password)
}
