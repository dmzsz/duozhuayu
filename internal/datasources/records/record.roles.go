package records

import (
	"time"
)

type Roles struct {
	CreatedAt   time.Time  `db:"created_at"`
	DeletedAt   *time.Time `db:"deleted_at"`
	Description *string    `db:"description"`
	Id          string     `db:"id"`
	IsDeleted   bool       `db:"is_deleted"`
	Name        string     `db:"name"`
	ParentId    *string    `db:"parent_id"`
	PositionId  *string    `db:"position_id"`
	UpdatedAt   *time.Time `db:"updated_at"`
}
