package records

import (
	"time"
)

type Users struct {
	Avatar                 string     `db:"avatar"`
	balance                string     `db:"balance"`
	CreatedAt              time.Time  `db:"created_at"`
	DeletedAt              *time.Time `db:"deleted_at"`
	Email                  string     `db:"email"`
	EmailNonce             string     `db:"email_nonce"`
	EmailCipher            string     `db:"email_cipher"`
	FirstName              string     `db:"password"`
	Gender                 string     `db:"gender"`
	Id                     string     `db:"id"`
	IsActive               bool       `db:"is_active"`
	IsDeleted              bool       `db:"is_deleted"`
	IsLocked               bool       `db:"is_locked"`
	IsOnline               bool       `db:"is_online"`
	IsVerified             string     `db:"is_verified"`
	LastName               string     `db:"last_name"`
	Level                  string     `db:"level"`
	Nickname               string     `db:"nickname"`
	Password               string     `db:"password"`
	PhoneNumber            string     `db:"phone_number"`
	Reason                 string     `db:"reason"`
	ResetPasswordExpiresIn string     `db:"reset_password_expires_in"`
	ResetPasswordToken     string     `db:"reset_password_token"`
	SessionKey             string     `db:"session_key"`
	Unionid                string     `db:"unionid"`
	UpdatedAt              *time.Time `db:"updated_at"`
	Username               string     `db:"username"`
	WXAccessToken          string     `db:"wx_access_token"`
	WXExpiresInstring      string     `db:"wx_expires_in"`
	WXNickname             string     `db:"wx_nickname"`
	WXOpenid               string     `db:"wx_openid"`
	WXRefreshToken         string     `db:"wx_refresh_token"`
	WXHeadimgurl           string     `db:"wx_headimgurl"`
}
