package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
)

type SignUpRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	supabaseURL := os.Getenv("LOCAL_SUPABASE_URL")
	supabaseAnonKey := os.Getenv("LOCAL_ANON_KEY")
	numUsers := 10

	if supabaseURL == "" || supabaseAnonKey == "" {
		log.Fatalf("Error: SUPABASE_URL, SUPABASE_ANON_KEY, and NUM_USERS_TO_SEED environment variables must be set.")
	}

	authURL := fmt.Sprintf("%s/auth/v1/signup", supabaseURL)

	fmt.Printf("Attempting to seed %d users to Supabase Auth at %s\n", numUsers, authURL)

	client := &http.Client{
		Timeout: 10 * time.Second, // Set a timeout for HTTP requests
	}

	for i := 1; i <= numUsers; i++ {
		email := fmt.Sprintf("user%d@test.com", i)
		password := "123456" // Use a strong password in production

		signUpReq := SignUpRequest{
			Email:    email,
			Password: password,
		}

		jsonBody, err := json.Marshal(signUpReq)
		if err != nil {
			log.Printf("Error marshalling JSON for user %s: %v", email, err)
			continue
		}

		req, err := http.NewRequest("POST", authURL, bytes.NewBuffer(jsonBody))
		if err != nil {
			log.Printf("Error creating request for user %s: %v", email, err)
			continue
		}

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("apikey", supabaseAnonKey)                                  // Supabase Anon Key
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", supabaseAnonKey)) // Required for some Supabase endpoints

		resp, err := client.Do(req)
		if err != nil {
			log.Printf("Error sending request for user %s: %v", email, err)
			continue
		}
		defer resp.Body.Close()

		if resp.StatusCode == http.StatusCreated || resp.StatusCode == http.StatusOK {
			fmt.Printf("Successfully created user: %s\n", email)
		} else {
			var errorBody map[string]interface{}
			if err := json.NewDecoder(resp.Body).Decode(&errorBody); err != nil {
				log.Printf("Failed to decode error response for user %s: %v", email, err)
			}
			log.Printf("Failed to create user %s. Status: %s, Response: %v\n", email, resp.Status, errorBody)
		}
	}

	fmt.Println("\nUser seeding process completed.")
}
