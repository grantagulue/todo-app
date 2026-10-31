package main

// import (
// 	"bytes"
// 	"net/http"
// 	"net/http/httptest"
// 	"testing"
// )

// func TestHandleRoot(t *testing.T) {
// 	w := httptest.NewRecorder()


// 	handleRoot(w, nil)

// 	desiredCode := http.StatusOK

// 	if w.Code != desiredCode {
// 		t.Errorf("bad response code, expected %v but got %v\nbody: %s\n", desiredCode, w.Code, w.Body.String())
// 	}

// 	expectedMessage := []byte("Welcome to our Homepage!\n")

// 	if !bytes.Equal(expectedMessage, w.Body.Bytes()) {
// 		t.Errorf("bad return, got %q, expected %q", w.Body.String(), expectedMessage)
// 	}
// }

// func TestHandleGoodbye(t *testing.T) {
// 	w := httptest.NewRecorder()


// 	handleGoodbye(w, nil)

// 	desiredCode := http.StatusOK

// 	if w.Code != desiredCode {
// 		t.Errorf("bad response code, expected %v but got %v\nbody: %s\n", desiredCode, w.Code, w.Body.String())
// 	}

// 	expectedMessage := []byte("Goodbye!\n")

// 	if !bytes.Equal(expectedMessage, w.Body.Bytes()) {
// 		t.Errorf("bad return, got %q, expected %q", w.Body.String(), expectedMessage)
// 	}
// }

// func TestHandleParameterized(t *testing.T) {
// 	req := httptest.NewRequest(http.MethodGet, "/hello?user=TestMan", nil)

// 	w := httptest.NewRecorder()

// 	handleParameterized(w, req)

// 	desiredCode := http.StatusOK
	
// 	if w.Code != desiredCode {
// 		t.Error("bad response code")
// 	}
// }