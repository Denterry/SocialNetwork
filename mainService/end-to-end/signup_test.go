package end_to_end

import (
	"bytes"
	"encoding/json"
	"net/http"
	"strconv"
	"testing"

	"gotest.tools/v3/assert"
)

func SignUp(t *testing.T, test_secret int) {
	requestBody := map[string]string{
		"username":              "testuser" + strconv.Itoa(test_secret) + "@example.com",
		"password":              "password123",
		"password_confirmation": "password123",
	}

	jsonValue, _ := json.Marshal(requestBody)
	resp, err := http.Post("http://localhost:8080/api/users/sign-up", "application/json", bytes.NewBuffer(jsonValue))

	if err != nil {
		t.Fatalf("Failed to send sign-up request: %v", err)
	}
	defer resp.Body.Close()

	assert.Equal(t, http.StatusOK, resp.StatusCode)
}
