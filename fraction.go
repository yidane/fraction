package fraction

import (
	"errors"
	"fmt"
	"strconv"
)

var errDivideByZero = errors.New("fraction: attempt to divide by zero")

type Fraction struct {
	divisor  int64
	dividend int64
}

func New(divisor, dividend int64) (Fraction, error) {
	if dividend == 0 {
		return Fraction{}, errDivideByZero
	}

	if divisor == 0 {
		return Fraction{divisor: 0, dividend: 1}, nil
	}

	return newFraction(divisor, dividend), nil
}

func newFraction(divisor, dividend int64) Fraction {
	if dividend < 0 {
		divisor = -divisor
		dividend = -dividend
	}

	positive := true
	if divisor < 0 {
		positive = false
		divisor = -divisor
	}

	maxDivisor := findMaxCommonDividend(divisor, dividend)

	divisor = divisor / maxDivisor
	dividend = dividend / maxDivisor

	if !positive {
		divisor = -divisor
	}

	return Fraction{divisor: divisor, dividend: dividend}
}

func findMaxCommonDividend(i, j int64) int64 {
	if i < 0 {
		i = -i
	}

	if j < 0 {
		j = -j
	}

	if i < j {
		i, j = j, i
	}

	for j != 0 {
		i, j = j, i%j
	}
	return i
}

func (fraction Fraction) Add(other Fraction) Fraction {
	maxDivisor := findMaxCommonDividend(fraction.dividend, other.dividend)
	fTimes, oTimes := fraction.dividend/maxDivisor, other.dividend/maxDivisor

	dividend := fTimes * oTimes * maxDivisor
	divisor := fraction.divisor*oTimes + other.divisor*fTimes

	return newFraction(divisor, dividend)
}

func (fraction Fraction) Subtract(other Fraction) Fraction {
	if fraction.Equal(other) {
		return newFraction(0, 1)
	}

	maxDivisor := findMaxCommonDividend(fraction.dividend, other.dividend)
	fTimes, oTimes := fraction.dividend/maxDivisor, other.dividend/maxDivisor

	dividend := fTimes * oTimes * maxDivisor
	divisor := fraction.divisor*oTimes - other.divisor*fTimes

	return newFraction(divisor, dividend)
}

func (fraction *Fraction) Multiply(other Fraction) Fraction {
	maxDivisor1 := findMaxCommonDividend(fraction.divisor, other.dividend)
	maxDivisor2 := findMaxCommonDividend(fraction.dividend, other.divisor)

	divisor := (fraction.divisor / maxDivisor1) * (other.divisor / maxDivisor2)
	dividend := (fraction.dividend / maxDivisor2) * (other.dividend / maxDivisor1)

	return newFraction(divisor, dividend)
}

func (fraction Fraction) Divide(other Fraction) (Fraction, error) {
	if other.IsZero() {
		return Fraction{}, errDivideByZero
	}

	maxDivisor1 := findMaxCommonDividend(fraction.divisor, other.divisor)
	maxDivisor2 := findMaxCommonDividend(fraction.dividend, other.dividend)

	divisor := (fraction.divisor / maxDivisor1) * (other.dividend / maxDivisor2)
	dividend := (fraction.dividend / maxDivisor2) * (other.divisor / maxDivisor1)

	return newFraction(divisor, dividend), nil
}

func (fraction Fraction) Greater(other Fraction) bool {
	maxDivisor := findMaxCommonDividend(fraction.dividend, other.dividend)
	return fraction.divisor*(other.dividend/maxDivisor) > fraction.dividend*(other.divisor/maxDivisor)
}

func (fraction Fraction) String() string {
	if fraction.dividend == 1 {
		return strconv.FormatInt(fraction.divisor, 10)
	}

	return fmt.Sprintf("%d/%d", fraction.divisor, fraction.dividend)
}

func (fraction Fraction) Int() int {
	r := fraction.divisor / fraction.dividend
	return int(r)
}

func (fraction Fraction) Int64() int64 {
	return fraction.divisor / fraction.dividend
}

func (fraction Fraction) Float32() float32 {
	return float32(fraction.divisor) / float32(fraction.dividend)
}

func (fraction Fraction) Float64() float64 {
	return float64(fraction.divisor) / float64(fraction.dividend)
}

func (fraction Fraction) Positive() bool {
	return fraction.divisor >= 0
}

func (fraction Fraction) Equal(other Fraction) bool {
	if fraction.divisor != other.divisor {
		return false
	}

	return fraction.dividend == other.dividend
}

func (fraction Fraction) IsZero() bool {
	return fraction.divisor == 0
}
