package handler

import (
	"fmt"
	"league/utils"
	"net/http"
	"strings"
)

// EchoHandler -- handle endpoint echo, receive a matrix and return the matrix as a string in matrix format if the matrix is valid
func EchoHandler(w http.ResponseWriter, r *http.Request) {
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
	for _, row := range records {
		response = fmt.Sprintf("%s%s\n", response, strings.Join(row, ","))
	}
	fmt.Fprint(w, response)
}
