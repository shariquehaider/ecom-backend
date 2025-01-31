package utils

import "regexp"

func IsValidEmail(email string) bool {
	// Regular expression for basic email validation
	regex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	re := regexp.MustCompile(regex)
	return re.MatchString(email)
}

func VerifyNewPassword(password, confirmPassword string) bool {
	return password == confirmPassword
}
