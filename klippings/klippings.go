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
		klip, err := getKlip(filesystem, f.Name())
		if err != nil {
			return klips, err // do I actually need to fail here?
		}
		klips = append(klips, klip)
	}

	return klips, nil
}

func getKlip(filesystem fs.FS, fileName string) (Klip, error) {
	klipFile, err := filesystem.Open(fileName)

	if err != nil {
		return Klip{}, err
	}
	defer klipFile.Close()
	return newKlip(klipFile)
}

func newKlip(klipFile io.Reader) (Klip, error) {
	klipData, err := io.ReadAll(klipFile)
	if err != nil {
		return Klip{}, err
	}

	pageS := strings.SplitAfter(string(klipData), "- Your Highlight on page ")[1]
	pageS = strings.Split(pageS, " |")[0]
	page, err := strconv.Atoi(pageS)
	if err != nil {
		return Klip{}, err
	}

	klip := Klip{Page: page}

	return klip, nil
}
