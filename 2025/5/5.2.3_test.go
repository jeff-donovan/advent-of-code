package main

import "testing"

func TestIsOverlapping3(t *testing.T) {
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
			got := isOverlapping3(tc.a, tc.b)
			if got != tc.want {
				t.Fatalf("isOverlapping(%v, %v) = %v, want %v", tc.a, tc.b, got, tc.want)
			}
		})
	}
}
