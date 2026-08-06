package validator

import (
	"pm-with-essence/internal/api/app_errors"
	"unicode"
)

func ValidatePassword(password string) *app_errors.HttpError {
	if len(password) < 8 || len(password) > 72 {
		return app_errors.NewBadRequestError(
			"Password length must be between 8 and 72 characters",
		)
	}

	var hasUpper, hasLower, hasDigit bool

	for _, r := range password {
		switch {
		case unicode.IsUpper(r):
			hasUpper = true
		case unicode.IsLower(r):
			hasLower = true
		case unicode.IsDigit(r):
			hasDigit = true
		}
	}

	if !hasUpper || !hasLower || !hasDigit {
		return app_errors.NewBadRequestError(
			"Password must contain uppercase, lowercase and digit",
		)
	}

	return nil
}
