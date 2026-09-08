package main

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/sodieSam/podcast/23_packages/auth"
	"github.com/sodieSam/podcast/23_packages/user"
)

func main() {
	auth.LoginWithCredentials("Samrat", "123")

	session := auth.GetSession()
	fmt.Println("session : ", session)

	user := user.User{
		Email: "user@email.com",
		Name:  "John Doe",
	}
	// fmt.Println(user.Email, user.Name)

	color.Red(user.Name, user.Email)

}
