package utils

import (
	"regexp"
	"strings"
)

func ValidateEmail(email string) bool {
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return emailRegex.MatchString(email)
}

func ValidatePassword(password string) []ValidationError {
	var errors []ValidationError

	if len(password) < 6 {
		errors = append(errors, ValidationError{
			Field:   "password",
			Message: "Password must be at least 6 characters long",
		})
	}

	if len(password) > 128 {
		errors = append(errors, ValidationError{
			Field:   "password",
			Message: "Password must be less than 128 characters",
		})
	}

	return errors
}

func ValidateJournalContent(content string) []ValidationError {
	var errors []ValidationError

	content = strings.TrimSpace(content)
	if len(content) == 0 {
		errors = append(errors, ValidationError{
			Field:   "content",
			Message: "Journal content cannot be empty",
		})
	}

	if len(content) > 10000 {
		errors = append(errors, ValidationError{
			Field:   "content",
			Message: "Journal content must be less than 10,000 characters",
		})
	}

	return errors
}

func SanitizeInput(input string) string {
	return strings.TrimSpace(input)
}