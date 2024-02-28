package records

import (
	V1Domains "github.com/dmzsz/duozhuayu/internal/domains/v1"
)

func (r *Roles) ToV1Domain() V1Domains.RoleDomain {
	return V1Domains.RoleDomain{
		CreatedAt:   r.CreatedAt,
		Description: r.Description,
		Id:          r.Id,
		ParentId:    r.ParentId,
		PositionId:  r.PositionId,
		UpdatedAt:   r.UpdatedAt,
	}
}

func FromRolesV1Domain(r *V1Domains.RoleDomain) Roles {
	return Roles{
		CreatedAt:   r.CreatedAt,
		Description: r.Description,
		Id:          r.Id,
		ParentId:    r.ParentId,
		PositionId:  r.PositionId,
		UpdatedAt:   r.UpdatedAt,
	}
}

func ToArrayOfUsersV1Domain(u *[]Users) []V1Domains.UserDomain {
	var result []V1Domains.UserDomain

	for _, val := range *u {
		result = append(result, val.ToV1Domain())
	}

	return result
}
