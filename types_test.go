package goHPapi

import (
	"encoding/json"
	"testing"
)

func Test_Book_Consumer(t *testing.T) {
	book := bookConsumer{
		ID:            123,
		Number:        1,
		Title:         "title",
		OriginalTitle: "ogTitle",
		ReleaseDate:   "date",
		Description:   "desc",
		Pages:         10,
		Cover:         "url",
	}
	bookjson, err := json.Marshal(book)
	if err != nil {
		t.Fatal(err)
	}

	var consumer bookConsumer
	err = json.Unmarshal(bookjson, &consumer)
	if err != nil {
		t.Fatal(err)
	}

	if !(book.ID == consumer.ID &&
		book.Number == consumer.Number &&
		book.OriginalTitle == consumer.OriginalTitle &&
		book.ReleaseDate == consumer.ReleaseDate &&
		book.Description == consumer.Description &&
		book.Pages == consumer.Pages &&
		book.Cover == consumer.Cover) {
		t.Fatalf("Unmatched Book")
	}
}
func Test_Character_Consumer(t *testing.T) {
	book := characterConsumer{
		ID:            123,
		FullName:      "fullName",
		NickName:      "nickName",
		HogwartsHouse: "hogwartsHouse",
		InterpretedBy: "interpretedBy",
		Children:      []string{},
		Image:         "image",
		Birthdate:     "birthdate",
	}
	bookjson, err := json.Marshal(book)
	if err != nil {
		t.Fatal(err)
	}

	var consumer characterConsumer
	err = json.Unmarshal(bookjson, &consumer)
	if err != nil {
		t.Fatal(err)
	}

	if !(book.ID == consumer.ID &&
		book.FullName == consumer.FullName &&
		book.NickName == consumer.NickName &&
		book.HogwartsHouse == consumer.HogwartsHouse &&
		book.InterpretedBy == consumer.InterpretedBy &&
		book.Image == consumer.Image &&
		book.Birthdate == consumer.Birthdate) {
		t.Fatalf("Unmatched Character")
	}
}
func Test_house_Consumer(t *testing.T) {
	book := houseConsumer{
		ID:      123,
		House:   "house",
		Emoji:   "emoji",
		Founder: "founder",
		Colors:  []string{"c1", "c2"},
		Animal:  "animal",
	}
	bookjson, err := json.Marshal(book)
	if err != nil {
		t.Fatal(err)
	}

	var consumer houseConsumer
	err = json.Unmarshal(bookjson, &consumer)
	if err != nil {
		t.Fatal(err)
	}

	if !(book.ID == consumer.ID &&
		book.House == consumer.House &&
		book.Emoji == consumer.Emoji &&
		book.Founder == consumer.Founder &&
		//book.Colors == consumer.Colors &&
		book.Animal == consumer.Animal) {
		t.Fatalf("Unmatched House")
	}
}

func Test_Spell_Consumer(t *testing.T) {
	book := spellConsumer{
		ID:    123,
		Spell: "spell",
		Use:   "use",
	}
	bookjson, err := json.Marshal(book)
	if err != nil {
		t.Fatal(err)
	}

	var consumer spellConsumer
	err = json.Unmarshal(bookjson, &consumer)
	if err != nil {
		t.Fatal(err)
	}

	if !(book.ID == consumer.ID &&
		book.Spell == consumer.Spell &&
		book.Use == consumer.Use) {
		t.Fatalf("Unmatched Spell")
	}
}

func Test_Quote_Consumer(t *testing.T) {
	book := quoteConsumer{
		ID:      "123",
		Quote:   "quote",
		Speaker: "speaker",
		Story:   "story",
		Source:  "source",
	}
	bookjson, err := json.Marshal(book)
	if err != nil {
		t.Fatal(err)
	}

	var consumer quoteConsumer
	err = json.Unmarshal(bookjson, &consumer)
	if err != nil {
		t.Fatal(err)
	}

	if !(book.ID == consumer.ID &&
		book.Quote == consumer.Quote &&
		book.Speaker == consumer.Speaker &&
		book.Story == consumer.Story &&
		book.Source == consumer.Source) {
		t.Fatalf("Unmatched Quote")
	}
}
