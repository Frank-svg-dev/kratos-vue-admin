package data

import (
	"context"
	"errors"
	"github.com/go-kratos/kratos/v2/log"
	"go.uber.org/zap"
	"gorm.io/gorm"
	kvaAuth "kva-user/api/kva-user"
	"kva-user/internal/biz"
	"kva-user/internal/conf"
	"kva-user/internal/model"
	"kva-user/internal/model/request"
	"kva-user/internal/pkg"
)

type kvaAuthRepo struct {
	data *conf.Data
	log  *log.Helper
	Db 	 *gorm.DB
}

// NewkvaAuthRepo
func NewkvaAuthRepo(data *conf.Data, logger log.Logger, db *gorm.DB) biz.KvaAuthRepo {
	return &kvaAuthRepo{
		data: data,
		log:  log.NewHelper(logger),
		Db: db,
	}
}

func (au *kvaAuthRepo)Login(ctx context.Context, username, password string)(error, *model.SysUser){
	var user model.SysUser
	err := au.Db.Where("username = ?", username).Preload("Authorities").Preload("Authority").First(&user).Error
	if err == nil {
		if ok := pkg.BcryptCheck(password, user.Password); !ok {
			return errors.New("密码错误"), nil
		}
	}
	return nil, &user
}

// 登录以后签发jwt
func (au *kvaAuthRepo) tokenNext(ctx context.Context, user model.SysUser)(error, *kvaAuth.UserLoginRes, string) {
	j := &pkg.JWT{SigningKey: []byte(au.data.Jwt.SigningKey)} // 唯一签名
	claims := j.CreateClaims(request.BaseClaims{
		UUID:        user.UUID,
		ID:          user.ID,
		NickName:    user.NickName,
		Username:    user.Username,
		AuthorityId: user.AuthorityId,
	})
	token, err := j.CreateToken(claims)
	if err != nil {
		au.log.Error("获取token失败!", zap.Error(err))
		return err, nil, "获取token失败"
	}
	//if !global.GVA_CONFIG.System.UseMultipoint {
	//	response.OkWithDetailed(systemRes.LoginResponse{
	//		User:      user,
	//		Token:     token,
	//		ExpiresAt: claims.StandardClaims.ExpiresAt * 1000,
	//	}, "登录成功", c)
	//	return
	//}

	//if err, jwtStr := jwtService.GetRedisJWT(user.Username); err == redis.Nil {
	//	if err := jwtService.SetRedisJWT(token, user.Username); err != nil {
	//		global.GVA_LOG.Error("设置登录状态失败!", zap.Error(err))
	//		response.FailWithMessage("设置登录状态失败", c)
	//		return
	//	}
	//	response.OkWithDetailed(systemRes.LoginResponse{
	//		User:      user,
	//		Token:     token,
	//		ExpiresAt: claims.StandardClaims.ExpiresAt * 1000,
	//	}, "登录成功", c)
	//} else if err != nil {
	//	global.GVA_LOG.Error("设置登录状态失败!", zap.Error(err))
	//	response.FailWithMessage("设置登录状态失败", c)
	//} else {
	//	var blackJWT system.JwtBlacklist
	//	blackJWT.Jwt = jwtStr
	//	if err := jwtService.JsonInBlacklist(blackJWT); err != nil {
	//		response.FailWithMessage("jwt作废失败", c)
	//		return
	//	}
	//	if err := jwtService.SetRedisJWT(token, user.Username); err != nil {
	//		response.FailWithMessage("设置登录状态失败", c)
	//		return
	//	}
	return nil, &kvaAuth.UserLoginRes{
		Token:     token,
		ExpiresAt: string(claims.StandardClaims.ExpiresAt * 1000),
	}, "登录成功"
}

