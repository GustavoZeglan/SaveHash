package middleware

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/GustavoZeglan/SaveHash/web/utils"
	"github.com/go-playground/validator"
	"net/http"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

func ValidateBody[T any](next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var payload T
		var errors []string

		if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		if err := validate.Struct(payload); err != nil {
			for _, err := range err.(validator.ValidationErrors) {
				var el utils.ReqError
				el.Field = err.Field()
				el.Tag = err.Tag()
				el.Value = err.Param()
				errors = append(errors, fmt.Sprintf("Field validation for '%s' failed on the '%s' tag", err.Field(), err.Tag()))
			}

			utils.RespondWithError(w, http.StatusBadRequest, errors)
			return
		}

		ctx := context.WithValue(r.Context(), "payload", payload)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}
