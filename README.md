# goHPapi

Un-official golang wrapper for multiple Harry Potter API

[PotterAPI](https://github.com/fedeperin/potterapi)
[Harry Potter API](https://github.com/joeltgray/HarryPotterAPI/tree/main?tab=readme-ov-file)

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
  api.FetchBook(context.Background())		(Book		, error)
  api.FetchBooks(context.Background())		([]Book		, error)
  api.FetchCharacter(context.Background())	(Character	, error)
  api.FetchCharacters(context.Background())	([]Character, error)
  api.FetchHouse(context.Background())		(House		, error)
  api.FetchHouses(context.Background())		([]House	, error)
  api.FetchSpell(context.Background())		(Spell		, error)
  api.FetchSpells(context.Background())		([]Spell	, error)
  api.FetchQuote(context.Background())		(Quote		, error)
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
	ID            float64
	Number        int    
	Title         string 
	OriginalTitle string 
	ReleaseDate   string 
	Description   string 
	Pages         int    
	Cover         string 
}

type Character struct {
	ID            float64 
	FullName      string  
	NickName      string  
	HogwartsHouse string  
	InterpretedBy string  
	Children      []string
	Image         string  
	Birthdate     string  
}

type House struct {
	ID      float64 
	House   string  
	Emoji   string  
	Founder string  
	Colors  []string
	Animal  string  
}

type Spell struct {
	ID    float64
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
