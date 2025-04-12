package slices_test

import (
	"testing"

	"github.com/kurisu1024/go-boiler/slices"
	"github.com/stretchr/testify/assert"
)

func TestFilterInt(t *testing.T) {
	for _, tc := range []struct {
		name     string
		input    []int
		filterFn func(int) bool
		want     []int
	}{{
		name:     "filter_even_numbers",
		input:    []int{1, 2, 3, 4, 5, 6},
		filterFn: func(x int) bool { return x%2 == 0 },
		want:     []int{2, 4, 6},
	}, {
		name:     "filter_odd_numbers",
		input:    []int{1, 2, 3, 4, 5, 6},
		filterFn: func(x int) bool { return x%2 != 0 },
		want:     []int{1, 3, 5},
	}, {
		name:     "empty_input",
		input:    []int{},
		filterFn: func(x int) bool { return x%2 == 0 },
		want:     []int{},
	}, {
		name:     "No_matches",
		input:    []int{1, 3, 5},
		filterFn: func(x int) bool { return x%2 == 0 },
		want:     []int{},
	}} {
		t.Run(tc.name, func(t *testing.T) {
			got := slices.Filter(tc.input, tc.filterFn)
			assert.ElementsMatch(t, tc.want, got)
		})
	}

}

func TestFilterString(t *testing.T) {
	for _, tc := range []struct {
		name     string
		input    []string
		filterFn func(s string) bool
		want     []string
	}{{
		name:     "filter_exclude_zar",
		input:    []string{"foo", "bar", "zar"},
		filterFn: func(s string) bool { return s != "zar" },
		want:     []string{"foo", "bar"},
	}, {
		name:     "len_slice_grater than_3",
		input:    []string{"foo", "bar", "foo_long", "bar_long"},
		filterFn: func(s string) bool { return len(s) > 3 },
		want:     []string{"foo_long", "bar_long"},
	}, {
		name:     "len_slice_less than equal_3",
		input:    []string{"foo", "bar", "foo_long", "bar_long"},
		filterFn: func(s string) bool { return len(s) <= 3 },
		want:     []string{"foo", "bar"},
	}, {
		name:     "no_matches",
		input:    []string{"foo", "bar"},
		filterFn: func(s string) bool { return false },
		want:     []string{},
	}, {
		name:     "emtpy_input",
		input:    []string{},
		filterFn: func(s string) bool { return s != "zar" },
		want:     []string{},
	}} {
		t.Run(tc.name, func(t *testing.T) {
			got := slices.Filter(tc.input, tc.filterFn)
			assert.ElementsMatch(t, tc.want, got)
		})
	}
}
