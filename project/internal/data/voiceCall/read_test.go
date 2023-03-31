package voiceCall

import (
	"github.com/stretchr/testify/assert"
	"gitlab.skillbox.ru/denis_parovoi/profession-godev/finalWork/project/internal/models"
	"gitlab.skillbox.ru/denis_parovoi/profession-godev/finalWork/project/pkg/logging"
	"testing"
)

func TestRead(t *testing.T) {
	assert := assert.New(t)
	log := logging.GetLogger()

	expected := []models.VoiceCallData{{"BG", "40", "609", "E-Voice", float32(0.86), 160, 36, 5},
		{"DK", "11", "743", "JustPhone", float32(0.67), 82, 74, 41}}

	assert.Equal(expected, Read(log, "voice_test.data"))
}
