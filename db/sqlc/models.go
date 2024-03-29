// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package db

import (
	"database/sql"
	"database/sql/driver"
	"fmt"

	"github.com/google/uuid"
)

type UserRole string

const (
	UserRoleADMIN UserRole = "ADMIN"
	UserRoleMOD   UserRole = "MOD"
	UserRoleUSER  UserRole = "USER"
)

func (e *UserRole) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = UserRole(s)
	case string:
		*e = UserRole(s)
	default:
		return fmt.Errorf("unsupported scan type for UserRole: %T", src)
	}
	return nil
}

type NullUserRole struct {
	UserRole UserRole `json:"user_role"`
	Valid    bool     `json:"valid"` // Valid is true if UserRole is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullUserRole) Scan(value interface{}) error {
	if value == nil {
		ns.UserRole, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.UserRole.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullUserRole) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.UserRole), nil
}

type VoteType string

const (
	VoteTypeUP   VoteType = "UP"
	VoteTypeDOWN VoteType = "DOWN"
)

func (e *VoteType) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = VoteType(s)
	case string:
		*e = VoteType(s)
	default:
		return fmt.Errorf("unsupported scan type for VoteType: %T", src)
	}
	return nil
}

type NullVoteType struct {
	VoteType VoteType `json:"vote_type"`
	Valid    bool     `json:"valid"` // Valid is true if VoteType is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullVoteType) Scan(value interface{}) error {
	if value == nil {
		ns.VoteType, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.VoteType.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullVoteType) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.VoteType), nil
}

type Comment struct {
	ID        uuid.UUID     `json:"id"`
	Body      string        `json:"body"`
	PostID    uuid.UUID     `json:"post_id"`
	UserID    uuid.UUID     `json:"user_id"`
	ReplyID   uuid.NullUUID `json:"reply_id"`
	CreatedAt sql.NullTime  `json:"created_at"`
	UpdatedAt sql.NullTime  `json:"updated_at"`
}

type CommentVote struct {
	Vote      VoteType  `json:"vote"`
	CommentID uuid.UUID `json:"comment_id"`
	UserID    uuid.UUID `json:"user_id"`
}

type Post struct {
	ID          uuid.UUID      `json:"id"`
	Title       string         `json:"title"`
	Url         sql.NullString `json:"url"`
	Body        sql.NullString `json:"body"`
	UserID      uuid.NullUUID  `json:"user_id"`
	SubledditID uuid.NullUUID  `json:"subleddit_id"`
	CreatedAt   sql.NullTime   `json:"created_at"`
	UpdatedAt   sql.NullTime   `json:"updated_at"`
}

type PostImage struct {
	ID        uuid.UUID    `json:"id"`
	Url       string       `json:"url"`
	PostID    uuid.UUID    `json:"post_id"`
	CreatedAt sql.NullTime `json:"created_at"`
}

type PostVote struct {
	Vote   VoteType  `json:"vote"`
	PostID uuid.UUID `json:"post_id"`
	UserID uuid.UUID `json:"user_id"`
}

type Subleddit struct {
	ID        uuid.UUID     `json:"id"`
	Name      string        `json:"name"`
	UserID    uuid.NullUUID `json:"user_id"`
	CreatedAt sql.NullTime  `json:"created_at"`
	UpdatedAt sql.NullTime  `json:"updated_at"`
}

type User struct {
	ID        uuid.UUID      `json:"id"`
	Username  string         `json:"username"`
	Password  string         `json:"password"`
	Avatar    sql.NullString `json:"avatar"`
	Role      NullUserRole   `json:"role"`
	CreatedAt sql.NullTime   `json:"created_at"`
	UpdatedAt sql.NullTime   `json:"updated_at"`
}
