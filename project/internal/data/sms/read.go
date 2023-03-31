package sms

import (
	"bufio"
	"encoding/csv"
	"gitlab.skillbox.ru/denis_parovoi/profession-godev/finalWork/project/internal/models"
	"gitlab.skillbox.ru/denis_parovoi/profession-godev/finalWork/project/internal/provides"
	"gitlab.skillbox.ru/denis_parovoi/profession-godev/finalWork/project/pkg/countries"
	"gitlab.skillbox.ru/denis_parovoi/profession-godev/finalWork/project/pkg/logging"
	"io"
	"os"
)

func Read(log *logging.Logger, data string) []models.SMSData {
	countryList := countries.NewCountriesList()
	provideList := provides.NewSmsProvidesList()
	resultData := make([]models.SMSData, 0, 0)
	smsFile, err := os.Open(data)
	if err != nil {
		log.Error().Msgf("can't read smsFile \n %s", err)
		return resultData
	}
	defer smsFile.Close()
	reared := csv.NewReader(bufio.NewReader(smsFile))
	reared.Comma = ';'
	for {
		line, err := reared.Read()
		if err == io.EOF {
			break
		} else if len(line) != 4 {
			log.Trace().Msgf("this line doesn't have 4 items\n %s\n", line)
		} else if !provideList.CheckProvide(line[3]) {
			log.Trace().Msgf("this line has incorrect provider\n %s\n", line)
		} else if !countryList.CheckCountry(line[0]) {
			log.Trace().Msgf("this line has incorrect country\n %s\n", line)
		} else {
			resultData = append(resultData, models.SMSData{
				Country:      line[0],
				Bandwidth:    line[1],
				ResponseTime: line[2],
				Provider:     line[3],
			})
		}
	}
	log.Debug().Msg("return mms data is OK")
	return resultData
}
