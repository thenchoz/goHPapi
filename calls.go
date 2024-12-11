package goHPapi

import (
	"context"
	"slices"
	"strconv"
	"sync"
)

func (hp *HPapi) FetchBook(ctx context.Context) (book Book, err error) {

	var params = make(map[string]string)

	mainURL := baseURLFedeperin + hp.ExportedParams.lang + endpointBooksURL
	if hp.ExportedParams.id == "" {
		mainURL += "/random"
	} else {
		params["index"] = hp.ExportedParams.id
	}
	bcf, err := Fetch[bookConsumerFedeperin](ctx, mainURL, params)
	if err != nil {
		return
	}

	// successive, manage random + index
	mainURL = baseURLPotterhead + endpointBooksURL + "/" + strconv.Itoa(bcf.Number)
	bcp, err := Fetch[bookConsumerPotterhead](ctx, mainURL, make(map[string]string))
	if err != nil {
		// manage non retreived book
		bcp = bookConsumerPotterhead{}
	}

	book, err = mergeBook(bcf, bcp)
	return
}

func (hp *HPapi) FetchBooks(ctx context.Context) (books []Book, err error) {

	var (
		params = make(map[string]string)
		bcp    []bookConsumerPotterhead
		bcf    []bookConsumerFedeperin
		wg     sync.WaitGroup
	)

	ctx, cancel := context.WithCancel(ctx)

	// parallel to improve speed
	wg.Add(1)
	go func() {
		defer wg.Done()

		mainURL := baseURLPotterhead + endpointBooksURL
		bcp, err = Fetch[[]bookConsumerPotterhead](ctx, mainURL, make(map[string]string))
		if err != nil {
			cancel()
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()

		mainURL := baseURLFedeperin + hp.ExportedParams.lang + endpointBooksURL
		if hp.ExportedParams.search != "" {
			params["search"] = hp.ExportedParams.search
		}
		if hp.ExportedParams.max != "" {
			params["max"] = hp.ExportedParams.max
		}

		bcf, err = Fetch[[]bookConsumerFedeperin](ctx, mainURL, params)
		if err != nil {
			cancel()
		}
	}()

	wg.Wait()
	books = make([]Book, len(bcf))
	for i, book1 := range bcf {
		idx := slices.IndexFunc(
			bcp,
			func(c bookConsumerPotterhead) bool {
				val, err := strconv.Atoi(c.Number)
				return err == nil && val == book1.Number
			},
		)

		var book2 bookConsumerPotterhead
		if idx != -1 {
			book2 = bcp[idx]
		}

		books[i], err = mergeBook(book1, book2)
	}

	return
}

func (hp *HPapi) FetchCharacter(ctx context.Context) (cc Character, err error) {

	var params = make(map[string]string)

	mainURL := baseURLFedeperin + hp.ExportedParams.lang + endpointCharactersURL

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

	mainURL := baseURLFedeperin + hp.ExportedParams.lang + endpointBooksURL

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

	mainURL := baseURLFedeperin + hp.ExportedParams.lang + endpointHousesURL

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

	mainURL := baseURLFedeperin + hp.ExportedParams.lang + endpointBooksURL

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

	mainURL := baseURLFedeperin + hp.ExportedParams.lang + endpointSpellsURL

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

	mainURL := baseURLFedeperin + hp.ExportedParams.lang + endpointBooksURL

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
