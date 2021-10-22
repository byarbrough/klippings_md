package klippings

import (
	"io"
	"io/fs"
	"strconv"
	"strings"
)

type Klip struct {
	Page int
}

func NewKlipFromFS(filesystem fs.FS) ([]Klip, error) {
	dir, err := fs.ReadDir(filesystem, ".")
	if err != nil {
		return nil, err
	}

	var klips []Klip
	for _, f := range dir {
		klip, err := getKlip(filesystem, f)
		if err != nil {
			return klips, err // do I actually need to fail here?
		}
		klips = append(klips, klip)
	}

	return klips, nil
}

func getKlip(filesystem fs.FS, f fs.DirEntry) (Klip, error) {
	klipFile, err := filesystem.Open(f.Name())

	if err != nil {
		return Klip{}, err
	}
	defer klipFile.Close()
	return newKlip(klipFile)
}

func newKlip(klipFile fs.File) (Klip, error) {
	klipData, err := io.ReadAll(klipFile)
	if err != nil {
		return Klip{}, err
	}

	page_s := strings.SplitAfter(string(klipData), "- Your Highlight on page ")[1]
	page_s = strings.Split(page_s, " |")[0]
	page, err := strconv.Atoi(page_s)
	if err != nil {
		return Klip{}, err
	}

	klip := Klip{Page: page}

	return klip, nil
}
