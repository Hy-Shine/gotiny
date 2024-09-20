package image

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDownload(t *testing.T) {
	url := "https://go.dev/images/gophers/ladder.svg"
	err := Download(url, "./test.svg")
	if assert.Nil(t, err) {
		os.Remove("./test.svg")
	}
}
