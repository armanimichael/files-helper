package template

import (
	"os"
	"testing"
)

func Test_GenerateTestProject(t *testing.T) {
	GenerateTestProject()
	if err := os.RemoveAll(testFolder); err != nil {
		t.Fatal(err)
	}
}
