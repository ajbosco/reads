package goodreads

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetCurrentUserID(t *testing.T) {
	c, err := NewClient(&Config{})
	assert.NoError(t, err)

	c.baseURI = mockResponse("oauth_response.xml").URL

	user, err := c.GetCurrentUserID()
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, "12345", user.ID)
}
