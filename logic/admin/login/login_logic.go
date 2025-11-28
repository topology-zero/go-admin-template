package login

import (
	"strconv"
	"strings"

	"go-admin-template/config"
	"go-admin-template/pkg/jwt"
	"go-admin-template/pkg/redis"
	"go-admin-template/query"
	"go-admin-template/svc"
	"go-admin-template/types"

	"github.com/pkg/errors"
	"github.com/wenlng/go-captcha/v2/slide"
	"golang.org/x/crypto/bcrypt"
)

// Login 登录
func Login(ctx *svc.ServiceContext, req *types.LoginRequest) (resp types.LoginResponse, err error) {
	captchaKey := config.ServerConf.Name + ":CAPTCHAKEY:" + req.CodeID
	dbCode, _ := redis.Client.Get(ctx, captchaKey).Result()
	if len(dbCode) == 0 {
		err = errors.New("验证码已过期")
		return
	}
	redis.Client.Del(ctx, captchaKey) // 只允许用一次
	userCodeArr := strings.Split(req.Code, ",")
	dbCodeArr := strings.Split(dbCode, ",")

	userCodeX, _ := strconv.Atoi(userCodeArr[0])
	userCodeY, _ := strconv.Atoi(userCodeArr[1])
	dbCodeX, _ := strconv.Atoi(dbCodeArr[0])
	dbCodeY, _ := strconv.Atoi(dbCodeArr[1])

	if !slide.Validate(userCodeX, userCodeY, dbCodeX, dbCodeY, 5) {
		err = errors.New("验证码错误")
		return
	}

	userModel := query.AdminUserModel
	userInfo, _ := userModel.WithContext(ctx).Where(userModel.Username.Eq(req.Username)).First()
	if userInfo == nil {
		err = errors.New("用户名或密码错误")
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(userInfo.Password), []byte(req.Password))
	if err != nil {
		err = errors.New("用户名或密码错误")
		return
	}

	if userInfo.Status != 1 {
		err = errors.New("该用户已被管理员停用")
		return
	}

	token, err := jwt.MakeToken(userInfo.ID, userInfo.RoleID, userInfo.Phone)
	if err != nil {
		err = errors.New("系统错误")
		return
	}
	resp.Jwt = token
	return
}
