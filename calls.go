package goHPapi

import (
	"context"
)

func (hp *HPapi) FetchBook(ctx context.Context) (bc bookConsumer, err error) {

	var params = make(map[string]string)

	mainURL := baseURL + hp.ExportedParams.lang + endpointBooksURL

	if hp.ExportedParams.id == "" {
		mainURL += "/random"
	} else {
		params["index"] = hp.ExportedParams.id
	}

	bc, err = Fetch[bookConsumer](ctx, mainURL, params)
	return
}

func (hp *HPapi) FetchCharacter(ctx context.Context) (cc characterConsumer, err error) {

	var params = make(map[string]string)

	mainURL := baseURL + hp.ExportedParams.lang + endpointCharactersURL

	if hp.ExportedParams.id == "" {
		mainURL += "/random"
	} else {
		params["index"] = hp.ExportedParams.id
	}

	cc, err = Fetch[characterConsumer](ctx, mainURL, params)
	return
}

func (hp *HPapi) FetchHouse(ctx context.Context) (hc houseConsumer, err error) {

	var params = make(map[string]string)

	mainURL := baseURL + hp.ExportedParams.lang + endpointHousesURL

	if hp.ExportedParams.id == "" {
		mainURL += "/random"
	} else {
		params["index"] = hp.ExportedParams.id
	}

	hc, err = Fetch[houseConsumer](ctx, mainURL, params)
	return
}

func (hp *HPapi) FetchSpell(ctx context.Context) (sc spellConsumer, err error) {

	var params = make(map[string]string)

	mainURL := baseURL + hp.ExportedParams.lang + endpointSpellsURL

	if hp.ExportedParams.id == "" {
		mainURL += "/random"
	} else {
		params["index"] = hp.ExportedParams.id
	}

	sc, err = Fetch[spellConsumer](ctx, mainURL, params)
	return
}

func (hp *HPapi) FetchQuote(ctx context.Context) (qc quoteConsumer, err error) {

	var params = make(map[string]string)

	mainURL := quoteURL

	if hp.ExportedParams.id != "" {
		mainURL += "/" + hp.ExportedParams.id
	}

	qc, err = Fetch[quoteConsumer](ctx, mainURL, params)
	return
}
