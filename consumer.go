package goHPapi

const (
	baseURLFedeperin      string = "https://potterapi-fedeperin.vercel.app/"
	baseURLPotterhead     string = "https://potterhead-api.vercel.app/api"
	endpointBooksURL      string = "/books"
	endpointMoviesURL     string = "/movies"
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

type movieConsumerPotterhead struct {
	Number           string   `json:"serial"`
	Title            string   `json:"title"`
	Summary          string   `json:"summary"`
	Directors        []string `json:"directors"`
	ScreenWriters    []string `json:"screenwriters"`
	Producers        []string `json:"producers"`
	Cinematographers []string `json:"cinematographers"`
	Editors          []string `json:"editors"`
	Distributors     []string `json:"distributors"`
	MusicComposers   []string `json:"music_composers"`
	ReleaseDate      string   `json:"release_date"`
	RunningTime      string   `json:"running_time"`
	Budget           string   `json:"budget"`
	BoxOffice        string   `json:"box_office"`
	Rating           string   `json:"rating"`
	Trailer          string   `json:"trailer"`
	Poster           string   `json:"poster"`
	Wiki             string   `json:"wiki"`
}

type characterConsumerFedeperin struct {
	ID            float64  `json:"index"`
	FullName      string   `json:"fullName"`
	NickName      string   `json:"nickname"`
	HogwartsHouse string   `json:"hogwartsHouse"`
	InterpretedBy string   `json:"interpretedBy"`
	Children      []string `json:"children"`
	Image         string   `json:"image"`
	Birthdate     string   `json:"birthdate"`
}

type houseConsumerFedeperin struct {
	ID      float64  `json:"index"`
	House   string   `json:"house"`
	Emoji   string   `json:"emoji"`
	Founder string   `json:"founder"`
	Colors  []string `json:"colors"`
	Animal  string   `json:"animal"`
}

type spellConsumerFedeperin struct {
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
