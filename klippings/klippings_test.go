package klippings_test

import (
	"testing"
	"testing/fstest"

	"github.com/byarbrough/klippings_md/klippings"
)

func TestNewKlip(t *testing.T) {
	fs := fstest.MapFS{
		"one klip.txt":  {Data: []byte("Starship Troopers (Heinlein, Robert A.)")},
		"no author.txt": {Data: []byte("Invent and Wander")},
	}

	klip_files := klippings.NewKlipFromFS(fs)

	if len(klip_files) != len(fs) {
		t.Errorf("got %d klip files, want %d files", len(klip_files), len(fs))
	}

}
