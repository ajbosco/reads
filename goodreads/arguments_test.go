package goodreads

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestArguments(t *testing.T) {

	tt := []struct {
		name        string
		args        Arguments
		queryString string
	}{
		{
			name:        "single argument",
			args:        Arguments{"title": "Alice in Wonderland"},
			queryString: "title=Alice+in+Wonderland",
		},
		{
			name:        "multiple argument",
			args:        Arguments{"title": "Alice in Wonderland", "author": "Lewis Carroll"},
			queryString: "author=Lewis+Carroll&title=Alice+in+Wonderland",
		},
	}

	for _, ts := range tt {
		t.Run(ts.name, func(t *testing.T) {
			queryString := ts.args.ToURLValues().Encode()
			assert.Equal(t, ts.queryString, queryString)
		})
	}

}
