package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

type Payload struct {
	Service                 string                  `json:"service"`
	AuthPayload             AuthPayload             `json:"auth_payload,omitempty"`
	CompleteRegisterPayload CompleteRegisterPayload `json:"complete_register_payload,omitempty"`
}
type CompleteRegisterPayload struct {
	ID         string   `json:"id"`
	RawID      string   `json:"rawId"`
	Response   response `json:"response"`
	Type       string   `json:"type"`
	SessionKey string   `json:"sessionKey"`
}

type response struct {
	AttestationObject string `json:"attestationObject"`
	ClientDataJSON    string `json:"clientDataJSON"`
}

type AuthPayload struct {
	Email string `json:"email"`
}

// Read and parse the response body with structured typing
type ResponseData struct {
	Data struct {
		Options    map[string]any `json:"options"`
		SessionKey string         `json:"sessionKey"`
	} `json:"data"`
	Error   bool   `json:"error"`
	Message string `json:"message"`
}

func (app *application) Gateway(w http.ResponseWriter, r *http.Request) {
	var payload Payload
	err := app.readJSON(r, &payload)
	if err != nil {
		app.badRequestResponse(w, r, err)
	}

	switch payload.Service {
	case "begin-register":
		app.HandleBeginRegisterRoute(w, r, payload.AuthPayload)
	case "complete-register":
		app.HandleCompleteRegister(w, r, payload.CompleteRegisterPayload)
	case "begin-login":
		fmt.Print(payload.AuthPayload.Email)
		app.HandleBeginLogin(w, r, payload.AuthPayload)
	}

}

func (app *application) HandleBeginRegisterRoute(w http.ResponseWriter, r *http.Request, payload AuthPayload) {
	// Marshal the existing payload to JSON bytes
	jsonBody, err := json.Marshal(payload)
	if err != nil {
		fmt.Printf("error marshaling payload: %s\n", err)
		os.Exit(1)
	}
	bodyReader := bytes.NewReader(jsonBody)

	// Make the post request to the auth service with the body
	req, err := http.NewRequest(http.MethodPost, "http://localhost:8080/begin-register", bodyReader)
	if err != nil {
		fmt.Printf("client: could not create request: %s\n", err)
		os.Exit(1)
	}
	req.Header.Set("Content-Type", "application/json")

	// Set a timeout for 30 seconds to receive the response
	client := http.Client{
		Timeout: 30 * time.Second,
	}

	// execute the request
	res, err := client.Do(req)
	if err != nil {
		fmt.Printf("client: error making http request: %s\n", err)
		os.Exit(1)
	}
	defer res.Body.Close()

	var authResponse ResponseData
	body, err := io.ReadAll(res.Body)
	if err != nil {
		app.errorJSON(w, fmt.Errorf("error reading response body: %v", err), http.StatusInternalServerError)
		return
	}
	if err := json.Unmarshal(body, &authResponse); err != nil {
		app.errorJSON(w, fmt.Errorf("error parsing response: %v", err), http.StatusInternalServerError)
		return
	}

	// Check if auth service returned an error
	if authResponse.Error {
		app.errorJSON(w, fmt.Errorf("auth service error: %s", authResponse.Message), res.StatusCode)
		return
	}

	// Forward the entire response with original status code
	app.writeJSON(w, res.StatusCode, authResponse.Message, authResponse.Data)
}

func (app *application) HandleCompleteRegister(w http.ResponseWriter, r *http.Request, payload CompleteRegisterPayload) {
	// Marshal the existing payload to JSON bytes
	jsonBody, err := json.Marshal(payload)
	if err != nil {
		fmt.Printf("error marshaling payload: %s\n", err)
		os.Exit(1)
	}
	bodyReader := bytes.NewReader(jsonBody)

	// Make the post request to the auth service with the body
	req, err := http.NewRequest(http.MethodPost, "http://localhost:8080/complete-register", bodyReader)
	if err != nil {
		fmt.Printf("client: could not create request: %s\n", err)
		os.Exit(1)
	}
	req.Header.Set("Content-Type", "application/json")

	// Set a timeout for 30 seconds to recieve the response
	client := http.Client{
		Timeout: 30 * time.Second,
	}

	// execute the request
	res, err := client.Do(req)
	if err != nil {
		fmt.Printf("client: error making http request: %s\n", err)
		os.Exit(1)
	}
	defer res.Body.Close()

	var authResponse ResponseData
	body, err := io.ReadAll(res.Body)
	if err != nil {
		app.errorJSON(w, fmt.Errorf("error reading response body: %v", err), http.StatusInternalServerError)
		return
	}
	if err := json.Unmarshal(body, &authResponse); err != nil {
		app.errorJSON(w, fmt.Errorf("error parsing response: %v", err), http.StatusInternalServerError)
		return
	}

	// Check if auth service returned an error
	if authResponse.Error {
		app.errorJSON(w, fmt.Errorf("auth service error: %s", authResponse.Message), res.StatusCode)
		return
	}

	// Forward the entire response with original status code
	app.writeJSON(w, res.StatusCode, authResponse.Message, authResponse.Data)
}

func (app *application) HandleBeginLogin(w http.ResponseWriter, r *http.Request, payload AuthPayload) {
	// Marshal the existing payload to JSON bytes
	jsonBody, err := json.Marshal(payload)
	if err != nil {
		fmt.Printf("error marshaling payload: %s\n", err)
		os.Exit(1)
	}
	bodyReader := bytes.NewReader(jsonBody)

	// Make the post request to the auth service with the body
	req, err := http.NewRequest(http.MethodPost, "http://localhost:8080/begin-login", bodyReader)
	if err != nil {
		fmt.Printf("client: could not create request: %s\n", err)
		os.Exit(1)
	}
	req.Header.Set("Content-Type", "application/json")

	// Set a timeout for 30 seconds to receive the response
	client := http.Client{
		Timeout: 30 * time.Second,
	}

	// execute the request
	res, err := client.Do(req)
	if err != nil {
		fmt.Printf("client: error making http request: %s\n", err)
		os.Exit(1)
	}
	defer res.Body.Close()

	var authResponse ResponseData
	body, err := io.ReadAll(res.Body)
	if err != nil {
		app.errorJSON(w, fmt.Errorf("error reading response body: %v", err), http.StatusInternalServerError)
		return
	}
	if err := json.Unmarshal(body, &authResponse); err != nil {
		app.errorJSON(w, fmt.Errorf("error parsing response: %v", err), http.StatusInternalServerError)
		return
	}

	// Check if auth service returned an error
	if authResponse.Error {
		app.errorJSON(w, fmt.Errorf("auth service error: %s", authResponse.Message), res.StatusCode)
		return
	}

	// Forward the entire response with original status code
	app.writeJSON(w, res.StatusCode, authResponse.Message, authResponse.Data)
}
