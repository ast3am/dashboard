package support

import (
	"encoding/json"
	"errors"
	"gitlab.skillbox.ru/denis_parovoi/profession-godev/finalWork/project/internal/models"
	"gitlab.skillbox.ru/denis_parovoi/profession-godev/finalWork/project/pkg/logging"
	"io/ioutil"
	"net/http"
)

func Read(log *logging.Logger, url string) []models.SupportData {
	result := []models.SupportData{}
	r, err := http.Get(url)
	if err != nil {
		log.Err(err).Msg("")
		return result
	}
	if r.StatusCode == 500 {
		err = errors.New("no response from support server")
		log.Err(err).Msg("")
		return result
	}

	body, err := ioutil.ReadAll(r.Body)

	err = json.Unmarshal(body, &result)
	if err != nil {
		log.Error().Msgf("Support data error: %s", err)
		return result
	}

	log.Debug().Msg("return support data is OK")
	return result
}
