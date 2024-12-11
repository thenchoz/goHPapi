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
