package handler

import (
	"fmt"
	"league/utils"
	"net/http"
	"strings"
)

// FlattenHandler -- handle endpoint flatten, return the matrix as a 1 line string, with values separated by commas if the matrix is valid
func FlattenHandler(w http.ResponseWriter, r *http.Request) {
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
	var subResponse []string
	for _, nums := range records {
		subResponse = append(subResponse, nums...)
	}
	response = strings.Join(subResponse, ",")
	fmt.Fprint(w, response)
}
