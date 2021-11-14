package data

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

const api string = "https://loripsum.net/api"

func GetSamplePlaintext() []byte {
	res, err := http.Get(api + "/plaintext")
	if err != nil {
		log.Fatal(err)
	}
	checkForStatusCode(res.StatusCode)
	return readResponseBody(&res.Body)
}

func checkForStatusCode(code int) {
	httpCode := code
	if httpCode != http.StatusOK {
		log.Fatalf("status code: %d", httpCode)
	}
}

func readResponseBody(body *io.ReadCloser) []byte {
	defer (*body).Close()
	content, err := ioutil.ReadAll(*body)
	if err != nil {
		log.Fatal(err)
	}
	return content
}
