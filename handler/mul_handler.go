package handler

import (
	"fmt"
	"league/utils"
	"net/http"
)

// MultiplyHandler -- handle endpoint multiply, return the product of the integers in the matrix if the matrix is valid and there is no overflow
func MultiplyHandler(w http.ResponseWriter, r *http.Request) {
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
	var mulMatrix int64 = 1
	for _, row := range recordInt {
		for _, val := range row {
			mulMatrix, err = utils.CheckOverflowMul(mulMatrix, val)
			if err != nil {
				utils.WriteError(w, err)
				return
			}
		}
	}
	fmt.Fprint(w, mulMatrix)
}
