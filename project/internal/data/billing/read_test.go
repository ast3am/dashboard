package billing

import (
	"github.com/stretchr/testify/assert"
	"gitlab.skillbox.ru/denis_parovoi/profession-godev/finalWork/project/internal/models"
	"gitlab.skillbox.ru/denis_parovoi/profession-godev/finalWork/project/pkg/logging"
	"testing"
)

func TestBillingSet_Unmarshal(t *testing.T) {
	assert := assert.New(t)
	log := logging.GetLogger()

	var tests = []struct {
		data     byte
		expected *models.BillingData
	}{
		{0b00010011, &models.BillingData{true, true, false, false, true, false}},
		{0b00000011, &models.BillingData{true, true, false, false, false, false}},
		{0b00110011, &models.BillingData{true, true, false, false, true, true}},
		{0b01011010, &models.BillingData{false, true, false, true, true, false}},
	}

	for _, test := range tests {
		assert.Equal(test.expected, billingSet(test.data).Unmarshal(log))
	}
}
