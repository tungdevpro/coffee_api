package entity

import (
	"coffee_api/commons"
	"database/sql/driver"
)

type StatusAllowed string

const (
	Suspended StatusAllowed = "suspended"
	Active    StatusAllowed = "active"
	Inactive  StatusAllowed = "inactive"
)

func (r *StatusAllowed) Scan(value interface{}) error {
	*r = StatusAllowed(value.([]byte))
	return nil
}

func (r StatusAllowed) Value() (driver.Value, error) {
	return string(r), nil
}

type RoleAllowed string

const (
	admin  RoleAllowed = "admin"
	seller RoleAllowed = "seller"
	rider  RoleAllowed = "rider"
	member RoleAllowed = "member"
)

func (st *RoleAllowed) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		*st = RoleAllowed(b)
	}
	return nil
}

func (st RoleAllowed) Value() (driver.Value, error) {
	return string(st), nil
}

type User struct {
	*commons.SQLModel `json:",inline"`
	FullName          string         `json:"fullname" gorm:"column:fullname;"`
	Email             string         `json:"email" gorm:"column:email;type:varchar(100);unique_index"`
	Phone             string         `json:"phone" gorm:"column:phone;"`
	Address           string         `json:"address" gorm:"column:address;"`
	Avatar            *commons.Image `json:"avatar" gorm:"column:avatar;"`
	Status            StatusAllowed  `json:"status" gorm:"column:status;type:ENUM('active','suspended','inactive');default:'active'"`
	Role              RoleAllowed    `json:"role" gorm:"column:role;type:ENUM('admin','seller','rider','member');default:'member'"`
	// OTPCode
	// IsEmailVerified bool           `json:"is_email_verified" gorm:"default:false"`

	AccessToken string `json:"access_token" gorm:"column:access_token;"`
}

func (u *User) TableName() string {
	return "users"
}
