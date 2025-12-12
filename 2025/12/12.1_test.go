package main

import "testing"

func presentsEqual(a, b Present) bool {
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func TestRotate90DegreesNTimes(t *testing.T) {
	tests := []struct {
		name string
		a    Present
		n    int
		want Present
	}{
		{
			name: "rotating 90 degrees 0 times produces the same Present",
			a:    []string{"###", "##.", "##."},
			n:    0,
			want: []string{"###", "##.", "##."},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := rotate90DegreesNTimes(tc.a, tc.n)
			if !presentsEqual(got, tc.want) {
				t.Fatalf("rotate90DegreesNTimes(%v, %d) = %v, want %v", tc.a, tc.n, got, tc.want)
			}
		})
	}
}
