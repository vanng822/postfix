package postfix

import (
	"net/smtp"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClientFactory(t *testing.T) {
	defer func(f func() (*smtp.Client, error)) {
		clientFactory = f
	}(clientFactory)

	testClient := &smtp.Client{}
	f := func() (*smtp.Client, error) {
		return testClient, nil
	}
	SetClientFactory(f)
	actual, err := getClient()
	assert.Nil(t, err)
	assert.Equal(t, actual, testClient)
}
