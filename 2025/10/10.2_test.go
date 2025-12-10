package main

import "testing"

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
