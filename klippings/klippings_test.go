package klippings_test

import (
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/byarbrough/klippings_md/klippings"
)

func TestExtracKlips(t *testing.T) {

	klipfile := strings.NewReader(`Invent and Wander  
	- Your Highlight on page 37 | Location 616-618 | Added on Sunday, May 16, 2021 9:23:55 PM
		
	We will have to make many conscious and deliberate choices, some of which will be bold and unconventional. Hopefully, some will turn out to be winners. Certainly, some will turn out to be mistakes.
	==========`)

	// TODO: make standardiztion a function
	const dateFormat = "January 2, 2006 3:04:05 PM"
	entryTime, _ := time.Parse(dateFormat, "May 16, 2021 9:23:55 PM")

	want := klippings.Klip{Title: "Invent and Wander", Page: 37, Location: "616-618", Time: entryTime, Body: "We will have to make many conscious and deliberate choices, some of which will be bold and unconventional. Hopefully, some will turn out to be winners. Certainly, some will turn out to be mistakes."}
	got, err := klippings.ExtractKlips(klipfile)
	if err != nil {
		t.Fatal(err)
	}

	assertKlip(t, got, want)

}

// func TestNewKlip(t *testing.T) {
// 	/*
// 	   	ruskin := `Selected Writings (Oxford World's Classics) (Ruskin, John)
// 	   - Your Highlight on Location 1258-1261 | Added on Thursday, July 8, 2021 8:37:42 PM

// 	   God has lent us the earth for our life; it is a great entail. It belongs as much to those who are to come after us, and whose names are already written in the book of creation, as to us; and we have no right, by anything that we do or neglect, to involve them in unnecessary penalties, or deprive them of benefits which it was in our power to bequeath.
// 	   ==========`
// 	*/

// 	bezos := `Invent and Wander
// - Your Highlight on page 37 | Location 616-618 | Added on Sunday, May 16, 2021 9:23:55 PM

// We will have to make many conscious and deliberate choices, some of which will be bold and unconventional. Hopefully, some will turn out to be winners. Certainly, some will turn out to be mistakes.
// ==========`

// 	fs := fstest.MapFS{
// 		"one klip.txt": {Data: []byte(bezos)},
// 		//"second klip.txt": {Data: []byte(ruskin)},
// 		//"another klip.txt": {Data: []byte("Starship Troopers (Heinlein, Robert A.)\n- Your Highlight on page 98 | Location 1475-1477 | Added on Sunday, June 6, 2021 3:58:47 PM")},
// 	}

// 	klips, err := klippings.NewKlipFromFS(fs)

// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	// Proper number of files returned
// 	if len(klips) != len(fs) {
// 		t.Errorf("got %d klip files, want %d files", len(klips), len(fs))
// 	}

// 	// test fields
// 	const dateFormat = "January 2, 2006 3:04:05 PM"
// 	entryTime, _ := time.Parse(dateFormat, "May 16, 2021 9:23:55 PM")

// 	got := klips[0]
// 	want := klippings.Klip{Title: "Invent and Wander", Page: 37, Location: "616-618", Time: entryTime, Body: "We will have to make many conscious and deliberate choices, some of which will be bold and unconventional. Hopefully, some will turn out to be winners. Certainly, some will turn out to be mistakes."}

// 	assertKlip(t, got, want)

// }

func assertKlip(t *testing.T, got klippings.Klip, want klippings.Klip) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v want %+v", got, want)
	}
}
