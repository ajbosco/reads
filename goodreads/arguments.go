package goodreads

import (
	"net/url"
)

// Arguments are a map of key/value query strings
type Arguments map[string]string

// ToURLValues converts Arguments into query strings
func (args Arguments) ToURLValues() url.Values {
	v := url.Values{}
	for key, value := range args {
		v.Set(key, value)
	}
	return v
}
