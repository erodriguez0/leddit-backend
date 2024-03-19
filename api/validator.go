package api

import (
	"github.com/erodriguez0/leddit-backend/util"
	"github.com/go-playground/validator/v10"
)

var validUserRole validator.Func = func(fieldLevel validator.FieldLevel) bool {
	if role, ok := fieldLevel.Field().Interface().(string); ok {
		return util.IsSupportedRole(role)
	}

	return false
}
