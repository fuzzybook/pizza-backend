package models

import "database/sql/driver"

func (e *UserStatus) Scan(value interface{}) error {
	err := e.UnmarshalGQL(value)
	if err != nil {
		e.UnmarshalGQL("AWAITING")
		return nil
	}
	return nil
}

func (j UserStatus) Value() (driver.Value, error) {
	if !j.IsValid() {
		return "AWAITING", nil
	}
	return j.String(), nil
}
