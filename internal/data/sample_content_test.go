package data

import "testing"

func Test_GetSamplePlaintext(t *testing.T) {
	result := GetSamplePlaintext()
	if len(result) <= 0 {
		t.Fatal("no content found in api")
	}
}
