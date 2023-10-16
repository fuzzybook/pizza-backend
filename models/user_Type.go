package models

import (
	"database/sql/driver"
)

func (e *UserType) Scan(value interface{}) error {
	userType := value.(string)
	*e = ""
	var t UserType = ""
	err := t.UnmarshalGQL(userType)
	if err != nil {
		t = "INVALID"
	}
	*e = t

	return nil
}

func (e UserType) Value() (driver.Value, error) {
	if !e.IsValid() {
		return "INVALID", nil
	}
	return e.String(), nil
}
