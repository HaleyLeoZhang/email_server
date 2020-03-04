package util

import (
    "github.com/google/uuid"
    "fmt"
)

func GetUuid() string {
    u1 := uuid.Must(uuid.NewRandom())
    return fmt.Sprintf("%s", u1)
}
