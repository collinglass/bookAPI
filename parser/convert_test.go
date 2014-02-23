package convert

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConvertSuccess(t *testing.T) {
	_, err := ExtractMetadata("test.epub")

	assert.Nil(t, err)
}

func TestConvertFailure(t *testing.T) {
	_, err := ExtractMetadata("testifyyyy.epub")

	assert.NotNil(t, err)
}

func TestMetadata(t *testing.T) {
	book, _ := ExtractMetadata("test.epub")

	lang := []string{"en"}

	assert.Equal(t, book.Metadata.language, lang)
}
