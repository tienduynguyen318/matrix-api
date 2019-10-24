package utils

import (
	"testing"
)

// check sanitize empty matrix
// expected Matrix is empty
// actual Matrix is empty
func TestSanitizeEmptyMatrix(t *testing.T) {
	record := [][]string{}
	_, err := Sanitize(record, false)
	t.Log(err)
	t.Error()
}

// check sanitize different dimension matrix
// expected Matrix has different dimension
// actual Matrix has different dimension
func TestSanitizeDifferentDimension(t *testing.T) {
	record := [][]string{[]string{"1", "2", "3"}}
	_, err := Sanitize(record, false)
	t.Log(err)
	t.Error()
}

// check sanitize a cell in a matrix cannot be convert to integer
// expected A cell in matrix is not an integer
// actual A cell in matrix is not an integer
func TestSanitizeNotInteger(t *testing.T) {
	record := [][]string{[]string{"1", "c"}, []string{"2", "3"}}
	_, err := Sanitize(record, false)
	t.Log(err)
	t.Error()
}

// check convert string to int64 positive overflow interger
// expected Number is overflow
// actual Number is overflow
func TestSanitizeNumberOverflowPositive(t *testing.T) {
	record := [][]string{[]string{"1", "9223372036854775808"}, []string{"2", "3"}}
	_, err := Sanitize(record, false)
	t.Log(err)
	t.Error()
}

// check convert string to int64 positive overflow interger
// expected Number is overflow
// actual Number is overflow
func TestSanitizeNumberOverflowNegative(t *testing.T) {
	record := [][]string{[]string{"1", "-9223372036854775809"}, []string{"2", "3"}}
	_, err := Sanitize(record, false)
	t.Log(err)
	t.Error()
}

// check addition create a positive overflow interger
// expected -9223372036854775808 Number is overflow
// actual -9223372036854775808 Number is overflow
func TestSumOverflowPositive(t *testing.T) {
	num, err := CheckOverflowSum(9223372036854775807, 1)
	t.Log(num)
	t.Log(err)
	t.Error()
}

// check addition create a positive overflow interger
// expected 9223372036854775807 Number is overflow
// actual 9223372036854775807 Number is overflow
func TestSumOverflowNegative(t *testing.T) {
	num, err := CheckOverflowSum(-9223372036854775808, -1)
	t.Log(num)
	t.Log(err)
	t.Error()
}

// check addition create a positive overflow interger
// expected -2 Number is overflow
// actual -2 Number is overflow
func TestMulOverflowPositive(t *testing.T) {
	num, err := CheckOverflowMul(9223372036854775807, 2)
	t.Log(num)
	t.Log(err)
	t.Error()
}

// check addition create a positive overflow interger
// expected 2 Number is overflow
// actual 2 Number is overflow
func TestMulOverflowNegative(t *testing.T) {
	num, err := CheckOverflowMul(-9223372036854775807, 2)
	t.Log(num)
	t.Log(err)
	t.Error()
}
