# postfix

Sending mail through postfix SMTP server

# example

```go
import (
	"github.com/vanng822/postfix"
)
fromAddr := postfix.NewAddress(from, fromEmail)
toAddr := postfix.NewAddress(to, toEmail)
msg, err := postfix.MultipartMessage(fromAddr, toAddr, subject, text, html)
if err != nil {
	log.Fatal("Postfix mailing with error", err)
}
err := postfix.Send(msg)
if err != nil {
	log.Fatal("Postfix mailing with error", err)
}
```
