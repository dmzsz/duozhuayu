package v1

import (
	"context"
	"time"
)

type UserDomain struct {
	Roles *[]RoleDomain

	Avatar               string
	balance              string
	Email                string
	FirstName            string
	Gender               string
	IsActive             bool
	IsLocked             bool
	IsOnline             bool
	IsVerified           string
	LastName             string
	Level                string
	openid               string
	Password             string
	PhoneNumber          string
	Reason               string
	ResetPasswordExpires string
	ResetPasswordToken   string
	SessionToken         string
	AccessToken          string
	SessionKey           string
	Unionid              string
	Username             string

	WXAccessToken     string
	WXExpiresInstring string
	WXNickname        string
	WXOpenid          string
	WXRefreshToken    string
	WXHeadimgurl      string

	Id        string
	IsDeleted bool
	CreatedAt time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time
}

type UserUsecase interface {
	// 激活用户
	ActivateUser(ctx context.Context, email string) (statusCode int, err error)
	Delete(ctx context.Context, inDom *UserDomain) (statusCode int, err error)
	GetByEmail(ctx context.Context, inDom *UserDomain, decryptEmail bool) (outDom UserDomain, statusCode int, err error)
	Login(ctx context.Context, inDom *UserDomain) (outDom UserDomain, statusCode int, err error)
	// 发送验证码
	SendOTP(ctx context.Context, email string) (otpCode string, statusCode int, err error)
	Store(ctx context.Context, inDom *UserDomain) (outDom UserDomain, statusCode int, err error)
	// 校验验证码
	VerifOTP(ctx context.Context, email string, userOTP string, otpRedis string) (statusCode int, err error)
}
