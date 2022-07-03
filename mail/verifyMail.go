package mail

import (
	"fmt"
	"net/smtp"
)

type smtpServer struct {
	host string
	port string
}

func (s *smtpServer) Address() string {
	return s.host + ":" + s.port
}

func SendMail(toMail string, OTP string) (valid bool) {

	// Sender data.
	from := "madanthapa0707@gmail.com"

	password := "mpwgxpohnwblqbed"
	// Receiver email address.
	to := []string{toMail}
	authCodeMessage := "the authentication code is" + OTP
	message := []byte(authCodeMessage)
	// smtp server configuration.
	smtpServer := smtpServer{host: "smtp.gmail.com", port: "587"}
	// Message.

	// Authentication
	auth := smtp.PlainAuth("", from, password, smtpServer.host)
	// Sending email.
	err := smtp.SendMail(smtpServer.Address(), auth, from, to, message)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("mail is done")
	return true
}
