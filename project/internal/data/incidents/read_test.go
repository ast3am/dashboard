package incidents

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"gitlab.skillbox.ru/denis_parovoi/profession-godev/finalWork/project/internal/models"
	"gitlab.skillbox.ru/denis_parovoi/profession-godev/finalWork/project/pkg/logging"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestReadPositive(t *testing.T) {
	byt := `[
  {
	"topic": "Billing isn’t allowed in US",
	"status": "closed"
	},
	{
	"topic": "Wrong SMS delivery time",
	"status": "active"
	},
	{
	"topic": "Support overloaded because of EU affect",
	"status": "active"
	}
]
`
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, byt)
	}))
	defer srv.Close()
	assert := assert.New(t)
	log := logging.GetLogger()

	expected := []models.IncidentData{
		{"Billing isn’t allowed in US", "closed"},
		{"Wrong SMS delivery time", "active"},
		{"Support overloaded because of EU affect", "active"},
	}

	assert.Equal(expected, Read(log, srv.URL))
}

func TestReadNegative(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer srv.Close()
	assert := assert.New(t)
	log := logging.GetLogger()

	expected := []models.IncidentData{}

	assert.Equal(expected, Read(log, srv.URL))
}
