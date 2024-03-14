package jsonHelpers

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestWriteJSON(t *testing.T) {
	writer := httptest.NewRecorder()
	data := JsonResponse{
		Error:   false,
		Message: "Hit the broker",
	}
	wantStatus := http.StatusOK
	err := WriteJSON(writer, wantStatus, data)
	want := []string{`"message":"Hit the broker"`, `"error":false`}
	for i := range want {
		if !strings.Contains(writer.Body.String(), want[i]) {
			t.Fatalf(`writeJSON() with body: %q, want match for %#q, err: %v`, writer.Body.String(), want[i], err)
		}
	}
	if writer.Result().StatusCode != wantStatus {
		t.Fatalf(`writeJSON() with status code %d, want %d`, writer.Result().StatusCode, wantStatus)
	}
}
