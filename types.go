package goHPapi

const (
	baseURL               string = "https://potterapi-fedeperin.vercel.app/"
	endpointBooksURL      string = "/books"
	endpointCharactersURL string = "/characters"
	endpointHousesURL     string = "/houses"
	endpointSpellsURL     string = "/spells"
	quoteURL              string = "https://api.portkey.uk/quote"
)

type Params struct {
	id     string
	search string
	lang   string
}

type bookConsumer struct {
	ID            float64 `json:"index"`
	Number        int     `json:"number"`
	Title         string  `json:"title"`
	OriginalTitle string  `json:"originalTitle"`
	ReleaseDate   string  `json:"releaseDate"`
	Description   string  `json:"description"`
	Pages         int     `json:"pages"`
	Cover         string  `json:"cover"`
}

type characterConsumer struct {
	ID            float64  `json:"index"`
	FullName      string   `json:"fullName"`
	NickName      string   `json:"nickname"`
	HogwartsHouse string   `json:"hogwartsHouse"`
	InterpretedBy string   `json:"interpretedBy"`
	Children      []string `json:"children"`
	Image         string   `json:"image"`
	Birthdate     string   `json:"birthdate"`
}

type houseConsumer struct {
	ID      float64  `json:"index"`
	House   string   `json:"house"`
	Emoji   string   `json:"emoji"`
	Founder string   `json:"founder"`
	Colors  []string `json:"colors"`
	Animal  string   `json:"animal"`
}

type spellConsumer struct {
	ID    float64 `json:"index"`
	Spell string  `json:"spell"`
	Use   string  `json:"use"`
}

type quoteConsumer struct {
	ID      string `json:"id"`
	Quote   string `json:"quote"`
	Speaker string `json:"speaker"`
	Story   string `json:"story"`
	Source  string `json:"source"`
}
