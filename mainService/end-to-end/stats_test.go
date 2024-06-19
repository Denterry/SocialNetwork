package end_to_end

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"testing"
	"time"

	"gotest.tools/v3/assert"
)

func TestGetPostStatistics(t *testing.T) {
	// Generate a random integer between 0 and 100
	randomInt := rand.Intn(101)

	// Firstly, sign up user with secret
	SignUp(t, randomInt)

	// Firstly, sign in with secret to get the token
	token := SignIn(t, randomInt)

	// Then let's check the main functionality
	// 1. Let's check statistics before activity
	requestBody := map[string]int{
		"postId": 1,
	}

	jsonValue, _ := json.Marshal(requestBody)
	req, _ := http.NewRequest("GET", "http://localhost:8080/api/admin/posts/1/statistics", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	client := &http.Client{}
	resp, err := client.Do(req)

	assert.Equal(t, http.StatusOK, resp.StatusCode)

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

	prevTotalViews := response["totalViews"].(float64)
	prevTotalLikes := response["totalLikes"].(float64)
	fmt.Println(prevTotalViews)
	fmt.Println(prevTotalLikes)

	// 2. Let's add some activity
	requestBody = map[string]int{}

	jsonValue, _ = json.Marshal(requestBody)
	req, _ = http.NewRequest("POST", "http://localhost:8080/api/admin/posts/1/like", bytes.NewBuffer(jsonValue))
	req.Header.Set("Authorization", "Bearer "+token)

	client = &http.Client{}
	resp, _ = client.Do(req)

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	// 3. Let's chack statistics again
	requestBody = map[string]int{
		"postId": 1,
	}

	// VERY IMPORTANT IN THIS CASE WHILE CLICKHOUSE READ FROM KAFKA
	time.Sleep(time.Second * 10)

	jsonValue, _ = json.Marshal(requestBody)
	req, _ = http.NewRequest("GET", "http://localhost:8080/api/admin/posts/1/statistics", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	client = &http.Client{}
	resp, err = client.Do(req)

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	body, err = io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("Failed to read response body: %v", err)
	}

	var responseNew map[string]interface{}
	err = json.Unmarshal(body, &responseNew)
	if err != nil {
		t.Fatalf("Failed to unmarshal response: %v", err)
	}

	assert.Equal(t, prevTotalLikes+1, responseNew["totalLikes"].(float64))
	assert.Equal(t, prevTotalViews, responseNew["totalViews"].(float64))
}
