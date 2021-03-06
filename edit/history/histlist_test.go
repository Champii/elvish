package history

import (
	"testing"

	"github.com/elves/elvish/edit/eddefs"
	"github.com/elves/elvish/edit/ui"
)

var (
	theHistList = newHistlist([]string{"ls", "echo lalala", "ls"})

	histlistDedupFilterTests = []eddefs.ListingProviderFilterTest{
		{"", []eddefs.ListingShown{
			{"1", ui.Unstyled("echo lalala")},
			{"2", ui.Unstyled("ls")}}},
		{"l", []eddefs.ListingShown{
			{"1", ui.Unstyled("echo lalala")},
			{"2", ui.Unstyled("ls")}}},
	}

	histlistNoDedupFilterTests = []eddefs.ListingProviderFilterTest{
		{"", []eddefs.ListingShown{
			{"0", ui.Unstyled("ls")},
			{"1", ui.Unstyled("echo lalala")},
			{"2", ui.Unstyled("ls")}}},
		{"l", []eddefs.ListingShown{
			{"0", ui.Unstyled("ls")},
			{"1", ui.Unstyled("echo lalala")},
			{"2", ui.Unstyled("ls")}}},
	}
)

func TestHistlist(t *testing.T) {
	eddefs.TestListingProviderFilter(t, "theHistList", theHistList, histlistDedupFilterTests)
	theHistList.dedup = false
	eddefs.TestListingProviderFilter(t, "theHistList", theHistList, histlistNoDedupFilterTests)
}
