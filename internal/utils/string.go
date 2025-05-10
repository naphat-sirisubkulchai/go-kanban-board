package utils

import (
	"regexp"
)
func IsEmailValid(email string) bool {
	// Regular expression for basic email validation
	re := regexp.MustCompile(`^[a-z0-9._%+-]+@[a-z0-9.-]+\.[a-z]{2,}$`)
	return re.MatchString(email)
}
