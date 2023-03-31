package billing

import (
	"gitlab.skillbox.ru/denis_parovoi/profession-godev/finalWork/project/internal/models"
	"gitlab.skillbox.ru/denis_parovoi/profession-godev/finalWork/project/pkg/logging"
	"io/ioutil"
	"math"
)

type billingSet byte

const (
	CreateCustomer billingSet = 1 << iota
	Purchase
	Payout
	Recurring
	FraudControl
	CheckoutPage
)

func (k billingSet) Unmarshal(log *logging.Logger) *models.BillingData {
	result := models.BillingData{false, false, false, false, false, false}
	if k >= CheckoutPage<<1 {
		log.Info().Msgf("<unknown key: %d>", k)
	}
	for key := CreateCustomer; key <= CheckoutPage; key <<= 1 {
		if k&key != 0 {
			switch key {
			case CreateCustomer:
				result.CreateCustomer = true
			case Purchase:
				result.Purchase = true
			case Payout:
				result.Payout = true
			case Recurring:
				result.Recurring = true
			case FraudControl:
				result.FraudControl = true
			case CheckoutPage:
				result.CheckoutPage = true
			}
		}
	}
	log.Debug().Msg("return billing data is OK")
	return &result
}

func Read(log *logging.Logger, data string) *models.BillingData {
	content, err := ioutil.ReadFile(data)
	if err != nil {
		log.Error().Msgf("can't read billing File \n %s", err)
		return nil
	}
	var byteTemp float64
	for i := 0; i < len(content); i++ {
		if content[len(content)-(i+1)] == 49 {
			byteTemp += math.Pow(2.0, float64(i))
		}
	}
	byt := billingSet(byteTemp)
	return byt.Unmarshal(log)
}
