package example

import "testing"

func TestReverse(t *testing.T) {
	// unit test for Reverse
	// simple example of unit test in Go
	got := Reverse("x")
	if got != "x" {
		t.Errorf(`Reverse("x") == %v, want "x"`, got)
	}

	got = Reverse("Jane")
	if got != "enaJ" {
		t.Errorf(`Reverse("Jane") == %v, want "enaJ"`, got)
	}

	got = Reverse("Joe")
	if got != "eoJ" {
		t.Errorf(`Reverse("Joe") == %v, want "eoJ"`, got)
	}

	got = Reverse("日本語")
	if got != "語本日" {
		t.Errorf(`Reverse("日本語") == %v, want "語本日"`, got)
	}
}
