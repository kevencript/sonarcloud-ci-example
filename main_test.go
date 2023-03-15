package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandlerFirstHash(t *testing.T) {
	req, err := http.NewRequest("GET", "/first", nil)
	if err != nil {
		t.Fatal(err)
	}

	rec := httptest.NewRecorder()
	handler := http.HandlerFunc(handlerFirstHash)

	handler.ServeHTTP(rec, req)

	if status := rec.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := getFirstHash()
	if rec.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rec.Body.String(), expected)
	}
}

func TestHandlerSecondHash(t *testing.T) {
	req, err := http.NewRequest("GET", "/second", nil)
	if err != nil {
		t.Fatal(err)
	}

	rec := httptest.NewRecorder()
	handler := http.HandlerFunc(handlerSecondHash)

	handler.ServeHTTP(rec, req)

	if status := rec.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := getSecondHash()
	if rec.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rec.Body.String(), expected)
	}
}

func TestGetSecondHash(t *testing.T) {
	expected := "0cc175b9c0f1b6a831c399e269772662"
	result := getSecondHash()
	if result != expected {
		t.Errorf("getSecondHash() failed, expected %s but got %s", expected, result)
	}
}

func TestGeFirstHash(t *testing.T) {
	expected := "8743b52063cd84097a65d1633f5c74f6"
	result := getFirstHash()
	if result != expected {
		t.Errorf("getFirstHash() failed, expected %s but got %s", expected, result)
	}
}


