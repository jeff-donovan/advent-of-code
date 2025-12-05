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
			name: "simple overlap",
			a:    Range{1, 1},
			b:    Range{1, 1},
			want: true,
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
