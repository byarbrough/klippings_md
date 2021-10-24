package klippings_test

import (
	"reflect"
	"testing"
	"testing/fstest"
	"time"

	"github.com/byarbrough/klippings_md/klippings"
)

func TestNewKlip(t *testing.T) {
	fs := fstest.MapFS{
		"one klip.txt":    {Data: []byte("- Your Highlight on page 37 | Location 616-618 | Added on Sunday, May 16, 2021 9:23:55 PM")},
		"second klip.txt": {Data: []byte("- Your Highlight on page 194 | Location 2797-2797 | Added on Monday, June 7, 2021 11:05:04 PM")},
	}

	klips, err := klippings.NewKlipFromFS(fs)

	if err != nil {
		t.Fatal(err)
	}

	// Proper number of files returned
	if len(klips) != len(fs) {
		t.Errorf("got %d klip files, want %d files", len(klips), len(fs))
	}

	// test fields
	const dateFormat = "January 2, 2006 3:04:05 PM"
	entryTime, _ := time.Parse(dateFormat, "May 16, 2021 9:23:55 PM")

	got := klips[0]
	want := klippings.Klip{Page: 37, Location: "616-618", Time: entryTime}

	assertKlip(t, got, want)

}

func assertKlip(t *testing.T, got klippings.Klip, want klippings.Klip) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v want %+v", got, want)
	}
}
