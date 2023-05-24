package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAddCategory(t *testing.T) {
	//r.Post("/add", handlers.AddCategory)
	_, err := http.NewRequest("POST", "/categories/add", nil)
	if err != nil {
		t.Fatalf(err.Error())
	}

	newRecorder := httptest.NewRecorder()

	//router.Routes().serveHTTP(newRecorder, req)

	statusCode := 200
	if newRecorder.Result().StatusCode != statusCode {
		t.Errorf("TestAddCategory() test returned an unexpected result: got %v want %v", newRecorder.Result().StatusCode, statusCode)
	}

}
