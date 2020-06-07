package customStruct

import (
	"encoding/json"
	"errors"
	"strconv"
)

type CustomUint64 uint64

func (t *CustomUint64) UnmarshalJSON(b []byte) error {
	var data interface{}
	err := json.Unmarshal(b, &data)
	if err != nil {
		return err
	}
	switch data.(type) {
	case string:
		//d, err := strconv.Atoi(data.(string))
		d, err := strconv.ParseFloat(data.(string),32)
		if err != nil {
			*t = 0
		}
		*t = CustomUint64(d)
	case float64:
		*t = CustomUint64(data.(float64))
	case bool:
		if data.(bool) {
			*t = 1
		} else {
			*t = 0
		}
	default:
		*t = 0
	}
	return nil
}
func (t *CustomUint64) UnmarshalJSON1(b []byte) error {
	str := string(b)
	if len(str) < 1 {
		return errors.New("is zero")
	}
	if str[0] == '"' {
		var goString string
		err := json.Unmarshal(b, &goString)
		if err != nil {
			return err
		}
		if goString == "" {
			*t = 0
		} else {
			v, err := strconv.ParseUint(goString, 10, 32)
			if err != nil {
				return err
			}
			*t = CustomUint64(v)
		}
	} else if str == "false" {
		*t = 0
	} else if str == "true" {
		*t = 1
	} else if str == "null" {
		*t = 0
	} else {
		v, err := strconv.ParseFloat(str, 32)
		if err != nil {
			return err
		}
		*t = CustomUint64(v)
	}
	return nil
}
