package request

import (
	"errors"
	"fmt"
	"github.com/GustavoZeglan/SaveHash/web/utils"
	"github.com/go-playground/validator"
	"strings"
)

type Validator[T SignUpRequest | LoginRequest] struct {
	body T
}

func (v *Validator[T]) Validate() error {
	var validationErrors []string

	if err := validate.Struct(v.body); err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var el utils.ReqError
			el.Field = err.Field()
			el.Tag = err.Tag()
			el.Value = err.Param()
			validationErrors = append(validationErrors, fmt.Sprintf("Field validation for '%s' failed on the '%s' tag", err.Field(), err.Tag()))
		}
		return errors.New(strings.Join(validationErrors, ", "))
	}

	return nil
}
