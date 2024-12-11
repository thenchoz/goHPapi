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
  book       , err := api.FetchBook(context.Background())
  character  , err := api.FetchCharacter(context.Background())
  house      , err := api.FetchHouse(context.Background())
  spell      , err := api.FetchSpell(context.Background())
  quote      , err := api.FetchQuote(context.Background())
```
Config available methods:
```go
  api.SetID(id string)
  api.SetSearch(keyword string)
  api.SetLang(lang string)      // default  "en"
  api.Reset()
```
Response Struct
```go
type bookConsumer struct {
	ID            float64
	Number        int    
	Title         string 
	OriginalTitle string 
	ReleaseDate   string 
	Description   string 
	Pages         int    
	Cover         string 
}

type characterConsumer struct {
	ID            float64 
	FullName      string  
	NickName      string  
	HogwartsHouse string  
	InterpretedBy string  
	Children      []string
	Image         string  
	Birthdate     string  
}

type houseConsumer struct {
	ID      float64 
	House   string  
	Emoji   string  
	Founder string  
	Colors  []string
	Animal  string  
}

type spellConsumer struct {
	ID    float64
	Spell string 
	Use   string 
}

type quoteConsumer struct {
	ID      string
	Quote   string
	Speaker string
	Story   string
	Source  string
}
