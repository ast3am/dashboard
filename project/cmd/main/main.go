package main

import (
	"github.com/gorilla/mux"
	"gitlab.skillbox.ru/denis_parovoi/profession-godev/finalWork/project/api/server"
	"gitlab.skillbox.ru/denis_parovoi/profession-godev/finalWork/project/internal/config"
	"gitlab.skillbox.ru/denis_parovoi/profession-godev/finalWork/project/pkg/logging"
)

func main() {
	cfg := config.GetConfig("../../config.yml")
	log := logging.GetLogger()
	r := mux.NewRouter()
	server.Start(r, log, cfg)
}
