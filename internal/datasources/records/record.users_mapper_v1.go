package records

import (
	V1Domains "github.com/dmzsz/duozhuayu/internal/domains/v1"
)

func (u *Users) ToV1Domain() V1Domains.UserDomain {
	return V1Domains.UserDomain{
		Id:        u.Id,
		Username:  u.Username,
		Email:     u.Email,
		Password:  u.Password,
		IsActive:  u.IsActive,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
		Roles:     &[]V1Domains.RoleDomain{},
	}
}

func FromUsersV1Domain(u *V1Domains.UserDomain) Users {
	return Users{
		Id:        u.Id,
		Username:  u.Username,
		Email:     u.Email,
		Password:  u.Password,
		IsActive:  u.IsActive,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}

func ToArrayOfRolesV1Domain(r *[]Roles) []V1Domains.RoleDomain {
	var result []V1Domains.RoleDomain

	for _, val := range *r {
		result = append(result, val.ToV1Domain())
	}

	return result
}
