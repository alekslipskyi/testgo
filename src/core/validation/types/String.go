package types

import (
	"core/validation/constants"
	"strconv"
)

type String struct {
	Min      int
	Max      int
	Required bool
}

func (str *String) Validate(field interface{}, key string, value interface{}, isRequiredDefault bool) (bool, string) {
	associatedField, _ := field.(String)

	if value == nil && (associatedField.Required || isRequiredDefault) {
		return false, key + constants.Required
	} else if value == nil {
		return true, "ok"
	}

	associatedValue, ok := value.(string)

	if !ok {
		return false, key + constants.WrongType + "string"
	}

	if associatedField.Min != 0 && len(associatedValue) < associatedField.Min {
		return false, key + constants.WrongMinLength + strconv.FormatInt(int64(associatedField.Min), 10)
	}

	if associatedField.Max != 0 && len(associatedValue) > associatedField.Max {
		return false, key + constants.WrongMaxLength + strconv.FormatInt(int64(associatedField.Max), 10)
	}

	return true, "ok"
}
