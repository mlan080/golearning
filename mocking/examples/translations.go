package example

import (
	"encoding/csv"
	"io"
	"os"
)

// Translations holds translations for a user interface.
type Translations struct {
	m map[string]map[string]string
}

// LoadTranslations loads translations from a CSV file.
func LoadTranslations(path string) (*Translations, error) {
	m := map[string]map[string]string{
		"en": {},
		"de": {},
		"fr": {},
	}
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	r := csv.NewReader(file)
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		id := record[0]
		m["en"][id] = record[1]
		m["de"][id] = record[2]
		m["fr"][id] = record[3]
	}
	return &Translations{m}, nil
}

// Lookup returns a translation in the requested language.
func (t *Translations) Lookup(language, id string) string {
	return t.m[language][id]
}
