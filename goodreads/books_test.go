package goodreads

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearchBooks(t *testing.T) {
	c, err := NewClient(&Config{})
	assert.NoError(t, err)

	c.baseURI = mockResponse("search.xml").URL

	books, err := c.SearchBooks("golang")
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, 3, len(books))
	assert.Equal(t, "Go Programming Blueprints - Solving Development Challenges with Golang", books[0].Title)
	assert.Equal(t, "astaxie", books[1].Author)
	assert.Equal(t, "3.42", books[1].AvgRating)
}
