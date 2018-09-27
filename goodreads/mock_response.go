package goodreads

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"path/filepath"
)

func mockResponse(paths ...string) *httptest.Server {
	parts := []string{".", "testdata"}
	filename := filepath.Join(append(parts, paths...)...)

	mockData, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	return httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		rw.Write(mockData)
	}))
}
