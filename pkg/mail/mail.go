package mail

import (
	"context"
	"fmt"
	"io"
	"time"

	"github.com/dmzsz/duozhuayu/internal/configs"
)

var GomailProvider string = "gamil"
var MailgunProvider string = "mailgun"

type MailImpl interface {
	Send(subject, text, to, htmlStr string, files *Files) (err error)
	SendOTP(subject, text, to, htmlStr, token string) (err error)
	SendUrl(subject, text, to, htmlStr, callbackUrl string) (err error)
}

type Message interface{}
type Files struct {
	Attaches []Attach
	Embeds   []Embed // 最好减少使用嵌入文件 采用Attaches附件方式比较好。嵌入式的体积较大，但是需要图文混编的话的确可以用这个减少静态服务器的压力，相比用静态服务器中的url包含在body中的写法。
}
type Attach struct {
	File            string // 服务器路径下的文件 /var/www/static/file.txt
	DisplayFilename string // 用户下载文件时看到的文件名 留空的话用源文件名
	Buffer          []byte // 使用byte数组传递文件内如数据。用于mailgun
}
type Embed struct {
	File            string // 需要上传的文件 例如磁盘中有这么一个文件 /var/www/static/file.txt
	DisplayFilename string // 用户下载文件时看到的文件名 留空的话用源文件名
	Buffer          []byte // 需要上传的文件 转换为的[]byte
	readCloser      io.ReadCloser
}
type ClientImpl interface {
	CreateNewMessage(from, subject, text, to, htmlStr string, files *Files) Message
	SendWithContext(ctx context.Context, message Message) error
}

type Config configs.EmailConfig

type Mail struct {
	conf   Config
	client ClientImpl
}

// NewMailgunClient creates new Mailgun client given config
func NewMail() MailImpl {
	config := Config(configs.AppConfig.EmailConfig)
	if config.Provider == MailgunProvider {
		return &Mail{
			conf: Config{
				Host:      config.Host,
				Port:      config.Port,
				FromEmail: config.FromEmail,
			},
			client: NewMailgun(config),
		}
	} else if config.Provider == GomailProvider {
		return &Mail{
			conf: Config{
				Host:      config.Host,
				Port:      config.Port,
				FromEmail: config.FromEmail,
			},
			client: NewGoMailer(config),
		}
	}
	return nil
}

// 普通邮件 附件可选
func (mg *Mail) Send(subject, text, to, htmlStr string, files *Files) (err error) {

	message := mg.client.CreateNewMessage(
		mg.conf.FromEmail,
		subject,
		text,
		to,
		htmlStr,
		files,
	)

	ctx, cancel := mg.setContext(10)
	defer cancel()
	return mg.client.SendWithContext(ctx, message)
}

// SendOTP implements MailImpl.
func (mg *Mail) SendOTP(subject, text, to, htmlStr, token string) (err error) {
	resetText := fmt.Sprintf(text, token)
	resetHTML := fmt.Sprintf(htmlStr, token)
	message := mg.client.CreateNewMessage(
		mg.conf.FromEmail,
		subject,
		resetText,
		to,
		resetHTML,
		nil,
	)

	ctx, cancel := mg.setContext(10)
	defer cancel()
	return mg.client.SendWithContext(ctx, message)
}

func (mg *Mail) SendUrl(subject, text, to, htmlStr, callbackUrl string) (err error) {
	// v := url.Values{}
	// v.Set("token", token)
	// v.Encode()

	resetURL := mg.getURL() + callbackUrl
	resetText := fmt.Sprintf(text, resetURL)
	resetHTML := fmt.Sprintf(htmlStr, resetURL)

	message := mg.client.CreateNewMessage(
		mg.conf.FromEmail,
		subject,
		resetText,
		to,
		resetHTML,
		nil,
	)

	ctx, cancel := mg.setContext(10)
	defer cancel()
	return mg.client.SendWithContext(ctx, message)
}

// ========= Private methods =========

func (mg *Mail) getURL() string {
	url := mg.conf.Host + ":" + mg.conf.Port
	return url
}

func (mg *Mail) setContext(seconds time.Duration) (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), time.Second*seconds)
}
