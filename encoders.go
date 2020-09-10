package dyanmic_params

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

const ErrCnvFailed = "conversion failed"

func convertToString(val interface{}) (string, error) {
	if v, ok := val.(string); ok {
		return v, nil
	} else if v, ok := val.(*string); ok {
		return *v, nil
	}
	return "", errors.New(ErrCnvFailed)
}

func convertToBytes(val interface{}) ([]byte, error) {
	if v, ok := val.([]byte); ok {
		return v, nil
	} else if v, ok := val.(*[]byte); ok {
		return *v, nil
	}
	return nil, errors.New(ErrCnvFailed)
}
func convertToInt(val interface{}) (int, error) {
	if v, ok := val.(int); ok {
		return v, nil
	} else if v, ok := val.(*int); ok {
		return *v, nil
	}
	return 0, errors.New(ErrCnvFailed)
}


func convertNumericStrToInt(val interface{}) (int, error) {
	str, err := convertToString(val)
	if err != nil {
		return 0, err
	}
	if strings.Index(str, "0") == 0 {
		return 0, errors.New("numeric string starts with zero")
	}
	rg := regexp.MustCompile(`^[0-9]+$`)
	if rg.Match([]byte(str)) {
		numInt, err := strconv.Atoi(str)
		if err != nil {
			return 0, err
		}
		return numInt, nil
	}
	return 0, errors.New(ErrCnvFailed)
}

func convertNumericStrToBool(val interface{}) (bool, error) {
	str, err := convertToString(val)
	if err != nil {
		return false, err
	}
	if str == "0" || str == "false" {
		return false, nil
	} else if str == "1" || str == "true" {
		return true, nil
	}
	return false, errors.New(ErrCnvFailed)
}


func convertToBool(val interface{}) (bool, error) {
	if v, ok := val.(bool); ok {
		return v, nil
	} else if v, ok := val.(*bool); ok {
		return *v, nil
	}
	return false, errors.New(ErrCnvFailed)
}
func convertToInt32(val interface{}) (int32, error) {
	if v, ok := val.(int32); ok {
		return v, nil
	} else if v, ok := val.(*int32); ok {
		return *v, nil
	}
	return 0, errors.New(ErrCnvFailed)
}
func convertToInt64(val interface{}) (int64, error) {
	if v, ok := val.(int64); ok {
		return v, nil
	} else if v, ok := val.(*int64); ok {
		return *v, nil
	}
	return 0, errors.New(ErrCnvFailed)
}

func convertToInt8(val interface{}) (int8, error) {
	if v, ok := val.(int8); ok {
		return v, nil
	} else if v, ok := val.(*int8); ok {
		return *v, nil
	}
	return 0, errors.New(ErrCnvFailed)
}
func convertToInt16(val interface{}) (int16, error) {
	if v, ok := val.(int16); ok {
		return v, nil
	} else if v, ok := val.(*int16); ok {
		return *v, nil
	}
	return 0, errors.New(ErrCnvFailed)
}