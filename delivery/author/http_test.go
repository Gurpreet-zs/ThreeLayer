package author

import (
	"net/http"
	"testing"
)

func TestBookHandler_Handler(t *testing.T) {
	testcases := []struct {
		desc string
		method   string
		expectedStatusCode int
	}{
		{"Getting Author","GET", http.},
		{"POST", http.StatusOK},
	}
}