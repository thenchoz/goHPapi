package goHPapi

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strconv"
)

type HPapi struct {
	ExportedParams Params
}

func New() *HPapi {
	return &HPapi{Params{lang: "en"}}
}

func (hp *HPapi) SetID(id string) {
	hp.ExportedParams.id = id
}

func (hp *HPapi) SetSearch(search string) {
	hp.ExportedParams.search = search
}

func (hp *HPapi) SetMax(max int) {
	hp.ExportedParams.max = strconv.Itoa(max)
}

func (hp *HPapi) SetLang(lang string) {
	hp.ExportedParams.lang = lang
}

func (hp *HPapi) Reset() {
	hp.ExportedParams = Params{lang: "en"}
}

func Fetch[C any](ctx context.Context, u string, params map[string]string) (consumer C, err error) {
	var reqUrl *url.URL
	reqUrl, err = url.Parse(u)
	if err != nil {
		return
	}

	query := reqUrl.Query()
	for key, value := range params {
		query.Set(key, value)
	}
	reqUrl.RawQuery = query.Encode()

	client := &http.Client{}
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, reqUrl.String(), nil)
	if err != nil {
		return
	}

	resp, err := client.Do(req)
	if err != nil {
		return
	}

	info, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}

	err = json.Unmarshal(info, &consumer)
	return
}
