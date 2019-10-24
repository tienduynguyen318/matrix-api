package handler

import (
	"fmt"
	"league/utils"
	"net/http"
)

// SumHandler -- handle endpoint multiply, return the sum of the integers in the matrix if the matrix is valid and there is no overflow
func SumHandler(w http.ResponseWriter, r *http.Request) {
	records, err := utils.ReadFile(w, r)
	if err != nil {
		utils.WriteError(w, err)
		return
	}
	recordInt, err := utils.Sanitize(records, true)
	if err != nil {
		utils.WriteError(w, err)
		return
	}
	var sumMatrix int64
	for _, row := range recordInt {
		for _, val := range row {
			sumMatrix, err = utils.CheckOverflowSum(sumMatrix, val)
			if err != nil {
				utils.WriteError(w, err)
				return
			}
		}
	}
	fmt.Fprint(w, sumMatrix)
}
