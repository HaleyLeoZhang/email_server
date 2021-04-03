package util

import (
	"fmt"
	"github.com/google/uuid"
)

func GetUuid() string {
	u1 := uuid.Must(uuid.NewRandom())
	return fmt.Sprintf("%s", u1)
}
