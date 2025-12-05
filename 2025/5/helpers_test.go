package main

import "testing"

func rangeSlicesEqual(a, b []Range) bool {
	// same length
	// all elements in a are in b
	// all elements in b are in a
	if len(a) != len(b) {
		return false
	}

	for _, ra := range a {
		seen := false
		for _, rb := range b {
			if rangeEqual(ra, rb) {
				seen = true
			}
		}
		if !seen {
			return false
		}
	}

	for _, rb := range b {
		seen := false
		for _, ra := range a {
			if rangeEqual(rb, ra) {
				seen = true
			}
		}
		if !seen {
			return false
		}
	}

	return true
}

func rangeEqual(a, b Range) bool {
	return a.start == b.start && a.end == b.end
}

func TestDedupe(t *testing.T) {
	tests := []struct {
		name string
		a    []Range
		want []Range
	}{
		{
			name: "empty returns empty",
			a:    []Range{},
			want: []Range{},
		},
		{
			name: "empty returns nil",
			a:    []Range{},
			want: nil,
		},
		{
			name: "nil returns empty",
			a:    nil,
			want: []Range{},
		},
		{
			name: "nil returns nil",
			a:    nil,
			want: nil,
		},
		{
			name: "two duplicate ranges reduced to 1",
			a:    []Range{{1, 1}, {1, 1}},
			want: []Range{{1, 1}},
		},
		{
			name: "three duplicate ranges reduced to 1",
			a:    []Range{{1, 1}, {1, 1}, {1, 1}},
			want: []Range{{1, 1}},
		},
		{
			name: "multiple sets of dupes reduced to just theirs, order doesn't matter",
			a:    []Range{{1, 2}, {2, 1}, {1, 2}, {2, 1}},
			want: []Range{{2, 1}, {1, 2}},
		},
		{
			name: "unduplicated sets are not removed",
			a:    []Range{{1, 2}, {3, 4}, {1, 2}},
			want: []Range{{3, 4}, {1, 2}},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := dedupe(tc.a)
			if !rangeSlicesEqual(got, tc.want) {
				t.Fatalf("dedupe(%v) = %v, want %v", tc.a, got, tc.want)
			}
		})
	}
}

func TestIsOverlapping(t *testing.T) {
	tests := []struct {
		name string
		a    Range
		b    Range
		want bool
	}{
		{
			name: "same ranges overlap",
			a:    Range{1, 1},
			b:    Range{1, 1},
			want: true,
		},
		{
			name: "inner range and outer range overlap",
			a:    Range{1, 1},
			b:    Range{0, 2},
			want: true,
		},
		{
			name: "inner range and outer range overlap - flipped",
			a:    Range{0, 2},
			b:    Range{1, 1},
			want: true,
		},
		{
			name: "",
			a:    Range{1, 3},
			b:    Range{2, 4},
			want: true,
		},
		{
			name: "",
			a:    Range{0, 1},
			b:    Range{2, 3},
			want: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := isOverlapping(tc.a, tc.b)
			if got != tc.want {
				t.Fatalf("isOverlapping(%v, %v) = %v, want %v", tc.a, tc.b, got, tc.want)
			}
		})
	}
}
