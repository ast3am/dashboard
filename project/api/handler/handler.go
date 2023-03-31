package handler

import (
	"encoding/json"
	"gitlab.skillbox.ru/denis_parovoi/profession-godev/finalWork/project/internal/config"
	"gitlab.skillbox.ru/denis_parovoi/profession-godev/finalWork/project/internal/models"
	"gitlab.skillbox.ru/denis_parovoi/profession-godev/finalWork/project/internal/service"
	"gitlab.skillbox.ru/denis_parovoi/profession-godev/finalWork/project/pkg/logging"
	"net/http"
	"time"
)

var cacheResult []byte
var cacheTime time.Time

func HandleConnection(log *logging.Logger, cfg *config.Config) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		nowTime := time.Now()
		if cacheTime.After(nowTime) {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
			w.Write(cacheResult)
			log.Info().Msg("request processed, from cache")
			return
		}

		resultT := models.ResultT{}
		data := service.GetResultData(log, cfg)
		if checkData(&data) {
			resultT.Status = true
			resultT.Data = data
			log.Debug().Msg("GetResult true")
		} else {
			resultT.Error = "Error on collect data"
			log.Debug().Msg("GetResult error")
		}

		answer, _ := json.Marshal(resultT)
		cacheResult = answer
		cacheTime = nowTime.Add((time.Duration(cfg.CacheTime)) * time.Second)
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Write(answer)
		log.Info().Msg("request processed, cache has been updated")
	}
}

func checkData(t *models.ResultSetT) bool {
	switch {
	case len(t.SMS[0]) == 0:
		return false
	case len(t.MMS[0]) == 0:
		return false
	case len(t.VoiceCall) == 0:
		return false
	case len(t.Email) == 0:
		return false
	case t.Support[1] == -1:
		return false
	case len(t.Incidents) == 0:
		return false
	case t.Billing == nil:
		return false
	}
	return true
}
