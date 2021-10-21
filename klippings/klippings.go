package klippings

import (
	"io/fs"
	"testing/fstest"
)

type Klip struct{}

func NewKlipFromFS(filesystem fstest.MapFS) []Klip {
	dir, _ := fs.ReadDir(filesystem, ".")

	var klips []Klip
	for range dir {
		klips = append(klips, Klip{})
	}

	return klips
}
