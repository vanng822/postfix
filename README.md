# postfix

Sending mail through postfix SMTP server

# example

```go
import (
	"github.com/vanng822/postfix"
)
fromAddr := postfix.NewAddress(from, fromEmail)
toAddr := postfix.NewAddress(to, toEmail)
attachment := &Attachment{
		Filename:    "test.txt",
		ContentType: "text/plain",
		Content:     []byte("test attachment content"),
	}

attachment2 := NewAttachment("test2.txt", "text/plain", []byte("test attachment content 2"))

msg, err := MultipartMessage(from, to, subject, text, html, attachment, attachment2)
if err != nil {
	log.Fatal("Postfix mailing with error", err)
}
err := postfix.Send(msg)
if err != nil {
	log.Fatal("Postfix mailing with error", err)
}
```
