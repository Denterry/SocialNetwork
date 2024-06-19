package end_to_end

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"testing"

	"gotest.tools/v3/assert"
)

func SignIn(t *testing.T, test_secret int) string {
	requestBody := map[string]string{
		"username": "testuser" + strconv.Itoa(test_secret) + "@example.com",
		"password": "password123",
	}

	jsonValue, _ := json.Marshal(requestBody)
	resp, err := http.Post("http://localhost:8080/api/users/sign-in", "application/json", bytes.NewBuffer(jsonValue))

	if err != nil {
		t.Fatalf("Failed to send sign-in request: %v", err)
	}
	defer resp.Body.Close()

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("Failed to read response body: %v", err)
	}

	var response map[string]interface{}
	err = json.Unmarshal(body, &response)
	if err != nil {
		t.Fatalf("Failed to unmarshal response: %v", err)
	}

	token, ok := response["token"].(string)
	if !ok {
		t.Fatalf("Token not found in response")
	}

	return token
}
