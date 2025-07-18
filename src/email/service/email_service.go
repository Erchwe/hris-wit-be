package service

import (
	"context"
	"fmt"

	"github.com/go-gomail/gomail"
	"github.com/pkg/errors"
	"github.com/wit-id/blueprint-backend-go/common/httpservice"
	"github.com/wit-id/blueprint-backend-go/src/repository/payload"
	"github.com/wit-id/blueprint-backend-go/toolkit/config"
	"github.com/wit-id/blueprint-backend-go/toolkit/log"
)

type EmailService struct {
	cfg config.KVStore
}

func NewEmailService(cfg config.KVStore) *EmailService {
	return &EmailService{
		cfg: cfg,
	}
}

func (s *EmailService) SendReportEmail(
	ctx context.Context,
	request payload.SendReportEmailPayload,
) error {
	smtpHost := s.cfg.GetString("smtp.host")
	smtpPort := s.cfg.GetInt("smtp.port")
	smtpUser := s.cfg.GetString("smtp.username")
	smtpPass := s.cfg.GetString("smtp.password")
	senderEmail := s.cfg.GetString("smtp.sender_email")

	log.FromCtx(ctx).Info(fmt.Sprintf("SMTP Config Loaded: Host=%s, Port=%d, User=%s, PassLen=%d, Sender=%s",
		smtpHost, smtpPort, smtpUser, len(smtpPass), senderEmail))

	if smtpHost == "" || smtpPort == 0 || smtpUser == "" || smtpPass == "" || senderEmail == "" {
		log.FromCtx(ctx).Error(nil, "SMTP configuration is incomplete")
		return errors.WithStack(httpservice.ErrUnknownSource)
	}

	m := gomail.NewMessage()
	m.SetHeader("From", fmt.Sprintf("%s <%s>", request.SenderName, senderEmail))
	m.SetHeader("To", request.RecipientEmail)
	m.SetHeader("Subject", request.Subject)
	m.SetBody("text/html", request.BodyHTML)

	d := gomail.NewDialer(smtpHost, smtpPort, smtpUser, smtpPass)

	//d.SSL = true

	if err := d.DialAndSend(m); err != nil {
		log.FromCtx(ctx).Error(err, "failed to send email")
		return errors.WithStack(httpservice.ErrUnknownSource)
	}

	return nil
}
