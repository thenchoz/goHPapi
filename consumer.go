package goHPapi

const (
	baseURLFedeperin      string = "https://potterapi-fedeperin.vercel.app/"
	baseURLPotterhead     string = "https://potterhead-api.vercel.app/api"
	endpointBooksURL      string = "/books"
	endpointCharactersURL string = "/characters"
	endpointHousesURL     string = "/houses"
	endpointSpellsURL     string = "/spells"
	quoteURL              string = "https://api.portkey.uk/quote"
)

type bookConsumerFedeperin struct {
	ID            float64 `json:"index"`
	Number        int     `json:"number"`
	Title         string  `json:"title"`
	OriginalTitle string  `json:"originalTitle"`
	ReleaseDate   string  `json:"releaseDate"`
	Description   string  `json:"description"`
	Pages         int     `json:"pages"`
	Cover         string  `json:"cover"`
}

type bookConsumerPotterhead struct {
	Number      string `json:"serial"`
	Title       string `json:"title"`
	Summary     string `json:"summary"`
	ReleaseDate string `json:"release_date"`
	Dedication  string `json:"dedication"`
	Pages       string `json:"pages"`
	Cover       string `json:"cover"`
	Wiki        string `json:"wiki"`
}

type CharacterConsumerFedeperin struct {
	ID            float64  `json:"index"`
	FullName      string   `json:"fullName"`
	NickName      string   `json:"nickname"`
	HogwartsHouse string   `json:"hogwartsHouse"`
	InterpretedBy string   `json:"interpretedBy"`
	Children      []string `json:"children"`
	Image         string   `json:"image"`
	Birthdate     string   `json:"birthdate"`
}

type HouseConsumerFedeperin struct {
	ID      float64  `json:"index"`
	House   string   `json:"house"`
	Emoji   string   `json:"emoji"`
	Founder string   `json:"founder"`
	Colors  []string `json:"colors"`
	Animal  string   `json:"animal"`
}

type SpellConsumerFedeperin struct {
	ID    float64 `json:"index"`
	Spell string  `json:"spell"`
	Use   string  `json:"use"`
}

type QuoteConsumer struct {
	ID      string `json:"id"`
	Quote   string `json:"quote"`
	Speaker string `json:"speaker"`
	Story   string `json:"story"`
	Source  string `json:"source"`
}
