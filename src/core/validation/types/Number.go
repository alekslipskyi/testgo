package types

import (
	"core/validation/constants"
	"strconv"
)

type Number struct {
	Min      float64
	Max      float64
	Required bool
}

func (num *Number) Validate(field interface{}, key string, value interface{}) (bool, string) {
	if value == nil {
		return false, key + constants.Required
	}

	associatedField, _ := field.(Number)
	associatedValue, ok := value.(float64)

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
