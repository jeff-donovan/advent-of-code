package main

import "testing"

func TestIsImpossiblePath(t *testing.T) {
	tests := []struct {
		name    string
		req     JoltageRequirement
		current JoltageRequirement
		want    bool
	}{
		{
			name:    "impossible - at least one value in current is greater than corresponding value in req",
			req:     JoltageRequirement{0},
			current: JoltageRequirement{1},
			want:    true,
		},

		{
			name:    "possible - all values in current are less than or equal to corresponding value in req",
			req:     JoltageRequirement{0},
			current: JoltageRequirement{0},
			want:    false,
		},

		{
			name:    "possible - JoltageRequirement of length > 1 - all values in current are less than or equal to corresponding value in req",
			req:     JoltageRequirement{0, 1},
			current: JoltageRequirement{0, 0},
			want:    false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := isImpossiblePath(tc.req, tc.current)
			if got != tc.want {
				t.Fatalf("isImpossiblePath(%v, %v) = %v, want %v", tc.req, tc.current, got, tc.want)
			}
		})
	}
}

func TestAreRequirementsEqual(t *testing.T) {
	tests := []struct {
		name string
		a    JoltageRequirement
		b    JoltageRequirement
		want bool
	}{
		{
			name: "empty reqs are equal",
			a:    nil,
			b:    nil,
			want: true,
		},

		{
			name: "equal reqs are equal",
			a:    JoltageRequirement{0, 1, 2},
			b:    JoltageRequirement{0, 1, 2},
			want: true,
		},

		{
			name: "reqs with diff values are NOT equal",
			a:    JoltageRequirement{0, 2, 1},
			b:    JoltageRequirement{0, 1, 2},
			want: false,
		},

		{
			name: "reqs with diff lengths are NOT equal",
			a:    JoltageRequirement{0, 1},
			b:    JoltageRequirement{0, 1, 2},
			want: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := areRequirementsEqual(tc.a, tc.b)
			if got != tc.want {
				t.Fatalf("areRequirementsEqual(%v, %v) = %v, want %v", tc.a, tc.b, got, tc.want)
			}
		})
	}
}
