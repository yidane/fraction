package fraction

import (
	"fmt"
	"strconv"
)

type Fraction struct {
	divisor  int64
	dividend int64
}

func New(divisor, dividend int64) Fraction {
	if dividend == 0 {
		dividend = 1
		divisor = 0

		return Fraction{divisor: divisor, dividend: dividend}
	}

	maxDivisor := findMaxCommonDividend(divisor, dividend)

	if maxDivisor == 1 {
		return Fraction{divisor: divisor, dividend: dividend}
	}

	return Fraction{divisor: divisor / maxDivisor, dividend: dividend / maxDivisor}
}

func findMaxCommonDividend(i, j int64) int64 {
	m := i % j
	for m != 0 {
		if j > m {
			i, j = j, m
		} else {
			i = m
		}

		m = i % j
	}

	return j
}

func (fraction Fraction) Add(other Fraction) Fraction {
	maxDivisor := findMaxCommonDividend(fraction.dividend, other.dividend)
	fTimes, oTimes := fraction.dividend/maxDivisor, other.dividend/maxDivisor

	dividend := fTimes * oTimes * maxDivisor
	divisor := fraction.divisor*oTimes + other.divisor*fTimes

	return New(divisor, dividend)
}

func (fraction Fraction) Subtract(other Fraction) Fraction {
	if fraction.Equal(other) {
		return New(0, 1)
	}

	maxDivisor := findMaxCommonDividend(fraction.dividend, other.dividend)
	fTimes, oTimes := fraction.dividend/maxDivisor, other.dividend/maxDivisor

	dividend := fTimes * oTimes * maxDivisor
	divisor := (fraction.divisor*oTimes - other.divisor*fTimes) * maxDivisor

	return New(divisor, dividend)
}

func (fraction *Fraction) Multiply(other Fraction) Fraction {
	maxDivisor1 := findMaxCommonDividend(fraction.divisor, other.dividend)
	maxDivisor2 := findMaxCommonDividend(fraction.dividend, other.divisor)

	divisor := fraction.divisor * (other.dividend / maxDivisor1)
	dividend := fraction.dividend * (other.divisor / maxDivisor2)

	return New(divisor, dividend)
}

func (fraction Fraction) Divide(other Fraction) Fraction {
	if other.IsZero() {
		panic("dividend is zero")
	}

	maxDivisor := findMaxCommonDividend(fraction.dividend, other.divisor)

	dividend := fraction.dividend * (other.divisor / maxDivisor)
	divisor := fraction.divisor * (other.divisor / maxDivisor)

	return New(divisor, dividend)
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

//func (fraction Fraction)Float64Size(s int)float32{
//
//}

func (fraction Fraction) Float64() float64 {
	return float64(fraction.divisor) / float64(fraction.dividend)
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
