package klippings

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"strconv"
	"strings"
	"time"
)

type Klip struct {
	Title    string
	Page     int
	Location string
	Time     time.Time
	Body     string
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
	return ExtractKlips(klipFile)
}

func ExtractKlips(klipFile io.Reader) (Klip, error) {

	scanner := bufio.NewScanner(klipFile)

	readLine := func() string {
		scanner.Scan()
		return strings.TrimSpace(scanner.Text())
	}

	klip := Klip{Title: readLine()}

	klip, err := parseMetadataLine(readLine(), klip)
	if err != nil {
		return Klip{}, err
	}

	readLine() // skip a line

	klip.Body = readLine()

	// this should be the end of the klip
	if readLine() != "==========" {
		return klip, errors.New("improperly formatted highlight")
	}

	return klip, nil
}

// parseMetadataLine handles the second line of a highlight
func parseMetadataLine(line string, klip Klip) (Klip, error) {
	line = strings.TrimPrefix(line, "- Your Highlight on page ")
	pageS := strings.Split(line, " |")[0]
	page, err := strconv.Atoi(pageS)
	if err != nil {
		return Klip{}, fmt.Errorf("unable to parse page number %w", err)
	}

	location := strings.Split(line, " | ")[1]
	location = strings.TrimPrefix(location, "Location ")

	const dateFormat = "January 2, 2006 3:04:05 PM"
	timeS := strings.Split(line, "day, ")[1]
	t, err := time.Parse(dateFormat, timeS)
	if err != nil {
		return Klip{}, fmt.Errorf("unable to parse location %w", err)
	}

	klip.Page = page
	klip.Location = location
	klip.Time = t
	return klip, nil
}
