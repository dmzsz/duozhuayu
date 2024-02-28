package v1

import (
	"encoding/hex"
	"strconv"
	"time"

	"github.com/dmzsz/duozhuayu/internal/configs"
	"github.com/dmzsz/duozhuayu/internal/constants"
	"github.com/dmzsz/duozhuayu/internal/datasources/caches"
	"github.com/dmzsz/duozhuayu/pkg/helpers"
	"github.com/dmzsz/duozhuayu/pkg/mail"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
)

// EmailService interface
type EmailService interface {
	Welcome(toEmail string) error
	ResetPassword(toEmail, token string) error
	VerifyNewUser(toEmail, token string) error
	VerifyUpdatedEmail(toEmail, token string) error
}

type emailService struct {
	client mail.MailImpl
}

// NewEmailService instantiates a Email Service
func NewEmailService(client mail.MailImpl) EmailService {
	return &emailService{
		client: client,
	}
}

func (es *emailService) Welcome(toEmail string) error {
	return es.client.SendUrl(welcomeSubject, welcomeText, toEmail, welcomeHTML)

}

func (es *emailService) ResetPassword(toEmail, token string) error {
	data := struct {
		key   string
		value string
	}{}

	if configs.IsPasswordRecoverCodeUUIDv4() {
		data.key = constants.PasswordResetKeyPrefix + uuid.NewString()
	} else {
		code := helpers.SecureRandomNumber(configs.AppConfig.EmailConfig.PasswordResetCodeLength)
		data.key = constants.PasswordResetKeyPrefix + strconv.FormatUint(code, 10)
	}

	data.value = toEmail

	// when encryption at rest is used
	if configs.IsCipher() {
		var err error

		// hash of the email in hexadecimal string format
		value, err := helpers.CalcHash(
			toEmail,
			configs.AppConfig.SecurityConfig.Blake2bSec,
		)
		if err != nil {
			log.WithError(err).Error("error code: 406.1")
			return err
		}
		data.value = hex.EncodeToString(value)
	}

	caches.GetRedisCache().Set(
		data.key,
		data.value,
		time.Duration(configs.AppConfig.EmailConfig.PasswordResetValidityPeriod)*time.Second)

	return es.client.ResetPassword(resetSubject, resetTextTmpl, toEmail, resetHTMLTmpl, token)
}

func (es *emailService) VerifyNewUser(toEmail, token string) error {
	data := struct {
		key   string
		value string
	}{}

	if configs.IsEmailVerificationCodeUUIDv4() {
		data.key += constants.EmailVerificationKeyPrefix + uuid.NewString()
	} else {
		code := helpers.SecureRandomNumber(configs.AppConfig.EmailConfig.EmailVerificationCodeLength)
		data.key += strconv.FormatUint(code, 10)
	}

	keyTTL := configs.AppConfig.EmailConfig.EmailVerifyValidityPeriod
	if keyTTL == 0 {
		keyTTL = 15 * 60
	}

	data.value = toEmail

	// when encryption at rest is used
	if configs.IsCipher() {
		var err error

		// hash of the email in hexadecimal string format
		value, err := helpers.CalcHash(
			toEmail,
			configs.AppConfig.SecurityConfig.Blake2bSec,
		)
		if err != nil {
			log.WithError(err).Error("error code: 406.1")
			return err
		}
		data.value = hex.EncodeToString(value)
	}

	caches.GetRedisCache().Set(data.key, data.value, time.Duration(keyTTL)*time.Second)

	return es.client.VerifyNewUser(newSubject, verifyTextTmpl, toEmail, verifyHTMLTmpl, token)
}

func (es *emailService) VerifyUpdatedEmail(toEmail, token string) error {

	data := struct {
		key   string
		value string
	}{}

	if configs.IsEmailVerificationCodeUUIDv4() {
		data.key += constants.EmailVerificationKeyPrefix + uuid.NewString()
	} else {
		code := helpers.SecureRandomNumber(configs.AppConfig.EmailConfig.EmailVerificationCodeLength)
		data.key += strconv.FormatUint(code, 10)
	}

	keyTTL := configs.AppConfig.EmailConfig.EmailVerifyValidityPeriod
	if keyTTL == 0 {
		keyTTL = 15 * 60
	}

	data.value = toEmail

	// when encryption at rest is used
	if configs.IsCipher() {
		var err error

		// hash of the email in hexadecimal string format
		value, err := helpers.CalcHash(
			toEmail,
			configs.AppConfig.SecurityConfig.Blake2bSec,
		)
		if err != nil {
			log.WithError(err).Error("error code: 406.1")
			return err
		}
		data.value = hex.EncodeToString(value)
	}

	caches.GetRedisCache().Set(data.key, data.value, time.Duration(keyTTL)*time.Second)

	return es.client.VerifyUpdatedEmail(resetSubject, verifyTextTmpl, toEmail, verifyHTMLTmpl, token)
}
