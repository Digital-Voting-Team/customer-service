package helpers

import (
	"gitlab.com/distributed_lab/logan/v3/errors"
	"strconv"
	"time"
)

func IsInteger(value interface{}) error {
	if integer, ok := value.(*int64); ok {
		if *integer >= 0 {
			return nil
		}
	}

	if v, ok := value.(*string); ok {
		if integer, err := strconv.Atoi(*v); err == nil {
			if integer >= 0 {
				return nil
			}
			return errors.New("value is less or equal 0")
		}
		return errors.New("value is not a number")
	}

	return errors.New("unknown value type")
}

func IsDate(value interface{}) error {
	if _, ok := value.(**time.Time); ok {
		return nil
	}
	if _, ok := value.(*time.Time); ok {
		return nil
	}
	if _, ok := value.(time.Time); ok {
		return nil
	}
	return errors.New("value is not an valid date")
}
