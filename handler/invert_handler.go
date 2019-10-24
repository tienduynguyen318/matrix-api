package handler

import (
	"fmt"
	"league/utils"
	"net/http"
	"strings"
)

// InvertHandler -- handle endpoint invert, return the matrix as a string in matrix format where the columns and rows are inverted if the matrix is valid
func InvertHandler(w http.ResponseWriter, r *http.Request) {
	records, err := utils.ReadFile(w, r)
	if err != nil {
		utils.WriteError(w, err)
		return
	}
	_, err = utils.Sanitize(records, false)
	if err != nil {
		utils.WriteError(w, err)
		return
	}
	var response string
	for i := 0; i < len(records); i++ {
		for j := 0; j < i; j++ {
			records[i][j], records[j][i] = records[j][i], records[i][j]
		}
	}
	for _, row := range records {
		response = fmt.Sprintf("%s%s\n", response, strings.Join(row, ","))
	}
	fmt.Fprint(w, response)
}
