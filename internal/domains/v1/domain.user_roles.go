package v1

import (
	"context"
	"time"
)

type UserRoleDomain struct {
	UserId string
	RoleId int

	Id        string
	IsDeleted bool
	CreatedAt time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time
}

type UserRolesAction interface {
	Delete(ctx context.Context, inDom *UserRoleDomain) (outDom UserRoleDomain, statusCode int, err error)
	Store(ctx context.Context, inDom *UserRoleDomain) (outDom UserRoleDomain, statusCode int, err error)
}

type UserRolesRepository interface {
	Delete(ctx context.Context, inDom *UserRoleDomain) (err error)
	GetRolesByUser(ctx context.Context, inDom *UserDomain) (outDomain []RoleDomain, err error)
	GetUsersByRole(ctx context.Context, inDom *RoleDomain) (outDomain []UserRoleDomain, err error)
	Store(ctx context.Context, inDom *UserDomain) (err error)
	GetById(ctx context.Context, id string) (outDomain UserRoleDomain, err error)
}
