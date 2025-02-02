package utils

import (
	"regexp"
	"strings"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func IsValidEmail(email string) bool {
	// Regular expression for basic email validation
	regex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	re := regexp.MustCompile(regex)
	return re.MatchString(email)
}

func VerifyNewPassword(password, confirmPassword string) bool {
	return password == confirmPassword
}

func VerifyObjectId(id string) (*primitive.ObjectID, error) {
	trimmedObjectId := strings.TrimPrefix(id, "ObjectID(")
	trimmedObjectId = strings.TrimSuffix(trimmedObjectId, ")")
	trimmedObjectId = strings.Trim(trimmedObjectId, `"`)

	objID, err := primitive.ObjectIDFromHex(trimmedObjectId)
	if err != nil {
		return nil, err
	}

	return &objID, nil
}
