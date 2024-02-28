package v1

import (
	"context"
	"time"
)

type RoleDomain struct {
	Name        string
	Description *string
	ParentId    *string
	PositionId  *string

	Id        string
	IsDeleted bool
	CreatedAt time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time
}

type RoleAction interface {
	Delete(ctx context.Context, inDom *RoleDomain) (outDom RoleDomain, statusCode int, err error)
	GetById(ctx context.Context, id string) (outDom RoleDomain, statusCode int, err error)
	Store(ctx context.Context, inDom *RoleDomain) (outDom RoleDomain, statusCode int, err error)
}

type RoleRepository interface {
	Delete(ctx context.Context, inDom *RoleDomain) (err error)
	GetById(ctx context.Context, id string) (outDomain RoleDomain, err error)
	Store(ctx context.Context, inDom *RoleDomain) (err error)
}
