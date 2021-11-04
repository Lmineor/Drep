package utils

import (
	uuid "github.com/satori/go.uuid"
)

func GenerateUUID() string {
	uuid := uuid.NewV4()
	return uuid.String()
}
