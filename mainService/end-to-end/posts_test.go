package end_to_end

import (
	"bytes"
	"encoding/json"
	"io"
	"math/rand"
	"net/http"
	"testing"

	"github.com/google/uuid"
	"gotest.tools/v3/assert"
)

func TestCreatePost(t *testing.T) {
	// Generate a random integer between 0 and 100
	randomInt := rand.Intn(101)

	// Firstly, sign up user with secret
	SignUp(t, randomInt)

	// Firstly, sign in with secret to get the token
	token := SignIn(t, randomInt)

	// Then let's check the main functionality
	requestBody := map[string]string{
		"title":    "Test Post",
		"content":  "This is a test post",
		"authorId": uuid.Max.String(),
	}

	jsonValue, _ := json.Marshal(requestBody)
	req, _ := http.NewRequest("POST", "http://localhost:8080/api/admin/posts", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("Failed to read response body: %v", err)
	}

	var response map[string]interface{}
	err = json.Unmarshal(body, &response)
	if err != nil {
		t.Fatalf("Failed to unmarshal response: %v", err)
	}

	assert.Equal(t, http.StatusCreated, resp.StatusCode)
	assert.Equal(t, "Test Post", response["title"])
	assert.Equal(t, "This is a test post", response["content"])
}
