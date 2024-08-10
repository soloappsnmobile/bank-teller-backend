package validators

import (
	"fmt"
	"regexp"
)

func ValidateRequest(req map[string]interface{}, req_type string) error {
	var requiredFields []string
	if req_type == "Teller" {
		requiredFields = []string{"first_name", "last_name", "phone", "email", "password"}
	} else if req_type == "Customer" {
		requiredFields = []string{"first_name", "last_name", "phone", "email", "password"}
	} else {
		return fmt.Errorf("invalid request type")
	}

	for _, field := range requiredFields {

		// For other fields, check if they are strings
		if _, ok := req[field].(string); !ok {
			return fmt.Errorf("%s is required", field)
		}

	}

	// if field == "dob" {
	// 		// Special handling for dob
	// 		dobStr, ok := req[field].(string)
	// 		if !ok {
	// 			return fmt.Errorf("%s is required", field)
	// 		}
	// 		_, err := time.Parse("2006-01-02", dobStr)
	// 		if err != nil {
	// 			return fmt.Errorf("invalid dob format, expected YYYY-MM-DD")
	// 		}
	// 	} else {

	// Validate email
	email, ok := req["email"].(string)
	if !ok {
		return fmt.Errorf("email is required")
	}
	if !IsValidEmail(email) {
		return fmt.Errorf("invalid email format")
	}

	// // Additional check for address field when req_type is "Customer"
	// if req_type == "Customer" {
	// 	address, ok := req["address"].(string)
	// 	if !ok || address == "" {
	// 		return fmt.Errorf("address is required and cannot be empty")
	// 	}
	// }

	return nil
}

func IsValidEmail(email string) bool {
	emailRegex := `^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`
	match := regexp.MustCompile(emailRegex).MatchString
	return match(email)
}
