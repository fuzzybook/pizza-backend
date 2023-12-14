package models

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

type MenuDataCategory struct {
	Text     string `json:"text"`
	Priority int    `json:"priority"`
}

type MenuDataCategoies []MenuDataCategory

func (e *MenuDataCategoies) Scan(value interface{}) error {
	v := MenuDataCategoies{}
	err := json.Unmarshal([]byte(value.(string)), &v)
	if err != nil {
		return err
	}
	*e = v
	if len([]rune(value.(string))) > 3 {
		fmt.Println("unmarshal", e)
	}
	return nil
}

func (j MenuDataCategoies) Value() (driver.Value, error) {
	m, err := json.Marshal(&j)
	if err != nil {
		return m, err
	}
	return m, nil
}
