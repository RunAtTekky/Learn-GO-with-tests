package poker_test

import (
	"io"
	"os"
	"testing"

	poker "github.com/runattekky/go-app"
)

func TestTapeWrite(t *testing.T) {
	file, clean := createTempFile(t, "12345")
	defer clean()

	tape := poker.Tape{file.(*os.File)}
	tape.Write([]byte("abc"))

	file.Seek(0, io.SeekStart)
	newFileContents, _ := io.ReadAll(file)

	got := string(newFileContents)
	want := "abc"

	if got != want {
		t.Errorf("got %s but want %s", got, want)
	}
}
