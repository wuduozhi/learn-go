package main

import (
	"fmt"
	"smtp-server/models"
	"smtp-server/smtp"
)

func main() {
	authHandle := func(username, password, remoteAddress string) error {
		user := models.Login(username,password)

		if user.Email != ""{
			return nil
		}

		return smtp.ErrorAuthError

	}


	srv := &smtp.Server{
		Name: "mail.my.server",
		Addr:        ":25025",
		MaxBodySize: 5 * 1024,
		Handler:     smtp.StoreEmail,
		Auth :    authHandle,
	}
	fmt.Println(srv.ListenAndServe())
}