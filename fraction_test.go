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
		{-2, 4, "-1/2"},
		{2, -4, "-1/2"},
		{-2, -4, "1/2"},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%d/%d", tt.divisor, tt.dividend), func(t *testing.T) {
			fraction, _ := New(tt.divisor, tt.dividend)

			fras := fraction.String()
			if fras != tt.want {
				t.Fatalf("%d/%d should be equal %s but %s", tt.divisor, tt.dividend, tt.want, fras)
			}
		})
	}
}

func TestFraction_Add(t *testing.T) {
	tests := [][]int64{
		{1, 2, 3, 4, 5, 4},
		{3, 7, 4, 7, 1, 1},
		{3, 7, 4, 9, 55, 63},
		{1, 2, -2, 4, 0, 1},
		{1, 2, 2, -4, 0, 2},
	}
	for _, arr := range tests {
		t.Run(fmt.Sprint(arr), func(t *testing.T) {
			fraction, _ := New(arr[0], arr[1])
			other, _ := New(arr[2], arr[3])
			result, _ := New(arr[4], arr[5])
			if got := fraction.Add(other); !got.Equal(result) {
				t.Errorf("Fraction.Add() = %v, want %v", got, result)
			}
		})
	}
}

func BenchmarkFraction_Add(b *testing.B) {
	tests := [][]int64{
		{1, 2, 3, 4, 5, 4},
		{3, 7, 4, 7, 1, 1},
		{3, 7, 4, 9, 55, 63},
		{1, 2, -2, 4, 0, 1},
		{1, 2, 2, -4, 0, 2},
	}

	for i := 0; i < b.N; i++ {
		for _, arr := range tests {
			fraction, _ := New(arr[0], arr[1])
			other, _ := New(arr[2], arr[3])
			result, _ := New(arr[4], arr[5])
			if got := fraction.Add(other); !got.Equal(result) {
				b.Errorf("Fraction.Add() = %v, want %v", got, result)
			}
		}
	}
}

func TestFraction_Subtract(t *testing.T) {
	tests := [][]int64{
		{1, 2, 3, 4, -1, 4},
		{3, 7, 4, 7, -1, 7},
		{4, 9, 3, 7, 1, 63},
		{1, 2, -2, 4, 1, 1},
		{1, 2, 2, -4, 2, 2},
	}
	for _, arr := range tests {
		t.Run(fmt.Sprint(arr), func(t *testing.T) {
			fraction, _ := New(arr[0], arr[1])
			other, _ := New(arr[2], arr[3])
			result, _ := New(arr[4], arr[5])
			if got := fraction.Subtract(other); !got.Equal(result) {
				t.Errorf("Fraction.Subtract() = %v, want %v", got, result)
			}
		})
	}
}

func TestFraction_Multiply(t *testing.T) {
	tests := [][]int64{
		{1, 2, 3, 4, 3, 8},
		{3, 7, 4, 7, 12, 49},
		{4, 9, 3, 7, 4, 21},
		{1, 2, -2, 4, -1, 4},
		{1, 2, 2, -4, -1, 4},
	}
	for _, arr := range tests {
		t.Run(fmt.Sprint(arr), func(t *testing.T) {
			fraction, _ := New(arr[0], arr[1])
			other, _ := New(arr[2], arr[3])
			result, _ := New(arr[4], arr[5])
			if got := fraction.Multiply(other); !got.Equal(result) {
				t.Errorf("Fraction.Multiply() = %v, want %v", got, result)
			}
		})
	}
}

