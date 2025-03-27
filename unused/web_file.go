package utils

import (
	"encoding/json"
	"io"
	"net/http"
)

type WebFileGetter struct {
	BaseURL string
}

func (wfg *WebFileGetter) GetJSON(path string, dest any) error {

	resp, err := http.Get(wfg.BaseURL + path)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	dataBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	return json.Unmarshal(dataBytes, dest)
}
