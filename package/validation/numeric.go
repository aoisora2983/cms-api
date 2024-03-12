package validation

import (
	"errors"
	"strconv"
)

func Numeric(value interface{}) error {
	numeric, _ := value.(string)

	if _, err := strconv.Atoi(numeric); err == nil {
		return errors.New("半角数値を入力してください。")
	}

	return nil
}
