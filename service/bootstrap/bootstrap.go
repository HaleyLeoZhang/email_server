package bootstrap

import (
	"email_server/service/email_service"
	"fmt"
)

func Setup() {
	email := &email_service.Email{}

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("捕获异常:", err)
		}
	}()

	go email.DoPull()
}
