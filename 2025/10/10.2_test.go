package main

import (
	"math"
	"testing"
)

func TestCalculateFewestButtonClicksRemaining(t *testing.T) {
	tests := []struct {
		name    string
		machine Machine
		current JoltageRequirement
		want    int
	}{
		{
			name:    "0 when Machine requirements match current",
			machine: Machine{nil, nil, JoltageRequirement{0, 1, 2}},
			current: JoltageRequirement{0, 1, 2},
			want:    0,
		},

		{
			name:    "math.MaxInt when we've gone past machine requirements",
			machine: Machine{nil, nil, JoltageRequirement{0}},
			current: JoltageRequirement{1},
			want:    math.MaxInt,
		},

		{
			name:    "1 when 1 button click will produce machine requirements",
			machine: Machine{nil, []Button{{0}}, JoltageRequirement{1}},
			current: JoltageRequirement{0},
			want:    1,
		},

		{
			name:    "1 when second button click will produce machine requirements",
			machine: Machine{nil, []Button{{0}, {1}}, JoltageRequirement{0, 1}},
			current: JoltageRequirement{0, 0},
			want:    1,
		},

		{
			name:    "test example",
			machine: Machine{nil, []Button{{3}, {1, 3}, {2}, {2, 3}, {0, 2}, {0, 1}}, JoltageRequirement{3, 5, 4, 7}},
			current: JoltageRequirement{0, 0, 0, 0},
			want:    10,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := calculateFewestButtonClicksRemaining(tc.machine, tc.current)
			if got != tc.want {
				t.Fatalf("calculateFewestButtonClicksRemaining(%v, %v) = %v, want %v", tc.machine, tc.current, got, tc.want)
			}
		})
	}
}

func TestMakeNextJoltageRequirement(t *testing.T) {
	tests := []struct {
		name    string
		current JoltageRequirement
		click   Button
		want    JoltageRequirement
	}{
		{
			name:    "click an empty button produces same JoltageRequirement",
			current: JoltageRequirement{0, 1, 2, 3},
			click:   nil,
			want:    JoltageRequirement{0, 1, 2, 3},
		},

		{
			name:    "click a button with 1 element produces new JoltageRequirement",
			current: JoltageRequirement{0, 1, 2, 3},
			click:   Button{0},
			want:    JoltageRequirement{1, 1, 2, 3},
		},

		{
			name:    "click a button with multiple elements produces new JoltageRequirement",
			current: JoltageRequirement{0, 1, 2, 3},
			click:   Button{0, 3},
			want:    JoltageRequirement{1, 1, 2, 4},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := makeNextJoltageRequirement(tc.current, tc.click)
			if !areRequirementsEqual(got, tc.want) {
				t.Fatalf("makeNextJoltageRequirement(%v, %v) = %v, want %v", tc.current, tc.click, got, tc.want)
			}
		})
	}
}

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
