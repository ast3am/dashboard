package sms

import (
	"github.com/stretchr/testify/assert"
	"gitlab.skillbox.ru/denis_parovoi/profession-godev/finalWork/project/internal/models"
	"gitlab.skillbox.ru/denis_parovoi/profession-godev/finalWork/project/pkg/logging"
	"testing"
)

func TestRead(t *testing.T) {
	assert := assert.New(t)
	log := logging.GetLogger()

	expected := []models.SMSData{{"US", "36", "1576", "Rond"},
		{"BL", "68", "1594", "Kildy"}}

	assert.Equal(Read(log, "sms_test.data"), expected)
}
