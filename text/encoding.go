package text

import "github.com/emersion/go-imap/utf7"

func Utf72Utf8(name string) string {
	dec := utf7.Encoding.NewDecoder()
	s, err := dec.String(name)
	if err != nil {
		return name
	}
	return s
}
