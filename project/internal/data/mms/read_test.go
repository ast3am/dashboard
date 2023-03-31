package mms

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
    "country": "U5",
    "provider": "Topol",
    "bandwidth": "41",
    "response_time": "910"
  },
  {
    "country": "US",
    "provider": "Rond",
    "bandwidth": "36",
    "response_time": "1576"
  },
  {
    "country": "GB",
    "provider": "kilty",
    "bandwidth": "28",
    "response_time": "495"
  },
  {
    "country": "F2",
    "provider": "Topolo",
    "bandwidth": "9",
    "response_time": "484"
  },
  {
    "country": "BL",
    "provider": "Kildy",
    "bandwidth": "68",
    "response_time": "1594"
  }
]
`
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, byt)
	}))
	defer srv.Close()
	assert := assert.New(t)
	log := logging.GetLogger()

	expected := []models.MMSData{{"US", "Rond", "36", "1576"},
		{"BL", "Kildy", "68", "1594"}}

	assert.Equal(expected, Read(log, srv.URL))
}

func TestReadNegative(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer srv.Close()
	assert := assert.New(t)
	log := logging.GetLogger()

	expected := []models.MMSData{}

	assert.Equal(expected, Read(log, srv.URL))
}
