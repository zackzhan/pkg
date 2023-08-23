package text

import (
	"fmt"
	"io"
	"os"
	"testing"

	"github.com/gabriel-vasile/mimetype"
)

func TestContentType(t *testing.T) {
	file, err := os.Open("../testdata/test.svg")
	if err != nil {
		return
	}
	defer file.Close()

	all, err := io.ReadAll(file)
	if err != nil {
		return
	}

	fmt.Println(mimetype.Detect(all).String())
}
