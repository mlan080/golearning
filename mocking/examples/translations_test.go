package example

import "testing"

func TestTranslations(t *testing.T) {
	// unit test for LoadTranslations and Lookup
	// it's hard to test one of these by itself, so we're doing both in one

	// example of using a “testdata” directory

	// to be able to test Lookup, we also need an actual file with translations, so we have one
	// just for testing:
	translations, err := LoadTranslations("testdata/translations.csv")
	if err != nil {
		t.Fatalf("LoadTranslations returned error: %v", err)
	}

	cases := []struct {
		language, id string
		want         string
	}{
		{"en", "yes", "Yes"},
		{"de", "thanks", "Danke schön!"},
		{"fr", "hello", "Salut!"},
	}
	for _, c := range cases {
		got := translations.Lookup(c.language, c.id)
		if got != c.want {
			t.Errorf("translations.Lookup(%v, %v) == %v, want %v",
				c.id, c.language, got, c.want)
		}
	}

	_, err = LoadTranslations("testdata/no-such-file")
	if err == nil {
		t.Error("LoadTranslations didn't return error for nonexitent file")
	}

	_, err = LoadTranslations("testdata/translations-invalid.csv")
	if err == nil {
		t.Error("LoadTranslations didn't return error for invalid file")
	}
}