func Test_findMaxCommonDividend(t *testing.T) {
	tests := []struct {
		i    int64
		j    int64
		want int64
	}{
		{1, 2, 1},
		{2, 4, 2},
		{11, 33, 11},
		{-2, 6, 2},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprint(tt.i, tt.j, tt.want), func(t *testing.T) {
			if got := findMaxCommonDividend(tt.i, tt.j); got != tt.want {
				t.Errorf("findMaxCommonDividend() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkFindMaxCommonDividend(b *testing.B) {
	tests := []struct {
		i    int64
		j    int64
		want int64
	}{
		{1, 2, 1},
		{2, 4, 2},
		{11, 33, 11},
		{-2, 6, 2},
	}

	for i := 0; i < b.N; i++ {
		for _, tt := range tests {
			if got := findMaxCommonDividend(tt.i, tt.j); got != tt.want {
				b.Errorf("findMaxCommonDividend() = %v, want %v", got, tt.want)
			}
		}
	}
}

func TestFraction_Divide(t *testing.T) {
	tests := [][]int64{
		{1, 2, 3, 4, 2, 3},
		{3, 7, 4, 7, 3, 4},
		{4, 9, 3, 7, 28, 27},
		{1, 2, -2, 4, -1, 1},
		{1, 2, 2, -4, -1, 1},
	}
	for _, arr := range tests {
		t.Run(fmt.Sprint(arr), func(t *testing.T) {
			fraction, _ := New(arr[0], arr[1])
			other, _ := New(arr[2], arr[3])
			result, _ := New(arr[4], arr[5])
			if got, _ := fraction.Divide(other); !got.Equal(result) {
				t.Errorf("Fraction.Multiply() = %v, want %v", got, result)
			}
		})
	}
}

func TestFraction_Greater(t *testing.T) {
	tests := []struct {
		args []int64
		want bool
	}{
		{[]int64{2, 4, 1, 4}, true},
		{[]int64{3, 7, 2, 5}, true},
		{[]int64{-2, 7, -2, 5}, true},
		{[]int64{2, 7, -2, 5}, true},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprint(tt.args), func(t *testing.T) {
			fraction, _ := New(tt.args[0], tt.args[1])
			otherFraction, _ := New(tt.args[2], tt.args[3])

			if got := fraction.Greater(otherFraction); got != tt.want {
				t.Errorf("Fraction.Greater() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFraction_Int(t *testing.T) {
	tests := []struct {
		args []int64
		want int
	}{
		{[]int64{1, 2}, 0},
		{[]int64{2, 2}, 1},
		{[]int64{3, 2}, 1},
		{[]int64{4, 3}, 1},
		{[]int64{5, 3}, 1},
		{[]int64{-5, 3}, -1},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprint(tt.args), func(t *testing.T) {
			fraction, _ := New(tt.args[0], tt.args[1])
			if got := fraction.Int(); got != tt.want {
				t.Errorf("Fraction.Int() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFraction_Int64(t *testing.T) {
	tests := []struct {
		args []int64
		want int64
	}{
		{[]int64{1, 2}, 0},
		{[]int64{2, 2}, 1},
		{[]int64{3, 2}, 1},
		{[]int64{4, 3}, 1},
		{[]int64{5, 3}, 1},
		{[]int64{-5, 3}, -1},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprint(tt.args), func(t *testing.T) {
			fraction, _ := New(tt.args[0], tt.args[1])
			if got := fraction.Int64(); got != tt.want {
				t.Errorf("Fraction.Int() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFraction_Float32(t *testing.T) {
	tests := [][]int64{
		{1, 2},
		{2, 2},
		{3, 2},
		{4, 3},
		{5, 3},
		{-5, 3},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprint(tt), func(t *testing.T) {
			fraction, _ := New(tt[0], tt[1])
			if got := fraction.Float32(); got != float32(tt[0])/float32(tt[1]) {
				t.Errorf("Fraction.Float32() = %v", got)
			}
		})
	}
}

func TestFraction_Float64(t *testing.T) {
	tests := [][]int64{
		{1, 2},
		{2, 2},
		{3, 2},
		{4, 3},
		{5, 3},
		{-5, 3},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprint(tt), func(t *testing.T) {
			fraction, _ := New(tt[0], tt[1])
			if got := fraction.Float64(); got != float64(tt[0])/float64(tt[1]) {
				t.Errorf("Fraction.Float64() = %v", got)
			}
		})
	}
}

func TestFraction_Positive(t *testing.T) {
	tests := []struct {
		args []int64
		want bool
	}{
		{[]int64{1, 2}, true},
		{[]int64{-1, 2}, false},
		{[]int64{-1, -2}, true},
		{[]int64{0, -2}, true},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprint(tt.args), func(t *testing.T) {
			fraction, _ := New(tt.args[0], tt.args[1])
			if got := fraction.Positive(); got != tt.want {
				t.Errorf("Fraction.Positive() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFraction_Equal(t *testing.T) {
	tests := []struct {
		args []int64
		want bool
	}{
		{[]int64{2, 4, 1, 2}, true},
		{[]int64{3, 7, 2, 5}, false},
		{[]int64{-2, 7, -2, 5}, false},
		{[]int64{2, 7, -2, 5}, false},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprint(tt.args), func(t *testing.T) {
			fraction, _ := New(tt.args[0], tt.args[1])
			otherFraction, _ := New(tt.args[2], tt.args[3])

			if got := fraction.Equal(otherFraction); got != tt.want {
				t.Errorf("Fraction.Equal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFraction_IsZero(t *testing.T) {
	fraction, _ := New(0, 100)

	if !fraction.IsZero() {
		t.Fatal("it should be zero")
	}
}
