package support

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"gitlab.skillbox.ru/denis_parovoi/profession-godev/finalWork/project/internal/models"
	"gitlab.skillbox.ru/denis_parovoi/profession-godev/finalWork/project/pkg/logging"
	"net/http"
	"net/http/httptest"
	"testing"
)

func main() {

}

func TestReadPositive(t *testing.T) {
	byt := `[
  {
	"topic": "SMS",
	"active_tickets": 3
	},
	{
	"topic": "MMS",
	"active_tickets": 9
	},
	{
	"topic": "Billing",
	"active_tickets": 0
	}
]
`
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, byt)
	}))
	defer srv.Close()
	assert := assert.New(t)
	log := logging.GetLogger()

	expected := []models.SupportData{
		{"SMS", 3},
		{"MMS", 9},
		{"Billing", 0},
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

	expected := []models.SupportData{}

	assert.Equal(expected, Read(log, srv.URL))
}
