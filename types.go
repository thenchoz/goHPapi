package goHPapi

type Params struct {
	id     string
	search string
	max    string
	lang   string
}

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

type Character struct {
	ID            float64  `json:"index"`
	FullName      string   `json:"fullName"`
	NickName      string   `json:"nickname"`
	HogwartsHouse string   `json:"hogwartsHouse"`
	InterpretedBy string   `json:"interpretedBy"`
	Children      []string `json:"children"`
	Image         string   `json:"image"`
	Birthdate     string   `json:"birthdate"`
}

type House struct {
	ID      float64  `json:"index"`
	House   string   `json:"house"`
	Emoji   string   `json:"emoji"`
	Founder string   `json:"founder"`
	Colors  []string `json:"colors"`
	Animal  string   `json:"animal"`
}

type Spell struct {
	ID    float64 `json:"index"`
	Spell string  `json:"spell"`
	Use   string  `json:"use"`
}

type Quote struct {
	ID      string `json:"id"`
	Quote   string `json:"quote"`
	Speaker string `json:"speaker"`
	Story   string `json:"story"`
	Source  string `json:"source"`
}
