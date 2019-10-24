package main

import (
	"bytes"
	"fmt"
	"io"
	"league/handler"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func readFile(test string) (bytes.Buffer, string) {
	var filePath string
	switch test {
	case "normal":
		filePath = "data/matrix.csv"
	case "matrix_overflow":
		filePath = "data/matrix_overflow.csv"
	case "sum_overflow":
		filePath = "data/sum_overflow.csv"
	case "mul_overflow":
		filePath = "data/mul_overflow.csv"
	case "empty_file":
		filePath = "data/empty.csv"
	}
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	var requestBody bytes.Buffer
	multipartWriter := multipart.NewWriter(&requestBody)
	fileWriter, err := multipartWriter.CreateFormFile("file", "matrix.csv")
	if err != nil {
		fmt.Println(err)
	}
	_, err = io.Copy(fileWriter, file)
	if err != nil {
		fmt.Println(err)
	}
	multipartWriter.Close()
	return requestBody, multipartWriter.FormDataContentType()
}

// read matrix from csv file and echo back
// expected 1,2,3\n4,5,6\n7,8,9\n
// actual 1,2,3\n4,5,6\n7,8,9\n
func TestEchoHandler(t *testing.T) {
	requestBody, contentType := readFile("normal")
	nr := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/echo", &requestBody)
	if err != nil {
		t.Log(err)
	}
	req.Header.Set("Content-Type", contentType)
	serverHandler := http.HandlerFunc(handler.EchoHandler)
	serverHandler.ServeHTTP(nr, req)
	if status := nr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
	expected := "1,2,3\n4,5,6\n7,8,9\n"
	result := nr.Body.String()
	if result != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", result, expected)
	}
}

// read matrix from csv file and flip the matrix
// expected 1,4,7\n2,5,8\n3,6,9\n
// actual 1,4,7\n2,5,8\n3,6,9\n
func TestInvertHandler(t *testing.T) {
	requestBody, contentType := readFile("normal")
	nr := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/invert", &requestBody)
	if err != nil {
		t.Log(err)
	}
	req.Header.Set("Content-Type", contentType)
	serverHandler := http.HandlerFunc(handler.InvertHandler)
	serverHandler.ServeHTTP(nr, req)
	if status := nr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
	expected := "1,4,7\n2,5,8\n3,6,9\n"
	result := nr.Body.String()
	if result != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", result, expected)
	}
}

// read matrix from csv file and flatten the matrix
// expected 1,2,3,4,5,6,7,8,9
// actual 1,2,3,4,5,6,7,8,9
func TestFlattenHandler(t *testing.T) {
	requestBody, contentType := readFile("normal")
	nr := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/flatten", &requestBody)
	if err != nil {
		t.Log(err)
	}
	req.Header.Set("Content-Type", contentType)
	serverHandler := http.HandlerFunc(handler.FlattenHandler)
	serverHandler.ServeHTTP(nr, req)
	if status := nr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
	expected := "1,2,3,4,5,6,7,8,9"
	result := nr.Body.String()
	if result != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", result, expected)
	}
}

// read matrix from csv file and return the sum of each element the matrix
// expected 45
// actual 45
func TestSumHandler(t *testing.T) {
	requestBody, contentType := readFile("normal")
	nr := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/sum", &requestBody)
	if err != nil {
		t.Log(err)
	}
	req.Header.Set("Content-Type", contentType)
	serverHandler := http.HandlerFunc(handler.SumHandler)
	serverHandler.ServeHTTP(nr, req)
	if status := nr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
	expected := "45"
	result := nr.Body.String()
	if result != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", result, expected)
	}
}

// read matrix from csv file and return the multiple of each element the matrix
// expected 362880
// actual 362880
func TestMultiplyHandler(t *testing.T) {
	requestBody, contentType := readFile("normal")
	nr := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/multiply", &requestBody)
	if err != nil {
		t.Log(err)
	}
	req.Header.Set("Content-Type", contentType)
	serverHandler := http.HandlerFunc(handler.MultiplyHandler)
	serverHandler.ServeHTTP(nr, req)
	if status := nr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
	expected := "362880"
	result := nr.Body.String()
	if result != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", result, expected)
	}
}

// read matrix from empty csv file
// expected error Matrix is empty
// actual error Matrix is empty
func TestEmptyFile(t *testing.T) {
	requestBody, contentType := readFile("empty_file")
	nr := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/echo", &requestBody)
	if err != nil {
		t.Log(err)
	}
	req.Header.Set("Content-Type", contentType)
	serverHandler := http.HandlerFunc(handler.EchoHandler)
	serverHandler.ServeHTTP(nr, req)
	if status := nr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
	expected := "error Matrix is empty"
	result := nr.Body.String()
	if result != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", result, expected)
	}
}

// read matrix from csv file
// expected error Number is overflow
// actual error Number is overflow
func TestMatrixOverflow(t *testing.T) {
	requestBody, contentType := readFile("matrix_overflow")
	nr := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/echo", &requestBody)
	if err != nil {
		t.Log(err)
	}
	req.Header.Set("Content-Type", contentType)
	serverHandler := http.HandlerFunc(handler.EchoHandler)
	serverHandler.ServeHTTP(nr, req)
	if status := nr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
	expected := "error Number is overflow"
	result := nr.Body.String()
	if result != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", result, expected)
	}
}

// read matrix from csv file
// expected error Number is overflow
// actual error Number is overflow
func TestSumOverflow(t *testing.T) {
	requestBody, contentType := readFile("sum_overflow")
	nr := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/sum", &requestBody)
	if err != nil {
		t.Log(err)
	}
	req.Header.Set("Content-Type", contentType)
	serverHandler := http.HandlerFunc(handler.SumHandler)
	serverHandler.ServeHTTP(nr, req)
	if status := nr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
	expected := "error Number is overflow"
	result := nr.Body.String()
	if result != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", result, expected)
	}
}

// read matrix from csv file
// expected error Number is overflow
// actual error Number is overflow
func TestMulOverflow(t *testing.T) {
	requestBody, contentType := readFile("mul_overflow")
	nr := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/mul", &requestBody)
	if err != nil {
		t.Log(err)
	}
	req.Header.Set("Content-Type", contentType)
	serverHandler := http.HandlerFunc(handler.MultiplyHandler)
	serverHandler.ServeHTTP(nr, req)
	if status := nr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
	expected := "error Number is overflow"
	result := nr.Body.String()
	if result != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", result, expected)
	}
}
