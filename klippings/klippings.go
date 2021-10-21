package klippings

import (
	"io/fs"
)

type Klip struct{}

func NewKlipFromFS(filesystem fs.FS) ([]Klip, error) {
	dir, err := fs.ReadDir(filesystem, ".")
	if err != nil {
		return nil, err
	}

	var klips []Klip
	for range dir {
		klips = append(klips, Klip{})
	}

	return klips, nil
}
