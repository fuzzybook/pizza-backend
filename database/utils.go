package database

import (
	"database/sql"
	"fmt"
	"strconv"
)

func valueToInt(data sql.RawBytes) int {
	v, err := strconv.Atoi(fmt.Sprintf("%s", data))
	if err != nil {
		return 0
	}
	return v
}

func valueToString(data sql.RawBytes) string {
	return fmt.Sprintf("%s", data)
}
