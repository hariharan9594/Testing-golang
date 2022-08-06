package model

import "net/smtp"

func SendMail() {

	from := "hariharanTesting@yahoo.com"
	password := "Green005@1995"

	toEmailAddress := "hariharan9594py@gmail.com"
	to := []string{toEmailAddress}

	host := "smtp.mail.yahoo.com"
	port := "465"
	address := host + ":" + port

	subject := "Subject: This is the subject of the mail\n"
	body := "This is the body of the mail"
	message := []byte(subject + body)

	auth := smtp.PlainAuth("", from, password, host)

	err := smtp.SendMail(address, auth, from, to, message)
	if err != nil {
		panic(err)
	}

}
