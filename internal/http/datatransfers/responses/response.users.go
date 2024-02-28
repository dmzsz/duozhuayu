package responses

import (
	"time"

	V1Domains "github.com/dmzsz/duozhuayu/internal/domains/v1"
)

type RoleId struct {
	Id string
}

type UserResponse struct {
	Id        string     `json:"id"`
	Username  string     `json:"username"`
	Email     string     `json:"email"`
	Password  string     `json:"password,omitempty"`
	RoleIds   []RoleId   `json:"role_ids"`
	Token     string     `json:"token,omitempty"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

func (u *UserResponse) ToV1Domain() V1Domains.UserDomain {
	return V1Domains.UserDomain{
		Id:        u.Id,
		Username:  u.Username,
		Password:  u.Password,
		Email:     u.Email,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}

func FromV1Domain(u V1Domains.UserDomain) UserResponse {
	var RoleIds []RoleId
	for _, role := range *u.Roles {
		RoleIds = append(RoleIds, RoleId{Id: role.Id})
	}
	return UserResponse{
		Id:        u.Id,
		Username:  u.Username,
		Email:     u.Email,
		Password:  u.Password,
		Token:     u.ResetPasswordToken,
		RoleIds:   RoleIds,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}

func ToResponseList(domains []V1Domains.UserDomain) []UserResponse {
	var result []UserResponse

	for _, val := range domains {
		result = append(result, FromV1Domain(val))
	}

	return result
}
