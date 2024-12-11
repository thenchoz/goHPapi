package goHPapi

import (
	"testing"
)

func Test_Params(t *testing.T) {
	testID := "123"
	testSearch := "search"
	testMax := 5
	testLang := "language"

	api := New()

	if !(api.ExportedParams.id == "" &&
		api.ExportedParams.search == "" &&
		api.ExportedParams.lang == "en") {
		t.Fatalf("Wrong base api")
	}

	api.SetID(testID)
	api.SetSearch(testSearch)
	api.SetMax(testMax)
	api.SetLang(testLang)

	if !(api.ExportedParams.id == testID &&
		api.ExportedParams.search == testSearch &&
		api.ExportedParams.max == "5" &&
		api.ExportedParams.lang == testLang) {
		t.Fatalf("Wrong setup api")
	}

	api.Reset()

	if !(api.ExportedParams.id == "" &&
		api.ExportedParams.search == "" &&
		api.ExportedParams.max == "" &&
		api.ExportedParams.lang == "en") {
		t.Fatalf("Wrong reset api")
	}
}

func Test_Merge_Book(t *testing.T) {
	book1 := bookConsumerFedeperin{
		ID:            1,
		Number:        1,
		Title:         "title",
		OriginalTitle: "ogTitle",
		ReleaseDate:   "release",
		Description:   "desc",
		Pages:         10,
		Cover:         "cover",
	}

	var book2 bookConsumerPotterhead

	book, err := mergeBook(book1, book2)
	if err != nil {
		t.Fatal(err)
	}

	if !(book.Number == book1.Number &&
		book.Title == book1.Title &&
		book.OriginalTitle == book1.OriginalTitle &&
		book.ReleaseDate == book1.ReleaseDate &&
		book.Description == book1.Description &&
		book.Summary == "" &&
		book.Pages == book1.Pages &&
		book.Cover[0] == book1.Cover &&
		book.Dedication == "" &&
		book.Wiki == "") {
		t.Fatalf("Wrong partial book from book 1")
	}

	book2 = bookConsumerPotterhead{
		Number:      "1",
		Title:       "t2",
		Summary:     "summary",
		ReleaseDate: "rel2",
		Dedication:  "dedi",
		Pages:       "10",
		Cover:       "cov2",
		Wiki:        "wiki",
	}

	book, err = mergeBook(book1, book2)
	if err != nil {
		t.Fatal(err)
	}

	if !(book.Number == book1.Number &&
		book.Title == book1.Title &&
		book.OriginalTitle == book1.OriginalTitle &&
		book.ReleaseDate == book2.ReleaseDate &&
		book.Description == book1.Description &&
		book.Summary == book2.Summary &&
		book.Pages == book1.Pages &&
		book.Cover[0] == book1.Cover &&
		book.Cover[1] == book2.Cover &&
		book.Dedication == book2.Dedication &&
		book.Wiki == book2.Wiki) {
		t.Fatalf("Wrong merged book")
	}

	book1 = bookConsumerFedeperin{}

	book, err = mergeBook(book1, book2)
	if err != nil {
		t.Fatal(err)
	}

	if !(book.Number == 1 &&
		book.Title == book2.Title &&
		book.OriginalTitle == book2.Title &&
		book.ReleaseDate == book2.ReleaseDate &&
		book.Description == book2.Summary &&
		book.Summary == book2.Summary &&
		book.Pages == 10 &&
		book.Cover[0] == book2.Cover &&
		book.Dedication == book2.Dedication &&
		book.Wiki == book2.Wiki) {
		t.Fatalf("Wrong partial book from book 2")
	}
}
