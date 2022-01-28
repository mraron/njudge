package pagination

import (
	"net/url"
	"testing"
)

func TestLinks(t *testing.T) {
	qu := url.Values{}
	qu.Set("njudge", "yes")

	removeErr := func(l []Link, err error) []Link {
		if err != nil {
			t.Error(err)
		}
		return l
	}

	var tests = []struct {
		got []Link
		expected []Link
	}{
		{
			got: removeErr(Links(2, 1, 3, qu)),
			expected: []Link{
				{Name: "&laquo;", Active: false, Disabled: false, Url: "?njudge=yes&page=1"},
				{Name: "1", Active: false, Disabled: false, Url: "?njudge=yes&page=1"},
				{Name: "2", Active: true, Disabled: true, Url: "?njudge=yes&page=2"},
				{Name: "3", Active: false, Disabled: false, Url: "?njudge=yes&page=3"},
				{Name: "&raquo;", Active: false, Disabled: false, Url: "?njudge=yes&page=3"},
			},
		},
		{
			got: removeErr(Links(1, 1, 3, qu)),
			expected: []Link{
				{Name: "&laquo;", Active: false, Disabled: true, Url: "#"},
				{Name: "1", Active: true, Disabled: true, Url: "?njudge=yes&page=1"},
				{Name: "2", Active: false, Disabled: false, Url: "?njudge=yes&page=2"},
				{Name: "3", Active: false, Disabled: false, Url: "?njudge=yes&page=3"},
				{Name: "&raquo;", Active: false, Disabled: false, Url: "?njudge=yes&page=2"},
			},
		},
		{
			got: removeErr(LinksWithCountLimit(50, 1, 100, qu, 2)),
			expected: []Link{
				{Name: "&laquo;", Active: false, Disabled: false, Url: "?njudge=yes&page=49"},
				{Name: "1", Active: false, Disabled: false, Url: "?njudge=yes&page=1"},
				{Name: "...", Active: false, Disabled: true, Url: "#"},
				{Name: "48", Active: false, Disabled: false, Url: "?njudge=yes&page=48"},
				{Name: "49", Active: false, Disabled: false, Url: "?njudge=yes&page=49"},
				{Name: "50", Active: true, Disabled: true, Url: "?njudge=yes&page=50"},
				{Name: "51", Active: false, Disabled: false, Url: "?njudge=yes&page=51"},
				{Name: "52", Active: false, Disabled: false, Url: "?njudge=yes&page=52"},
				{Name: "...", Active: false, Disabled: true, Url: "#"},
				{Name: "100", Active: false, Disabled: false, Url: "?njudge=yes&page=100"},
				{Name: "&raquo;", Active: false, Disabled: false, Url: "?njudge=yes&page=51"},
			},
		},
		{
			got: removeErr(LinksWithCountLimit(3, 1, 100, qu, 2)),
			expected: []Link{
				{Name: "&laquo;", Active: false, Disabled: false, Url: "?njudge=yes&page=2"},
				{Name: "1", Active: false, Disabled: false, Url: "?njudge=yes&page=1"},
				{Name: "2", Active: false, Disabled: false, Url: "?njudge=yes&page=2"},
				{Name: "3", Active: true, Disabled: true, Url: "?njudge=yes&page=3"},
				{Name: "4", Active: false, Disabled: false, Url: "?njudge=yes&page=4"},
				{Name: "5", Active: false, Disabled: false, Url: "?njudge=yes&page=5"},
				{Name: "...", Active: false, Disabled: true, Url: "#"},
				{Name: "100", Active: false, Disabled: false, Url: "?njudge=yes&page=100"},
				{Name: "&raquo;", Active: false, Disabled: false, Url: "?njudge=yes&page=4"},
			},
		},
		{
			got: removeErr(LinksWithCountLimit(3, 1, 100, qu, 1)),
			expected: []Link{
				{Name: "&laquo;", Active: false, Disabled: false, Url: "?njudge=yes&page=2"},
				{Name: "1", Active: false, Disabled: false, Url: "?njudge=yes&page=1"},
				{Name: "2", Active: false, Disabled: false, Url: "?njudge=yes&page=2"},
				{Name: "3", Active: true, Disabled: true, Url: "?njudge=yes&page=3"},
				{Name: "4", Active: false, Disabled: false, Url: "?njudge=yes&page=4"},
				{Name: "...", Active: false, Disabled: true, Url: "#"},
				{Name: "100", Active: false, Disabled: false, Url: "?njudge=yes&page=100"},
				{Name: "&raquo;", Active: false, Disabled: false, Url: "?njudge=yes&page=4"},
			},
		},
	}

	for idx, test := range tests {
		if len(test.expected) != len(test.got) {
			t.Error(idx, ": wrong length")
		}

		for i := 0; i < len(test.got); i++ {
			if test.got[i] != test.expected[i] {
				t.Error(idx, "(",i,")", ":", test.got[i], "!=", test.expected[i])
			}
		}
	}
}

