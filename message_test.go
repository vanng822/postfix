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
	msg, err := MultipartMessage(from, to, subject, text, html)
	assert.Nil(t, err)
	assert.NotNil(t, msg)
	assert.Equal(t, from, msg.From())
	assert.True(t, msg.To().Contain(to))
	actual := msg.Bytes()
	assert.True(t, bytes.Contains(actual, []byte(subject)))
}
