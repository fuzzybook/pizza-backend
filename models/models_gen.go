// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package models

import (
	"fmt"
	"io"
	"strconv"
)

type Menu struct {
	Ingredients []*MenuIngredient `json:"Ingredients"`
	Doughs      []*MenuDough      `json:"Doughs"`
	Condiments  []*MenuCondiment  `json:"Condiments"`
	Categories  []*MenuCategory   `json:"Categories"`
}

type UserPagesResponse struct {
	Page     int     `json:"page"`
	PageSize int     `json:"pageSize"`
	Count    int     `json:"count"`
	Users    []*User `json:"users,omitempty"`
}

type UserRole string

const (
	UserRoleAdmin        UserRole = "ADMIN"
	UserRoleUser         UserRole = "USER"
	UserRoleMusicmanager UserRole = "MUSICMANAGER"
)

var AllUserRole = []UserRole{
	UserRoleAdmin,
	UserRoleUser,
	UserRoleMusicmanager,
}

func (e UserRole) IsValid() bool {
	switch e {
	case UserRoleAdmin, UserRoleUser, UserRoleMusicmanager:
		return true
	}
	return false
}

func (e UserRole) String() string {
	return string(e)
}

func (e *UserRole) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = UserRole(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid UserRole", str)
	}
	return nil
}

func (e UserRole) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type UserStatus string

const (
	UserStatusAwaiting UserStatus = "AWAITING"
	UserStatusPending  UserStatus = "PENDING"
	UserStatusActive   UserStatus = "ACTIVE"
	UserStatusBlocked  UserStatus = "BLOCKED"
	UserStatusBanned   UserStatus = "BANNED"
)

var AllUserStatus = []UserStatus{
	UserStatusAwaiting,
	UserStatusPending,
	UserStatusActive,
	UserStatusBlocked,
	UserStatusBanned,
}

func (e UserStatus) IsValid() bool {
	switch e {
	case UserStatusAwaiting, UserStatusPending, UserStatusActive, UserStatusBlocked, UserStatusBanned:
		return true
	}
	return false
}

func (e UserStatus) String() string {
	return string(e)
}

func (e *UserStatus) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = UserStatus(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid UserStatus", str)
	}
	return nil
}

func (e UserStatus) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type UserType string

const (
	UserTypeSystem UserType = "SYSTEM"
	UserTypeSite   UserType = "SITE"
	UserTypeAuthor UserType = "AUTHOR"
)

var AllUserType = []UserType{
	UserTypeSystem,
	UserTypeSite,
	UserTypeAuthor,
}

func (e UserType) IsValid() bool {
	switch e {
	case UserTypeSystem, UserTypeSite, UserTypeAuthor:
		return true
	}
	return false
}

func (e UserType) String() string {
	return string(e)
}

func (e *UserType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = UserType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid UserType", str)
	}
	return nil
}

func (e UserType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
