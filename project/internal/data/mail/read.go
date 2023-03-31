package emailData

import (
	"bufio"
	"encoding/csv"
	"gitlab.skillbox.ru/denis_parovoi/profession-godev/finalWork/project/internal/models"
	"gitlab.skillbox.ru/denis_parovoi/profession-godev/finalWork/project/internal/provides"
	"gitlab.skillbox.ru/denis_parovoi/profession-godev/finalWork/project/pkg/countries"
	"gitlab.skillbox.ru/denis_parovoi/profession-godev/finalWork/project/pkg/logging"
	"io"
	"os"
	"strconv"
)

func Read(log *logging.Logger, data string) []models.EmailData {
	countryList := countries.NewCountriesList()
	provideList := provides.NewMailProvidesList()
	resultData := make([]models.EmailData, 0, 0)
	voiceData, err := os.Open(data)
	if err != nil {
		log.Error().Msgf("can't read e-mail File \n %s", err)
		return resultData
	}
	defer voiceData.Close()
	reared := csv.NewReader(bufio.NewReader(voiceData))
	reared.Comma = ';'
	for {
		line, err := reared.Read()
		if err == io.EOF {
			break
		} else if len(line) != 3 {
			log.Trace().Msgf("this line doesn't have 3 items\n %s\n", line)
		} else if !provideList.CheckProvide(line[1]) {
			log.Trace().Msgf("this line has incorrect provider\n %s\n", line)
		} else if !countryList.CheckCountry(line[0]) {
			log.Trace().Msgf("this line has incorrect country\n %s\n", line)
		} else {
			deliveryTime, err := strconv.Atoi(line[2])
			if err != nil {
				log.Trace().Msgf("can't parse deliveryTime from this data: %s", line[2])
				continue
			}
			resultData = append(resultData, models.EmailData{
				Country:      line[0],
				Provider:     line[1],
				DeliveryTime: deliveryTime,
			})
		}
	}
	log.Debug().Msg("return mail data is OK")
	return resultData
}
