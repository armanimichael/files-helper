package data

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

const api string = "https://loripsum.net/api"

func GetSamplePlaintext(paragraphs int) []byte {
	url := fmt.Sprintf("%s/plaintext/%d", api, paragraphs)
	return fetchBytes(url)
}

func GetSampleHtmltext(paragraphs int) []byte {
	url := fmt.Sprintf("%s/ul/ol/code/dl/headers/link/decorate/bq/%d", api, paragraphs)
	return fetchBytes(url)
}

func fetchBytes(url string) []byte {
	res, err := http.Get(url)
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
	content, err := ioutil.ReadAll(*body)
	if err != nil {
		log.Fatal(err)
	}
	if err := (*body).Close(); err != nil {
		log.Fatal(err)
	}
	return content
}
