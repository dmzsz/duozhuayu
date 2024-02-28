package mail

import (
	"bytes"
	"context"
	"errors"

	"gopkg.in/mail.v2"
)

type GoMailer struct {
	config Config
	client *mail.Dialer
}

func NewGoMailer(config Config) ClientImpl {
	return &GoMailer{
		config: config,
		client: mail.NewDialer("smtp.gmail.com", 587, config.FromEmail, config.Password),
	}
}

// CreateNewMessage implements ClientImpl.
func (mailer *GoMailer) CreateNewMessage(from string, subject string, text string, to string, htmlStr string, files *Files) Message {
	message := mail.NewMessage()
	message.SetHeader("From", from)
	message.SetHeader("To", to)
	message.SetHeader("Subject", subject)
	message.SetBody("text/html", htmlStr)
	// message.SetBody("text/plain", text)

	if files.Attaches != nil {
		for i := 0; i < len(files.Attaches); i++ {
			attach := files.Attaches[i]
			if attach.Buffer != nil {
				message.AttachReader(attach.File, bytes.NewReader(attach.Buffer), mail.Rename(attach.DisplayFilename))
			} else if attach.File != "" {
				// message.Attach("/tmp/0000146.jpg", mail.Rename("picture.jpg")) 用户下载文件名字是picture.jpg
				message.Attach(attach.File, mail.Rename(attach.DisplayFilename))
			} else {
				message.Attach(attach.File)
			}
		}
	}

	if files.Embeds != nil {
		for i := 0; i < len(files.Attaches); i++ {
			embed := files.Embeds[i]
			if embed.Buffer != nil {
				message.EmbedReader(embed.File, bytes.NewReader(embed.Buffer), mail.Rename(embed.DisplayFilename))
			} else if embed.File != "" {
				// message.Embed("/tmp/0000146.jpg", mail.Rename("picture.jpg"))
				// message.SetBody("text/html", `<img src="cid:picture.jpg" alt="My image" />`)
				// data:[<mediatype>][;base64],<data> body嵌入
				message.Embed(embed.File, mail.Rename(embed.DisplayFilename))
			} else {
				message.Embed(embed.File)
			}
		}
	}

	return message
}

// SendWithContext implements ClientImpl.
func (mailer *GoMailer) SendWithContext(ctx context.Context, message Message) error {
	if mgMessage, ok := message.(*mail.Message); ok {
		err := mailer.client.DialAndSend(mgMessage)
		if err != nil {
			return err
		}
		return nil
	}
	return errors.New("incompatible message type passed to SendWithContext")

}
