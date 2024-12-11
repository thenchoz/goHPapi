package goHPapi

import (
	"context"
	"encoding/json"
	"errors"
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

func mergeBook(book1 bookConsumerFedeperin, book2 bookConsumerPotterhead) (book Book, err error) {
	if book2.Title == "" {
		// book2 not found
		book = Book{
			Number:        book1.Number,
			Title:         book1.Title,
			OriginalTitle: book1.OriginalTitle,
			ReleaseDate:   book1.ReleaseDate,
			Description:   book1.Description,
			Pages:         book1.Pages,
			Cover:         []string{book1.Cover},
		}
		return
	}
	nb, err := strconv.Atoi(book2.Number)
	if err != nil {
		return
	}
	if book1.Title == "" {
		// book1 not found
		var page int
		page, err = strconv.Atoi(book2.Pages)
		if err != nil {
			return
		}
		book = Book{
			Number:        nb,
			Title:         book2.Title,
			OriginalTitle: book2.Title,
			ReleaseDate:   book2.ReleaseDate,
			Description:   book2.Summary,
			Summary:       book2.Summary,
			Pages:         page,
			Cover:         []string{book2.Cover},
			Dedication:    book2.Dedication,
			Wiki:          book2.Wiki,
		}
		return
	}
	if book1.Number != nb {
		err = errors.New("uncompatible books: " + strconv.Itoa(book1.Number) + " and " + book2.Number)
		return
	}

	book = Book{
		Number:        book1.Number, // type int
		Title:         book1.Title,  // with language
		OriginalTitle: book1.OriginalTitle,
		ReleaseDate:   book2.ReleaseDate, // better date format (in my opinion)
		Description:   book1.Description, // with language
		Summary:       book2.Summary,     // in english
		Pages:         book1.Pages,       // type int
		Cover:         []string{book1.Cover, book2.Cover},
		Dedication:    book2.Dedication,
		Wiki:          book2.Wiki,
	}
	return
}
