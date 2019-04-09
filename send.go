package postfix

import (
	"github.com/vanng822/mail"
)

// Send connects to a smtp server and send the message
// to all recipients
func Send(msg *mail.Message) error {
	addresses, err := msg.To().Addresses()
	if err != nil {
		return err
	}

	c, err := GetClient()
	if err != nil {
		return err
	}
	defer c.Close()

	if err = c.Mail(msg.From().String()); err != nil {
		return err
	}
	for _, adress := range addresses {
		if err = c.Rcpt(adress.String()); err != nil {
			return err
		}
	}

	w, err := c.Data()
	if err != nil {
		return err
	}
	_, err = w.Write(msg.Bytes())
	if err != nil {
		return err
	}
	err = w.Close()
	if err != nil {
		return err
	}
	return c.Quit()
}
