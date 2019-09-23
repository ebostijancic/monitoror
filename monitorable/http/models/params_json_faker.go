//+build faker

package models

import (
	"encoding/json"
	"regexp"

	"github.com/monitoror/monitoror/models/tiles"
)

type (
	HttpJsonParams struct {
		Url           string `json:"url" query:"url"`
		Key           string `json:"key" query:"key"`
		Regex         string `json:"regex" query:"regex"`
		StatusCodeMin *int   `json:"statusCodeMin" query:"statusCodeMin"`
		StatusCodeMax *int   `json:"statusCodeMax" query:"statusCodeMax"`

		Status  tiles.TileStatus `json:"status" query:"status"`
		Message string           `json:"message" query:"message"`
	}
)

func (p *HttpJsonParams) IsValid() bool {
	if !isValid(p.Url, p) {
		return false
	}

	if !isValidKey(p) {
		return false
	}

	return isValidRegex(p)
}

func (p *HttpJsonParams) GetStatusCodes() (min int, max int) {
	return getStatusCodes(p.StatusCodeMin, p.StatusCodeMax)
}

func (p *HttpJsonParams) GetRegex() string          { return p.Regex }
func (p *HttpJsonParams) GetRegexp() *regexp.Regexp { return getRegexp(p.GetRegex()) }

func (p *HttpJsonParams) GetKey() string { return p.Key }
func (p *HttpJsonParams) GetUnmarshaller() func(data []byte, v interface{}) error {
	return json.Unmarshal
}

func (p *HttpJsonParams) GetStatus() tiles.TileStatus { return p.Status }
func (p *HttpJsonParams) GetMessage() string          { return p.Message }