package server

import (
	"github.com/gorilla/mux"
	"gitlab.skillbox.ru/denis_parovoi/profession-godev/finalWork/project/api/handler"
	"gitlab.skillbox.ru/denis_parovoi/profession-godev/finalWork/project/internal/config"
	"gitlab.skillbox.ru/denis_parovoi/profession-godev/finalWork/project/pkg/logging"
	"net/http"
)

func Start(r *mux.Router, log *logging.Logger, cfg *config.Config) {
	cfgAddr := cfg.Listen.BindIP + ":" + cfg.Listen.Port
	srv := &http.Server{
		Handler: r,
		Addr:    cfgAddr,
	}
	r.HandleFunc("/api", handler.HandleConnection(log, cfg))
	log.Info().Msg("server is starting")
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal().Msgf("can't start server %s", err)
	}
}
