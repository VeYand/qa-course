package tests

import (
	"github.com/stretchr/testify/assert"
	"lab8/model"
	"testing"
)

func assertEquals(t *testing.T, expected, actual model.Product) {
	fields := []struct {
		name     string
		expected interface{}
		actual   interface{}
	}{
		{"CategoryID", expected.CategoryID, actual.CategoryID},
		{"Title", expected.Title, actual.Title},
		{"Content", expected.Content, actual.Content},
		{"Price", handleBigNumber(expected.Price), actual.Price},
		{"OldPrice", handleBigNumber(expected.OldPrice), actual.OldPrice},
		{"Status", expected.Status, actual.Status},
		{"Keywords", expected.Keywords, actual.Keywords},
		{"Description", expected.Description, actual.Description},
		{"Hit", expected.Hit, actual.Hit},
	}

	for _, field := range fields {
		assert.Equal(
			t,
			field.expected,
			field.actual,
			"%s mismatch",
			field.name,
		)
	}
}

func handleBigNumber(numStr string) string {
	if numStr == "99999999999999999999999999999999999999" {
		return "1e38"
	}
	if numStr == "-99999999999999999999999999999999999999" {
		return "-1e38"
	}

	return numStr
}
