package archive

import "testing"

func Test_TarFolder(t *testing.T) {
	if err := TarFolder("_test_proj_backup"); err != nil {
		t.Fatal(err)
	}
}
