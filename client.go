// Package postfix is for sending mail
//
// import (
//
//	"github.com/vanng822/postfix"
//
// )
//
// fromAddr := postfix.NewAddress(from, fromEmail)
// toAddr := postfix.NewAddress(to, toEmail)
// msg, err := postfix.MultipartMessage(fromAddr, toAddr, subject, text, html)
//
//	if err != nil {
//		log.Fatal("Postfix mailing with error", err)
//	}
//
// err := postfix.Send(msg)
//
//	if err != nil {
//		log.Fatal("Postfix mailing with error", err)
//	}
package postfix

import (
	"fmt"
	"net/smtp"
)

var (
	host          = "127.0.0.1"
	port          = 25
	clientFactory func() (c *smtp.Client, err error)
)

func init() {
	clientFactory = func() (c *smtp.Client, err error) {
		c, err = smtp.Dial(getAddr())
		return
	}
}

// SetClientFactory for setting a smtp client if needed
func SetClientFactory(f func() (c *smtp.Client, err error)) {
	clientFactory = f
}

func SetPort(p int) {
	port = p
}

func SetHost(h string) {
	host = h
}

func getAddr() string {
	return fmt.Sprintf("%s:%d", host, port)
}

func getClient() (c *smtp.Client, err error) {
	c, err = clientFactory()
	return
}

func Noop() error {
	c, err := getClient()
	if err != nil {
		return err
	}
	defer c.Close()
	return c.Noop()
}

func Hello() error {
	localName := "localhost"
	c, err := getClient()
	if err != nil {
		return err
	}
	defer c.Close()
	return c.Hello(localName)
}
