package validators

import (
	"fmt"
	"regexp"
)

func ValidateRequest(req map[string]interface{}) error {
	requiredFields := []string{"first_name", "last_name", "phone", "email", "password"}

	for _, field := range requiredFields {
		if _, ok := req[field].(string); !ok {
			return fmt.Errorf("%s is required", field)
		}
	}
	// Validate email
	email, ok := req["email"].(string)
	if !ok {
		return fmt.Errorf("email is required")
	}
	if !IsValidEmail(email) {
		return fmt.Errorf("invalid email format")
	}

	return nil
}

func IsValidEmail(email string) bool {
	emailRegex := `^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`
	match := regexp.MustCompile(emailRegex).MatchString
	return match(email)
}
