# goHPapi

Un-official golang wrapper for multiple Harry Potter API. Concats data from multiple API on some request.

[PotterAPI](https://github.com/fedeperin/potterapi)

Books (merged), Characters, Houses and Spell

[Potterhead API](https://github.com/AcidOP/potterhead-api)

Books (merged), Movies

[Harry Potter API](https://github.com/joeltgray/HarryPotterAPI/tree/main?tab=readme-ov-file)

Quotes

## Install
Add ```github.com/thenchoz/goHPapi```to your go.mod(recommended) or:
```
EXPORT GO11MODULE=off
go get -u github.com/thenchoz/goHPapi
```

## Basic Usage
```go
import (
  "context"
  "github.com/thenchoz/goHPapi"
)

func main(){
  api := goHPapi.New()
  book, err := api.FetchBook(context.Background())
}
```
Available fetch commands:
```go
  api.FetchBook(Context)		(Book		, error)
  api.FetchBooks(Context)		([]Book		, error)
  api.FetchMovie(Context)		(Movie		, error)
  api.FetchMovies(Context)		([]Movie	, error)
  api.FetchCharacter(Context)	(Character	, error)
  api.FetchCharacters(Context)	([]Character, error)
  api.FetchHouse(Context)		(House		, error)
  api.FetchHouses(Context)		([]House	, error)
  api.FetchSpell(Context)		(Spell		, error)
  api.FetchSpells(Context)		([]Spell	, error)
  api.FetchQuote(Context)		(Quote		, error)
```
Config available methods:
```go
  api.SetID(id string)
  api.SetSearch(keyword string)
  api.SetMax(max int)
  api.SetLang(lang string)      // default  "en"
  api.Reset()
```
Response Struct
```go
type Book struct {
	Number        int
	Title         string
	OriginalTitle string
	ReleaseDate   string
	Description   string
	Summary       string
	Pages         int
	Cover         []string
	Dedication    string
	Wiki          string
}

type Movie struct {
	Number           int
	Title            string
	Summary          string
	Directors        []string
	ScreenWriters    []string
	Producers        []string
	Cinematographers []string
	Editors          []string
	Distributors     []string
	MusicComposers   []string
	ReleaseDate      string
	RunningTime      string
	Budget           string
	BoxOffice        string
	Rating           string
	Trailer          string
	Poster           string
	Wiki             string
}

type Character struct {
	ID            int
	FullName      string
	NickName      string
	HogwartsHouse string
	InterpretedBy string
	Children      []string
	Image         string
	Birthdate     string
}

type House struct {
	ID      int
	House   string
	Emoji   string
	Founder string
	Colors  []string
	Animal  string
}

type Spell struct {
	ID    int
	Spell string
	Use   string
}

type Quote struct {
	ID      string
	Quote   string
	Speaker string
	Story   string
	Source  string
}
```
