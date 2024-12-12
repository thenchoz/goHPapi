package goHPapi

import (
	"context"
	"math/rand/v2"
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

		if err != nil {
			return
		}
	}

	return
}

func (hp *HPapi) FetchMovie(ctx context.Context) (movie Movie, err error) {

	mainURL := baseURLPotterhead + endpointMoviesURL + "/"
	if hp.ExportedParams.id == "" {
		mainURL += strconv.Itoa(rand.IntN(8) + 1)
	} else {
		mainURL += hp.ExportedParams.id
	}

	mc, err := Fetch[movieConsumerPotterhead](ctx, mainURL, make(map[string]string))
	if err != nil {
		return
	}

	movie, err = mergeMovie(mc)
	return
}

func (hp *HPapi) FetchMovies(ctx context.Context) (movies []Movie, err error) {

	mainURL := baseURLPotterhead + endpointMoviesURL

	mc, err := Fetch[[]movieConsumerPotterhead](ctx, mainURL, make(map[string]string))
	if err != nil {
		return
	}

	movies = make([]Movie, len(mc))
	for i, m := range mc {
		movies[i], err = mergeMovie(m)
		if err != nil {
			return
		}
	}

	return
}

func (hp *HPapi) FetchCharacter(ctx context.Context) (character Character, err error) {

	var params = make(map[string]string)

	mainURL := baseURLFedeperin + hp.ExportedParams.lang + endpointCharactersURL

	if hp.ExportedParams.id == "" {
		mainURL += "/random"
	} else {
		params["index"] = hp.ExportedParams.id
	}

	cc, err := Fetch[characterConsumerFedeperin](ctx, mainURL, params)
	if err != nil {
		return
	}

	character, err = mergeCharacter(cc)
	return
}

func (hp *HPapi) FetchCharacters(ctx context.Context) (characters []Character, err error) {

	var params = make(map[string]string)

	mainURL := baseURLFedeperin + hp.ExportedParams.lang + endpointCharactersURL

	if hp.ExportedParams.search != "" {
		params["search"] = hp.ExportedParams.search
	}
	if hp.ExportedParams.max != "" {
		params["max"] = hp.ExportedParams.max
	}

	cc, err := Fetch[[]characterConsumerFedeperin](ctx, mainURL, params)
	if err != nil {
		return
	}

	characters = make([]Character, len(cc))
	for i, c := range cc {
		characters[i], err = mergeCharacter(c)
		if err != nil {
			return
		}
	}
	return
}

func (hp *HPapi) FetchHouse(ctx context.Context) (house House, err error) {

	var params = make(map[string]string)

	mainURL := baseURLFedeperin + hp.ExportedParams.lang + endpointHousesURL

	if hp.ExportedParams.id == "" {
		mainURL += "/random"
	} else {
		params["index"] = hp.ExportedParams.id
	}

	hc, err := Fetch[houseConsumerFedeperin](ctx, mainURL, params)
	if err != nil {
		return
	}

	house, err = mergeHouse(hc)
	return
}

func (hp *HPapi) FetchHouses(ctx context.Context) (houses []House, err error) {

	var params = make(map[string]string)

	mainURL := baseURLFedeperin + hp.ExportedParams.lang + endpointHousesURL

	if hp.ExportedParams.search != "" {
		params["search"] = hp.ExportedParams.search
	}
	if hp.ExportedParams.max != "" {
		params["max"] = hp.ExportedParams.max
	}

	hc, err := Fetch[[]houseConsumerFedeperin](ctx, mainURL, params)
	if err != nil {
		return
	}

	houses = make([]House, len(hc))
	for i, h := range hc {
		houses[i], err = mergeHouse(h)
		if err != nil {
			return
		}
	}
	return
}

func (hp *HPapi) FetchSpell(ctx context.Context) (spell Spell, err error) {

	var params = make(map[string]string)

	mainURL := baseURLFedeperin + hp.ExportedParams.lang + endpointSpellsURL

	if hp.ExportedParams.id == "" {
		mainURL += "/random"
	} else {
		params["index"] = hp.ExportedParams.id
	}

	sc, err := Fetch[spellConsumerFedeperin](ctx, mainURL, params)
	if err != nil {
		return
	}

	spell, err = mergeSpell(sc)
	return
}

func (hp *HPapi) FetchSpells(ctx context.Context) (spells []Spell, err error) {

	var params = make(map[string]string)

	mainURL := baseURLFedeperin + hp.ExportedParams.lang + endpointSpellsURL

	if hp.ExportedParams.search != "" {
		params["search"] = hp.ExportedParams.search
	}
	if hp.ExportedParams.max != "" {
		params["max"] = hp.ExportedParams.max
	}

	sc, err := Fetch[[]spellConsumerFedeperin](ctx, mainURL, params)
	if err != nil {
		return
	}

	spells = make([]Spell, len(sc))
	for i, s := range sc {
		spells[i], err = mergeSpell(s)
		if err != nil {
			return
		}
	}
	return
}

func (hp *HPapi) FetchQuote(ctx context.Context) (quote Quote, err error) {

	var params = make(map[string]string)

	mainURL := quoteURL

	if hp.ExportedParams.id != "" {
		mainURL += "/" + hp.ExportedParams.id
	}

	qc, err := Fetch[quoteConsumer](ctx, mainURL, params)
	if err != nil {
		return
	}

	quote, err = mergeQuote(qc)
	return
}
