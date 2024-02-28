package mail

import (
	"bytes"
	"context"
	"crypto/md5"
	"encoding/hex"
	"errors"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/google/uuid"
	"github.com/mailgun/mailgun-go/v4"
)

type MailgunClient struct {
	*mailgun.MailgunImpl
}

// NewMailgunClient creates new Mailgun client given config
func NewMailgun(config Config) ClientImpl {
	return &MailgunClient{
		mailgun.NewMailgun(config.Domain, config.APIKey),
	}
}

func (mg *MailgunClient) CreateNewMessage(from, subject, text, to, htmlStr string, files *Files) Message {
	message := mg.NewMessage(
		from,
		subject,
		text,
		to,
	)
	message.SetHtml(htmlStr)
	if files.Attaches != nil {
		for i := 0; i < len(files.Attaches); i++ {
			attach := files.Attaches[i]
			if attach.Buffer != nil {
				message.AddBufferAttachment(attach.DisplayFilename, attach.Buffer)
			} else if attach.File != "" {
				message.AddAttachment(attach.File)
			}
		}
	}

	if files.Embeds != nil {
		for i := 0; i < len(files.Attaches); i++ {
			embed := files.Embeds[i]
			if embed.Buffer != nil {
				message.AddReaderInline(embed.DisplayFilename, fromBytes(embed.Buffer))
			} else if embed.readCloser != nil {
				message.AddReaderInline(embed.DisplayFilename, embed.readCloser)
			} else if embed.File != "" {
				message.AddInline(embed.File)
			}
		}
	}
	return message
}

func (mg *MailgunClient) SendWithContext(ctx context.Context, message Message) error {
	if mgMessage, ok := message.(*mailgun.Message); ok {
		_, _, err := mg.Send(ctx, mgMessage)
		if err != nil {
			return err
		}
		return nil
	}
	return errors.New("incompatible message type passed to SendWithContext")
}

func fromBytes(data []byte) io.ReadCloser {

	reader := bytes.NewReader(data)
	closer := io.NopCloser(reader)

	// 不需要执行 Closer mailgun内部getPayloadBuffer()会执行defer file.value.Close()
	// defer closer.Close()

	return closer
}

func FromUrl(url string) io.ReadCloser {
	// resp, err := http.Get(url)
	// if err != nil {
	// 	return io.NopCloser(resp)
	// }
	// defer resp.Body.Close()
	// return resp.Body
	// 生成随机文件名
	uuid := uuid.New()
	tmpFile := filepath.Join(os.TempDir(), uuid.String()+".txt")

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	// defer resp.Body.Close()

	// 创建一个新的文件
	file, err := os.Create("/tmp/")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	// 将响应体复制到文件
	_, err = io.Copy(file, resp.Body)
	if err != nil {
		panic(err)
	}

	// 计算文件指纹
	fileBytes, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}

	md5 := md5.New()
	fingerprint := hex.EncodeToString(md5.Sum(fileBytes))

	// 重命名文件
	dstFile := filepath.Join(os.TempDir(), fingerprint+".txt")
	err = os.Rename(tmpFile, dstFile)
	if err != nil {
		panic(err)
	}

	return resp.Body

}
