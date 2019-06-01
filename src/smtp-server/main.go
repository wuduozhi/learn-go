package main

import (
	"smtp-server/models"
	"smtp-server/smtp"
	"smtp-server/pop3"
	"smtp-server/pkg/logging"
	"time"
)

func main() {
	/* 验证逻辑 */
	//AFd1ZHVvemhpQHFxLmNvbQB3ZHoxMjM=
	authHandle := func(username, password, remoteAddress string) error {
		user := models.Login(username,password)

		if user.Email != ""{
			return nil
		}

		return smtp.ErrorAuthError
	}

	saveHandle := func(request * smtp.Request) (error){
		from := request.From
		message := smtp.StreamToString(request.Message)
		
		tos := request.To

		for i := 0;i<len(request.To);i++ {
			email := models.Email{
				From:from,
				To :tos[i],
				Message:message,
				Create_at: int(time.Now().Unix()),
				Is_new:1,
			}

			err := email.Create()
			if err != nil {
				logging.Debug(err)
			}
		}

		return nil
	}

	smtp_srv := &smtp.Server{
		Name: "mail.my.server",
		Addr:        ":25025",
		MaxBodySize: 5 * 1024,
		Handler:     saveHandle,
		Auth :    authHandle,
	}

	go smtp_srv.ListenAndServe()

	pop_srv := &pop3.Server{
		Host:"wduozhi.xyz",
		Addr:":11011",
	}

	go pop_srv.ListenAndServe()

	for {

	}

}
