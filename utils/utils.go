package utils

import (
	"encoding/csv"
	"fmt"
	"league/constant"
	"mime/multipart"
	"net/http"
	"strconv"
	"strings"
)

// ReadFile -- read input csv file
// do not allow file bigger than 10MB
func ReadFile(w http.ResponseWriter, r *http.Request) (records [][]string, err error) {
	var file multipart.File
	// do not allow file bigger than 10MB
	err = r.ParseMultipartForm(10 << 20)
	if err != nil {
		return
	}
	file, _, err = r.FormFile("file")
	if err != nil {
		return
	}
	defer file.Close()
	// read csv or txt file
	records, err = csv.NewReader(file).ReadAll()
	if err != nil {
		return
	}
	return
}

// Sanitize -- check if input of a matrix is valid
// check empty matrix
// check matrix dimension
// check value of matrix
func Sanitize(records [][]string, convertToInt64 bool) (recordsInt [][]int64, err error) {
	// check if matrix is empty
	if len(records) == 0 || len(records[0]) == 0 {
		err = constant.ErrEmptyMatrix
		return
	}
	// check dimension of matrix
	if len(records) != len(records[0]) {
		err = constant.ErrDifferentDimension
		return
	}
	recordsInt = make([][]int64, len(records))
	for i := range recordsInt {
		recordsInt[i] = make([]int64, len(records))
	}
	// check if matrix has other type than integer or interger value is overflow
	for i := 0; i < len(records); i++ {
		for j := 0; j < len(records); j++ {
			num, parseErr := strconv.ParseInt(records[i][j], 10, 64)
			if parseErr != nil {
				if strings.Contains(parseErr.Error(), "value out of range") {
					err = constant.ErrOverFlow
				} else {
					err = constant.ErrNotInteger
				}
				break
			}
			if convertToInt64 {
				recordsInt[i][j] = num
			}
		}
	}
	return
}

// CheckOverflowSum -- check if addition of two integer resulted in an overflow integer
func CheckOverflowSum(a, b int64) (int64, error) {
	c := a + b
	if (c > a) == (b > 0) {
		return c, nil
	}
	return c, constant.ErrOverFlow
}

// CheckOverflowMul -- check if multiplication of two integer resulted in an overflow integer
func CheckOverflowMul(a, b int64) (int64, error) {
	c := a * b
	if a == 0 || b == 0 || a == 1 || b == 1 {
		return c, nil
	}
	if c/b == a {
		return c, nil
	}
	return c, constant.ErrOverFlow
}

// WriteError -- wrapper for returning error response
func WriteError(w http.ResponseWriter, err error) {
	w.Write([]byte(fmt.Sprintf("error %s", err.Error())))
}
