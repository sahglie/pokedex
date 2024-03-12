package repo

import (
	"testing"
)

type PageState struct {
	pagesBefore [2]int
	pagesAfter  [2]int
	fn          func() ([]string, error)
	err         error
}

func TestRepo_Pager(t *testing.T) {
	r := NewRepo()

	states := []PageState{
		{pagesBefore: [2]int{1, -1}, pagesAfter: [2]int{1, -1}, fn: r.LocationsPrev, err: ErrNoPrevPage},
		{pagesBefore: [2]int{1, -1}, pagesAfter: [2]int{2, 0}, fn: r.LocationsNext, err: nil},
		{pagesBefore: [2]int{2, 0}, pagesAfter: [2]int{2, 0}, fn: r.LocationsPrev, err: ErrNoPrevPage},
		{pagesBefore: [2]int{2, 0}, pagesAfter: [2]int{3, 1}, fn: r.LocationsNext, err: nil},
		{pagesBefore: [2]int{3, 1}, pagesAfter: [2]int{2, 0}, fn: r.LocationsPrev, err: nil},
	}

	for _, s := range states {
		if s.pagesBefore != r.pages() {
			t.Errorf("expected %v got %v", s.pagesBefore, r.pages())
		}

		s.fn()

		if s.pagesAfter != r.pages() {
			t.Errorf("expected %v got %v", s.pagesAfter, r.pages())
		}
	}
}

func TestRepo_CachesLocationResults(t *testing.T) {
	r := NewRepo()

	if r.cache.NumKeys() != 0 {
		t.Errorf("expected %d got %d\v", 0, r.cache.NumKeys())
	}

	r.LocationsNext()

	if r.cache.NumKeys() != 1 {
		t.Errorf("expected %d got %d\v", 1, r.cache.NumKeys())
	}
}
