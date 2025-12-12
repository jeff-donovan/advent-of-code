package main

import "testing"

func presentsEqual(a, b Present) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func TestFlipHorizontal(t *testing.T) {
	tests := []struct {
		name string
		a    Present
		want Present
	}{
		{
			name: "flipping horizontal",
			a:    []string{"###", "##.", "##."},
			want: []string{"###", ".##", ".##"},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := flipHorizontal(tc.a)
			if !presentsEqual(got, tc.want) {
				t.Fatalf("flipHorizontal(%v) = %v, want %v", tc.a, got, tc.want)
			}
		})
	}
}

func TestReverseString(t *testing.T) {
	tests := []struct {
		name string
		a    string
		want string
	}{
		{
			name: "reverse empty string produces empty string",
			a:    "",
			want: "",
		},

		{
			name: "reverse jeff",
			a:    "jeff",
			want: "ffej",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := reverseString(tc.a)
			if got != tc.want {
				t.Fatalf("reverseString(%v) = %v, want %v", tc.a, got, tc.want)
			}
		})
	}
}

func TestFlipVertical(t *testing.T) {
	tests := []struct {
		name string
		a    Present
		want Present
	}{
		{
			name: "flipping vertical",
			a:    []string{"###", "##.", "##."},
			want: []string{"##.", "##.", "###"},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := flipVertical(tc.a)
			if !presentsEqual(got, tc.want) {
				t.Fatalf("flipVertical(%v) = %v, want %v", tc.a, got, tc.want)
			}
		})
	}
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

		{
			name: "rotating 90 degrees 1 time",
			a:    []string{"###", "##.", "##."},
			n:    1,
			want: []string{"###", "###", "..#"},
		},

		{
			name: "rotating 90 degrees 2 times",
			a:    []string{"###", "##.", "##."},
			n:    2,
			want: []string{".##", ".##", "###"},
		},

		{
			name: "rotating 90 degrees 3 times",
			a:    []string{"###", "##.", "##."},
			n:    3,
			want: []string{"#..", "###", "###"},
		},

		{
			name: "rotating 90 degrees 4 times produces the same Present",
			a:    []string{"###", "##.", "##."},
			n:    4,
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
