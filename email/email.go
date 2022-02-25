package email

import (
	"crypto/tls"
	"fmt"
	"github.com/cuwand/pondasi/enum/environtmentEnums"
	"github.com/cuwand/pondasi/logger"
	"gopkg.in/gomail.v2"
)

type EmailConfig struct {
	emailDialler *gomail.Dialer
	logger       logger.Logger
	mailSender   string
	mailTrap     string
	environment  environtmentEnums.Environment
}

type SendMail struct {
	Subject string
	Content string
	From    string
	To      []string
	Cc      []string
	Bcc     []string
}

func InitConfig(mailPort int, mailHost, mailSender, username, password, mailTrap string, logger logger.Logger, tlsStatus bool, env environtmentEnums.Environment) EmailConfig {
	emailDialler := gomail.NewDialer(mailHost,
		mailPort,
		username,
		password)

	emailDialler.TLSConfig = &tls.Config{InsecureSkipVerify: tlsStatus}

	return EmailConfig{
		mailSender:   mailSender,
		emailDialler: emailDialler,
		logger:       logger,
		environment:  env,
		mailTrap:     mailTrap,
	}
}

func (cfg EmailConfig) SendMail(payload SendMail) error {
	if cfg.environment != environtmentEnums.Production {
		payload.To = []string{cfg.mailTrap}
	}

	mailer := gomail.NewMessage()
	mailer.SetHeader("From", fmt.Sprintf("%s <%s>", payload.From, cfg.mailSender))
	mailer.SetHeader("To", payload.To...)
	mailer.SetHeader("Cc", payload.Cc...)
	mailer.SetHeader("Bcc", payload.Bcc...)
	mailer.SetHeader("Subject", payload.Subject)
	mailer.SetBody("text/html", payload.Content)

	err := cfg.emailDialler.DialAndSend(mailer)

	if err != nil {
		//cfg.logger.Error(fmt.Sprintf("Mail Cannot Sent!! Error:: %s || Subject:: %s", err.Error(), payload.Subject))
		return err
	}

	cfg.logger.Info(fmt.Sprintf("Mail Sent!! Subject:: %s", payload.Subject))

	return nil
}
