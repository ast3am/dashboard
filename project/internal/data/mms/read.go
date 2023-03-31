package mms

import (
	"encoding/json"
	"errors"
	"gitlab.skillbox.ru/denis_parovoi/profession-godev/finalWork/project/internal/models"
	"gitlab.skillbox.ru/denis_parovoi/profession-godev/finalWork/project/internal/provides"
	"gitlab.skillbox.ru/denis_parovoi/profession-godev/finalWork/project/pkg/countries"
	"gitlab.skillbox.ru/denis_parovoi/profession-godev/finalWork/project/pkg/logging"
	"io/ioutil"
	"net/http"
)

func Read(log *logging.Logger, url string) []models.MMSData {
	r, err := http.Get(url)
	result := []models.MMSData{}
	if err != nil {
		log.Err(err).Msg("")
		return result
	}
	if r.StatusCode == 500 {
		err = errors.New("no response from mms server")
		log.Err(err).Msg("")
		return result
	}

	countryList := countries.NewCountriesList()
	provideList := provides.NewSmsProvidesList()
	body, err := ioutil.ReadAll(r.Body)
	buffer := []models.MMSData{}
	err = json.Unmarshal(body, &buffer)
	if err != nil {
		log.Error().Msgf("mms data error: %s", err)
		return result
	}
	for _, val := range buffer {
		check := true
		if !provideList.CheckProvide(val.Provider) {
			log.Trace().Msgf("this line has incorrect provider\n %s\n", val.Provider)
			check = false
		} else if !countryList.CheckCountry(val.Country) {
			log.Trace().Msgf("this line has incorrect country\n %s\n", val.Country)
			check = false
		}
		if check {
			result = append(result, val)
		}
	}
	log.Debug().Msg("return mms data is OK")
	return result
}
