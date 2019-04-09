package postfix

import (
	"bytes"

	"github.com/vanng822/go-premailer/premailer"
	"github.com/vanng822/mail"
)

// NewAddress return new mail.Address
func NewAddress(name, address string) *mail.Address {
	return &mail.Address{Name: name, Address: address}
}

// MultipartMessage create a message with multipart text and html
// html part is parsed and applied css inline styling
func MultipartMessage(from *mail.Address, to *mail.Address, subject, text, html string) (*mail.Message, error) {
	var (
		msg         *mail.Message
		alternative *mail.Multipart
		prem        premailer.Premailer
		pcontent    string
		err         error
	)

	msg = mail.NewMessage()
	msg.SetFrom(from)
	msg.To().Add(to)
	msg.SetSubject(subject)

	alternative = mail.NewMultipart("multipart/alternative", msg)
	if err = alternative.AddText("text/plain", bytes.NewReader([]byte(text))); err != nil {
		return nil, err
	}

	prem, err = premailer.NewPremailerFromString(html, nil)
	if err != nil {
		return nil, err
	}
	pcontent, err = prem.Transform()
	if err != nil {
		return nil, err
	}
	if err = alternative.AddText("text/html", bytes.NewReader([]byte(pcontent))); err != nil {
		return nil, err
	}
	if err = alternative.Close(); err != nil {
		return nil, err
	}

	return msg, nil
}
