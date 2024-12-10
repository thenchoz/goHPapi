package goHPapi

import (
	"context"
	"testing"
)

func Test_Fetch_Book(t *testing.T) {
	api := New()
	_, err := api.FetchBook(context.Background())

	if err != nil {
		t.Fatal(err)
	}
}

func Test_Fetch_Character(t *testing.T) {
	api := New()
	_, err := api.FetchCharacter(context.Background())

	if err != nil {
		t.Fatal(err)
	}
}

func Test_Fetch_House(t *testing.T) {
	api := New()
	_, err := api.FetchHouse(context.Background())

	if err != nil {
		t.Fatal(err)
	}
}

func Test_Fetch_Spell(t *testing.T) {
	api := New()
	_, err := api.FetchSpell(context.Background())

	if err != nil {
		t.Fatal(err)
	}
}

func Test_Fetch_Quote(t *testing.T) {
	api := New()
	_, err := api.FetchQuote(context.Background())

	if err != nil {
		t.Fatal(err)
	}
}
