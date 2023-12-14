package scalars

import (
	"encoding/json"
	"fmt"
	"strconv"
)

type F64 float64

func (f *F64) UnmarshalGQL(v interface{}) error {
	fmt.Printf("%T ", v)
	return nil
}

func (f *F64) UnmarshalFloat(v interface{}) (float64, error) {
	switch v := v.(type) {
	case string:
		return strconv.ParseFloat(v, 64)
	case int:
		return float64(v), nil
	case int64:
		return float64(v), nil
	case float64:
		return v, nil
	case json.Number:
		return strconv.ParseFloat(string(v), 64)
	default:
		return 0, fmt.Errorf("%T is not an float", v)
	}
}
