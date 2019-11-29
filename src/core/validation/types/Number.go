package types

import (
	"core/validation/constants"
	"strconv"
)

type Number struct {
	Min      int64
	Max      int64
	Required bool
}

func (num *Number) Validate(field interface{}, key string, value interface{}, isRequiredDefault bool) (bool, string) {
	associatedField, _ := field.(Number)

	if value == nil && (associatedField.Required || isRequiredDefault) {
		return false, key + constants.Required
	} else if value == nil {
		return true, "ok"
	}

	associatedValue, ok := value.(int64)

	if !ok {
		return false, key + constants.WrongType + "number"
	}

	if associatedField.Min != 0 && associatedValue < associatedField.Min {
		return false, key + constants.WrongMaxLength + strconv.FormatInt(int64(associatedField.Min), 10)
	}

	if associatedField.Max != 0 && associatedValue > associatedField.Max {
		return false, key + constants.WrongMinLength + strconv.FormatInt(int64(associatedField.Max), 10)
	}

	return true, "ok"
}
