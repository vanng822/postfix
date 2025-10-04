package postfix

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMultipartMessage(t *testing.T) {
	from := NewAddress("from", "from@example.com")
	to := NewAddress("to", "to@example.com")
	subject := "test subject"
	text := "test text"
	html := "<b>Click her</b>"
	msg, err := MultipartMessage(from, to, subject, text, html, nil)
	assert.Nil(t, err)
	assert.NotNil(t, msg)
	assert.Equal(t, from, msg.From())
	assert.True(t, msg.To().Contain(to))
	actual := msg.Bytes()
	assert.True(t, bytes.Contains(actual, []byte(subject)))
}

func TestMultipartMessageWithAttachment(t *testing.T) {
	from := NewAddress("from", "from@example.com")
	to := NewAddress("to", "to@example.com")
	subject := "test subject"
	text := "test text"
	html := "<b>Click her</b>"
	attachment := Attachment{
		Filename:    "test.txt",
		ContentType: "text/plain",
		Content:     []byte("test attachment content"),
	}

	attachment2 := Attachment{
		Filename:    "test2.txt",
		ContentType: "text/plain",
		Content:     []byte("test attachment content 2"),
	}
	msg, err := MultipartMessage(from, to, subject, text, html, []Attachment{attachment, attachment2})
	assert.Nil(t, err)
	assert.NotNil(t, msg)
	assert.Equal(t, from, msg.From())
	assert.True(t, msg.To().Contain(to))
	actual := msg.Bytes()
	assert.True(t, bytes.Contains(actual, []byte(subject)))
	assert.True(t, bytes.Contains(actual, []byte("Content-Location: test.txt")))
	assert.True(t, bytes.Contains(actual, []byte("Content-Location: test2.txt")))
}
