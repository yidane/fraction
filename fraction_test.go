package fraction

import (
	"fmt"
	"testing"
)

func TestNew(t *testing.T) {
	tests := []struct {
		divisor  int64
		dividend int64
		want     string
	}{
		{18, 20, "9/10"},
		{18, 19, "18/19"},
		{18, 18, "1"},
		{18, 17, "18/17"},
		{18, 16, "9/8"},
		{18, 15, "6/5"},
		{18, 14, "9/7"},
		{18, 13, "18/13"},
		{18, 12, "3/2"},
		{18, 11, "18/11"},
		{18, 10, "9/5"},
		{18, 9, "2"},
		{18, 8, "9/4"},
		{18, 7, "18/7"},
		{18, 6, "3"},
		{18, 5, "18/5"},
		{18, 4, "9/2"},
		{18, 3, "6"},
		{18, 2, "9"},
		{18, 1, "18"},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%d/%d", tt.divisor, tt.dividend), func(t *testing.T) {
			fraction := New(tt.divisor, tt.dividend)

			fmt.Println(fraction)

			if fraction.String() != tt.want {
				t.Fatalf("%d/%d should be equal %s", tt.divisor, tt.dividend, tt.want)
			}
		})
	}
}

func TestFraction_Add(t *testing.T) {
	tests := [][]int64{
		{1, 2, 3, 4, 5, 4},
		{3, 7, 4, 7, 1, 1},
		{3, 7, 4, 9, 55, 63},
	}
	for _, arr := range tests {
		t.Run(fmt.Sprint(arr), func(t *testing.T) {
			fraction := New(arr[0], arr[1])
			other := New(arr[2], arr[3])
			result := New(arr[4], arr[5])
			if got := fraction.Add(other); !got.Equal(result) {
				t.Errorf("Fraction.Add() = %v, want %v", got, result)
			}
		})
	}
}
