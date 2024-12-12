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
