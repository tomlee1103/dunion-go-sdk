package util

import (
	"github.com/google/uuid"
)

func uuid4() string {
	return uuid.New().String()
}
