package klippings

import (
	"bufio"
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

	scanner := bufio.NewScanner(klipFile)

	readLine := func() string {
		scanner.Scan()
		return scanner.Text()
	}

	// titleLine := readLine()

	klip, err := parseMetadataLine(readLine())
	if err != nil {
		return Klip{}, err
	}

	return klip, nil
}

func parseMetadataLine(line string) (Klip, error) {
	line = strings.TrimPrefix(line, "- Your Highlight on page ")
	pageS := strings.Split(line, " |")[0]
	page, err := strconv.Atoi(pageS)
	if err != nil {
		return Klip{}, err
	}

	location := strings.TrimSpace(strings.Split(line, "|")[1])
	location = strings.TrimPrefix(location, "Location ")

	const dateFormat = "January 2, 2006 3:04:05 PM"
	timeS := strings.TrimSpace(strings.Split(line, "day, ")[1])
	t, err := time.Parse(dateFormat, timeS)
	if err != nil {
		return Klip{}, err
	}

	return Klip{Page: page, Location: location, Time: t}, nil
}
