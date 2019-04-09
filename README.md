# postfix

Sending mail through postfix SMTP server

# example

```go
import (
	"github.com/vanng822/gamlich/core/postfix"
)
fromAddr := postfix.NewAddress(from, fromEmail)
toAddr := postfix.NewAddress(to, toEmail)
msg, err := postfix.MultipartMessage(fromAddr, toAddr, subject, text, html)
if err != nil {
	return err
}
err := postfix.Send(msg)
if err != nil {
	log.Println("Postfix mailing with error", err)
}
```
