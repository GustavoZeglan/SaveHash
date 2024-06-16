package utils

import (
	"encoding/json"
	"github.com/go-playground/validator"
)

func ErrorHandler(body interface{}) ([]byte, error) {
	var errors []ReqError
	validate := validator.New()

	err := validate.Struct(body)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var el ReqError
			el.Field = err.Field()
			el.Tag = err.Tag()
			el.Value = err.Param()
			errors = append(errors, el)
		}

		errorsMap := make(map[string][]ReqError)
		errorsMap["errors:"] = errors
		errMap, _ := json.Marshal(errorsMap)

		return errMap, err
	}

	return nil, nil
}
