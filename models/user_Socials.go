package models

import (
	"database/sql/driver"
	"encoding/json"
)

type UserSocials struct {
	Website   string `json:"socials"`
	Twitter   string `json:"twitter"`
	Linkedin  string `json:"linkedin"`
	Instagram string `json:"instagram"`
	Facebook  string `json:"facebook"`
	Youtube   string `json:"youtube"`
}

func (e *UserSocials) Scan(value interface{}) error {
	if err := json.Unmarshal(value.([]byte), e); err != nil {
		return nil
	}
	return nil
}

func (e UserSocials) Value() (driver.Value, error) {
	u, err := json.Marshal(e)
	if err != nil {
		return nil, nil
	}
	return u, nil
}
