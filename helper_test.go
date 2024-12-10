package goHPapi

import (
	"testing"
)

func Test_Params(t *testing.T) {
	testID := "123"
	testSearch := "search"
	testLang := "language"

	api := New()

	if !(api.ExportedParams.id == "" &&
		api.ExportedParams.search == "" &&
		api.ExportedParams.lang == "en") {
		t.Fatalf("Wrong base api")
	}

	api.setID(testID)
	api.setSearch(testSearch)
	api.setLang(testLang)

	if !(api.ExportedParams.id == testID &&
		api.ExportedParams.search == testSearch &&
		api.ExportedParams.lang == testLang) {
		t.Fatalf("Wrong setup api")
	}

	api.reset()

	if !(api.ExportedParams.id == "" &&
		api.ExportedParams.search == "" &&
		api.ExportedParams.lang == "en") {
		t.Fatalf("Wrong reset api")
	}
}
