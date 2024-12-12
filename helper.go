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

func mergeMovie(movie1 movieConsumerPotterhead) (movie Movie, err error) {
	nb, err := strconv.Atoi(movie1.Number)
	if err != nil {
		return
	}
	movie = Movie{
		Number:           nb,
		Title:            movie1.Title,
		Summary:          movie1.Summary,
		Directors:        movie1.Directors,
		ScreenWriters:    movie1.ScreenWriters,
		Producers:        movie1.Producers,
		Cinematographers: movie1.Cinematographers,
		Editors:          movie1.Editors,
		Distributors:     movie1.Distributors,
		MusicComposers:   movie1.MusicComposers,
		ReleaseDate:      movie1.ReleaseDate,
		RunningTime:      movie1.RunningTime,
		Budget:           movie1.Budget,
		BoxOffice:        movie1.BoxOffice,
		Rating:           movie1.Rating,
		Trailer:          movie1.Trailer,
		Poster:           movie1.Poster,
		Wiki:             movie1.Wiki,
	}
	return
}

func mergeCharacter(char1 characterConsumerFedeperin) (char Character, err error) {
	char = Character{
		ID:            int(char1.ID),
		FullName:      char1.FullName,
		NickName:      char1.NickName,
		HogwartsHouse: char1.HogwartsHouse,
		InterpretedBy: char1.InterpretedBy,
		Children:      char1.Children,
		Image:         char1.Image,
		Birthdate:     char1.Birthdate,
	}
	return
}

func mergeHouse(house1 houseConsumerFedeperin) (house House, err error) {
	house = House{
		ID:      int(house1.ID),
		House:   house1.House,
		Emoji:   house1.Emoji,
		Founder: house1.Founder,
		Colors:  house1.Colors,
		Animal:  house1.Animal,
	}
	return
}

func mergeSpell(spell1 spellConsumerFedeperin) (spell Spell, err error) {
	spell = Spell{
		ID:    int(spell1.ID),
		Spell: spell1.Spell,
		Use:   spell1.Use,
	}
	return
}

func mergeQuote(quote1 quoteConsumer) (quote Quote, err error) {
	quote = Quote(quote1)
	return
}
