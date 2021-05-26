package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetGithubRepositories(t *testing.T) {
	req, err := http.NewRequest("GET", "/any_user_name/repositories", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(getUserRepositoriesAction)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Github API always return OK for this endpoint
}
