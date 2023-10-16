package models

import (
	"database/sql/driver"
	"strings"

	"golang.org/x/exp/slices"
)

type UserRoles []UserRole

func (e *UserRoles) Scan(value interface{}) error {
	roles := strings.Split(value.(string), ",")
	*e = []UserRole{}
	for _, v := range roles {
		var r UserRole = ""
		err := r.UnmarshalGQL(v)
		if err != nil {
			return err
		}
		*e = append(*e, r)
	}
	return nil
}

func (j UserRoles) Value() (driver.Value, error) {
	if len(j) == 0 {
		return nil, nil
	}
	return strings.Join(j.Strings(), ","), nil
}

func (e UserRoles) Strings() []string {
	result := []string{}
	for _, v := range e {
		result = append(result, v.String())
	}
	return result
}

func (e UserRoles) CheckRoles(roles []string) bool {
	permissions := []string{}
	for _, v := range e {
		permissions = append(permissions, v.String())
	}
	for _, v := range roles {
		if slices.Contains(permissions, v) {
			return true
		}
	}
	return false
}
