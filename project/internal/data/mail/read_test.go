package emailData

import (
	"github.com/stretchr/testify/assert"
	"gitlab.skillbox.ru/denis_parovoi/profession-godev/finalWork/project/internal/models"
	"gitlab.skillbox.ru/denis_parovoi/profession-godev/finalWork/project/pkg/logging"
	"testing"
)

func TestRead(t *testing.T) {
	assert := assert.New(t)
	log := logging.GetLogger()

	expected := []models.EmailData{{"AT", "Hotmail", 487}}
	assert.Equal(expected, Read(log, "mail_test.data"))
}
