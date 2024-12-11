package goHPapi

import (
	"context"
)

func (hp *HPapi) FetchBook(ctx context.Context) (bc Book, err error) {

	var params = make(map[string]string)

	mainURL := baseURL + hp.ExportedParams.lang + endpointBooksURL

	if hp.ExportedParams.id == "" {
		mainURL += "/random"
	} else {
		params["index"] = hp.ExportedParams.id
	}

	bc, err = Fetch[Book](ctx, mainURL, params)
	return
}

func (hp *HPapi) FetchBooks(ctx context.Context) (bc []Book, err error) {

	var params = make(map[string]string)

	mainURL := baseURL + hp.ExportedParams.lang + endpointBooksURL

	if hp.ExportedParams.search != "" {
		params["search"] = hp.ExportedParams.search
	}
	if hp.ExportedParams.max != "" {
		params["max"] = hp.ExportedParams.max
	}

	bc, err = Fetch[[]Book](ctx, mainURL, params)
	return
}

func (hp *HPapi) FetchCharacter(ctx context.Context) (cc Character, err error) {

	var params = make(map[string]string)

	mainURL := baseURL + hp.ExportedParams.lang + endpointCharactersURL

	if hp.ExportedParams.id == "" {
		mainURL += "/random"
	} else {
		params["index"] = hp.ExportedParams.id
	}

	cc, err = Fetch[Character](ctx, mainURL, params)
	return
}

func (hp *HPapi) FetchCharacters(ctx context.Context) (cc []Character, err error) {

	var params = make(map[string]string)

	mainURL := baseURL + hp.ExportedParams.lang + endpointBooksURL

	if hp.ExportedParams.search != "" {
		params["search"] = hp.ExportedParams.search
	}
	if hp.ExportedParams.max != "" {
		params["max"] = hp.ExportedParams.max
	}

	cc, err = Fetch[[]Character](ctx, mainURL, params)
	return
}

func (hp *HPapi) FetchHouse(ctx context.Context) (hc House, err error) {

	var params = make(map[string]string)

	mainURL := baseURL + hp.ExportedParams.lang + endpointHousesURL

	if hp.ExportedParams.id == "" {
		mainURL += "/random"
	} else {
		params["index"] = hp.ExportedParams.id
	}

	hc, err = Fetch[House](ctx, mainURL, params)
	return
}

func (hp *HPapi) FetchHouses(ctx context.Context) (hc []House, err error) {

	var params = make(map[string]string)

	mainURL := baseURL + hp.ExportedParams.lang + endpointBooksURL

	if hp.ExportedParams.search != "" {
		params["search"] = hp.ExportedParams.search
	}
	if hp.ExportedParams.max != "" {
		params["max"] = hp.ExportedParams.max
	}

	hc, err = Fetch[[]House](ctx, mainURL, params)
	return
}

func (hp *HPapi) FetchSpell(ctx context.Context) (sc Spell, err error) {

	var params = make(map[string]string)

	mainURL := baseURL + hp.ExportedParams.lang + endpointSpellsURL

	if hp.ExportedParams.id == "" {
		mainURL += "/random"
	} else {
		params["index"] = hp.ExportedParams.id
	}

	sc, err = Fetch[Spell](ctx, mainURL, params)
	return
}

func (hp *HPapi) FetchSpells(ctx context.Context) (sc []Spell, err error) {

	var params = make(map[string]string)

	mainURL := baseURL + hp.ExportedParams.lang + endpointBooksURL

	if hp.ExportedParams.search != "" {
		params["search"] = hp.ExportedParams.search
	}
	if hp.ExportedParams.max != "" {
		params["max"] = hp.ExportedParams.max
	}

	sc, err = Fetch[[]Spell](ctx, mainURL, params)
	return
}

func (hp *HPapi) FetchQuote(ctx context.Context) (qc Quote, err error) {

	var params = make(map[string]string)

	mainURL := quoteURL

	if hp.ExportedParams.id != "" {
		mainURL += "/" + hp.ExportedParams.id
	}

	qc, err = Fetch[Quote](ctx, mainURL, params)
	return
}
