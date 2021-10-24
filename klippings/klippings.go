package klippings

import (
	"fmt"
	"io"
	"io/fs"
	"strconv"
	"strings"
	"time"
)

type Klip struct {
	Page     int
	Location string
	Time     time.Time
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

	line := strings.TrimPrefix(string(klipData), "- Your Highlight on page ")
	pageS := strings.Split(line, " |")[0]
	location := strings.TrimSpace(strings.Split(line, "|")[1])
	location = strings.TrimPrefix(location, "Location ")
	page, err := strconv.Atoi(pageS)
	if err != nil {
		return Klip{}, err
	}

	timeS := strings.TrimSpace(strings.Split(line, "day, ")[1])
	const dateFormat = "January 2, 2006 3:04:05 PM"
	t, err := time.Parse(dateFormat, timeS)
	if err != nil {
		return Klip{}, err
	}
	fmt.Println(t)

	klip := Klip{Page: page, Location: location, Time: t}

	return klip, nil
}
