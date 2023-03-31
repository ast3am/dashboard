package service

import (
	"gitlab.skillbox.ru/denis_parovoi/profession-godev/finalWork/project/internal/config"
	"gitlab.skillbox.ru/denis_parovoi/profession-godev/finalWork/project/internal/data/billing"
	"gitlab.skillbox.ru/denis_parovoi/profession-godev/finalWork/project/internal/data/incidents"
	"gitlab.skillbox.ru/denis_parovoi/profession-godev/finalWork/project/internal/data/mail"
	"gitlab.skillbox.ru/denis_parovoi/profession-godev/finalWork/project/internal/data/mms"
	"gitlab.skillbox.ru/denis_parovoi/profession-godev/finalWork/project/internal/data/sms"
	"gitlab.skillbox.ru/denis_parovoi/profession-godev/finalWork/project/internal/data/support"
	"gitlab.skillbox.ru/denis_parovoi/profession-godev/finalWork/project/internal/data/voiceCall"
	"gitlab.skillbox.ru/denis_parovoi/profession-godev/finalWork/project/internal/models"
	"gitlab.skillbox.ru/denis_parovoi/profession-godev/finalWork/project/pkg/countries"
	"gitlab.skillbox.ru/denis_parovoi/profession-godev/finalWork/project/pkg/logging"
	"sort"
	"sync"
)

func GetResultData(log *logging.Logger, cfg *config.Config) models.ResultSetT {
	result := models.ResultSetT{}
	countryList := *countries.NewCountriesList()

	var wg sync.WaitGroup
	wg.Add(7)

	//sms
	go func() {
		smsSortByCountry := sms.Read(log, cfg.Data.SMS)
		for i := range smsSortByCountry {
			smsSortByCountry[i].Country = countryList[smsSortByCountry[i].Country]
		}
		smsSortByProvider := make([]models.SMSData, len(smsSortByCountry))
		copy(smsSortByProvider, smsSortByCountry)
		sort.Slice(smsSortByCountry, func(i, j int) bool {
			return smsSortByCountry[i].Country < smsSortByCountry[j].Country
		})
		sort.Slice(smsSortByProvider, func(i, j int) bool {
			return smsSortByProvider[i].Provider < smsSortByProvider[j].Provider
		})
		result.SMS = [][]models.SMSData{smsSortByProvider, smsSortByCountry}
		wg.Done()
	}()

	//mms
	go func() {
		mmsSortByCountry := mms.Read(log, cfg.Data.MMS.URL)
		for i := range mmsSortByCountry {
			mmsSortByCountry[i].Country = countryList[mmsSortByCountry[i].Country]
		}
		mmsSortByProvider := make([]models.MMSData, len(mmsSortByCountry))
		copy(mmsSortByProvider, mmsSortByCountry)
		sort.Slice(mmsSortByCountry, func(i, j int) bool {
			return mmsSortByCountry[i].Country < mmsSortByCountry[j].Country
		})
		sort.Slice(mmsSortByProvider, func(i, j int) bool {
			return mmsSortByProvider[i].Provider < mmsSortByProvider[j].Provider
		})
		result.MMS = [][]models.MMSData{mmsSortByProvider, mmsSortByCountry}
		wg.Done()
	}()

	// voiceCall
	go func() {
		result.VoiceCall = voiceCall.Read(log, cfg.Data.VoiceCall)
		wg.Done()
	}()

	//email
	go func() {
		mailData := emailData.Read(log, cfg.Data.Email)
		countrySortMap := make(map[string][]models.EmailData)
		for _, val := range mailData {
			countrySortMap[val.Country] = append(countrySortMap[val.Country], val)
		}
		resultMap := make(map[string][][]models.EmailData)
		for key, val := range countrySortMap {
			sort.Slice(val, func(i, j int) bool {
				return val[i].DeliveryTime < val[j].DeliveryTime
			})
			if len(val) >= 3 {
				fast := val[:3]
				slow := val[len(val)-3:]
				resultMap[key] = [][]models.EmailData{fast, slow}
			} else {
				fast := val
				slow := val
				resultMap[key] = [][]models.EmailData{fast, slow}
			}
		}
		result.Email = resultMap
		wg.Done()
	}()

	//billing
	go func() {
		result.Billing = billing.Read(log, cfg.Data.Billing)
		wg.Done()
	}()

	// support
	go func() {
		supportData := support.Read(log, cfg.Data.Support.URL)
		supportResult := make([]int, 2, 2)
		var count int
		if len(supportData) == 0 {
			supportResult[0] = -1
		} else {
			for _, val := range supportData {
				count += val.ActiveTickets
			}
			switch {
			case count < 9:
				supportResult[0] = 1
			case count <= 16:
				supportResult[0] = 2
			default:
				supportResult[0] = 3
			}
		}
		supportResult[1] = count*60/18 + 1
		result.Support = supportResult
		wg.Done()
	}()

	//incidents
	go func() {
		incidentsData := incidents.Read(log, cfg.Data.Incidents.URL)
		sort.Slice(incidentsData, func(i, j int) bool {
			return incidentsData[i].Status < incidentsData[j].Status
		})
		result.Incidents = incidentsData
		wg.Done()

	}()

	wg.Wait()
	return result
}
