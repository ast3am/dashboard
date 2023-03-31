package voiceCall

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

func Read(log *logging.Logger, data string) []models.VoiceCallData {
	countryList := countries.NewCountriesList()
	provideList := provides.NewVoiceProvidesList()
	resultData := make([]models.VoiceCallData, 0, 0)
	voiceData, err := os.Open(data)
	if err != nil {
		log.Error().Msgf("can't read voiceCall \n %s", err)
		return resultData
	}
	defer voiceData.Close()
	reared := csv.NewReader(bufio.NewReader(voiceData))
	reared.Comma = ';'
	for {
		line, err := reared.Read()
		if err == io.EOF {
			break
		} else if len(line) != 8 {
			log.Trace().Msgf("this line doesn't have 8 items\n %s\n", line)
		} else if !provideList.CheckProvide(line[3]) {
			log.Trace().Msgf("this line has incorrect provider\n %s\n", line)
		} else if !countryList.CheckCountry(line[0]) {
			log.Trace().Msgf("this line has incorrect country\n %s\n", line)
		} else {
			connectionStability, err := strconv.ParseFloat(line[4], 32)
			if err != nil {
				log.Trace().Msgf("can't parse connectionStability from this data: %s", line[4])
				continue
			}
			tTFB, err := strconv.Atoi(line[5])
			if err != nil {
				log.Trace().Msgf("can't parse TTFB from this data: %s", line[5])
				continue
			}
			voicePurity, err := strconv.Atoi(line[6])
			if err != nil {
				log.Trace().Msgf("can't parse voicePurity from this data: %s", line[6])
				continue
			}
			medianTime, err := strconv.Atoi(line[7])
			if err != nil {
				log.Trace().Msgf("can't parse medianTime from this data: %s", line[7])
				continue
			}
			resultData = append(resultData, models.VoiceCallData{
				Country:             line[0],
				Bandwidth:           line[1],
				ResponseTime:        line[2],
				Provider:            line[3],
				ConnectionStability: float32(connectionStability),
				TTFB:                tTFB,
				VoicePurity:         voicePurity,
				MedianOfCallsTime:   medianTime,
			})
		}
	}
	log.Debug().Msg("return voiceCall data is OK")
	return resultData
}
