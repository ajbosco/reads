package goodreads

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListShelves(t *testing.T) {
	c, err := NewClient(&Config{})
	assert.NoError(t, err)

	c.baseURI = mockResponse("shelves.xml").URL

	shelves, err := c.ListShelves("test-user")
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, 3, len(shelves))
	assert.Equal(t, "currently-reading", shelves[1].Name)
	assert.Equal(t, "5", shelves[0].BookCount)
}

func TestListShelfBooks(t *testing.T) {
	c, err := NewClient(&Config{})
	assert.NoError(t, err)

	c.baseURI = mockResponse("shelf_books.xml").URL

	books, err := c.ListShelfBooks("test-user", "test-shelf")
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, 2, len(books))
	assert.Equal(t, "Old Man's War (Old Man's War, #1)", books[1].Title)
	assert.Equal(t, "4.27", books[0].AvgRating)
	assert.Equal(t, "https://www.goodreads.com/book/show/36111562-senlin-ascends", books[0].Link)
}
